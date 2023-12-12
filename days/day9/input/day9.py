lines = open("input/0.txt","rt").read().splitlines()

def f1(a):
  if a[0]==a[1]==a[-1]:
    return a[-1]
  return a[-1]+f1([a[i]-a[i-1] for i in range(1,len(a))])

for l in lines:
  print(f1([int(e) for e in l.split()]))

def f2(a):
  if a[0]==a[1]==a[-1]:
    return a[0]
  return a[0]-f2([a[i]-a[i-1] for i in range(1,len(a))])

print(sum(f2([int(e) for e in l.split()]) for l in lines))
