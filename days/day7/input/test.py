def create_equations(equations):
    split = [e.split(": ") for e in equations]
    read = [[int(e[0]), [int(x) for x in e[1].split(" ")]] for e in split]
    return read


def is_calculable(total, nums):
    queue = [[total, nums]]
    res = False
    while len(queue) > 0:
        total_i, nums_i = queue.pop(0)

        if len(nums_i) == 1:
            # finished with this combination
            if total_i == nums_i[0]:
                res = True
                break
            else:
                continue

        x = nums_i.pop()

        queue.append([total_i - x, [i for i in nums_i]])
        if total_i % x == 0:
            # only add the multiplication as a possibility if the
            # total is divisble by x
            queue.append([total_i / x, [i for i in nums_i]])

    return res


def sum_valid(equations):
    return sum([e[0] for e in equations if is_calculable(e[0], e[1])])


def is_calculable_2(total, nums):
    # same approach as part 1, but with concatenation as well
    queue = [[total, nums]]
    res = False

    while len(queue) > 0:
        total_i, nums_i = queue.pop(0)
        # little efficiency boost
        if total_i < 0:
            continue
        total_i = int(total_i)

        if len(nums_i) == 1:
            if total_i == nums_i[0]:
                res = True
                break
            else:
                continue

        x = nums_i.pop()

        # add combinations to queue
        # addition:
        queue.append([total_i - x, [i for i in nums_i]])
        # multiplcation:
        if total_i % x == 0:
            # only added if the total is divisible by x
            queue.append([total_i / x, [i for i in nums_i]])
        # concatenation:
        if str(total_i).endswith(str(x)) and len(str(total_i)) > len(str(x)):
            # only added if the total ends with x
            total_before_concat = int(str(total_i)[:-len(str(x))])
            queue.append([total_before_concat, [i for i in nums_i]])

    if res:
      print(total)
    return res


def sum_valid_2(equations):
    return sum([e[0] for e in equations if is_calculable_2(e[0], e[1])])


if __name__ == "__main__":
    with open("days/day7/input/1.txt") as f:
        input = [l.strip() for l in f.readlines()]
        equations = create_equations(input)
        t = sum_valid(equations)
        print(t, "< part 1")
        equations = create_equations(input)
        t_2 = sum_valid_2(equations)
        print(t_2, "< part 2")