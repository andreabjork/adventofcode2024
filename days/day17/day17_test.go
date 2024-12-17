package day17 

import (
	"testing"
)

func Test1(t *testing.T) {
	c := &Computer{0, 0, 9, []int{2,6}, 0, "", nil}
	c.initialize()
	c.run()

	if c.B != 1 {
		t.Fatalf(`Want B=1 when running program=%+v; got A,B,C=%d,%d,%d,strip=%s("")`, c.program, c.A, c.B, c.C, c.strip)
	}
}

func Test2(t *testing.T) {
	c := &Computer{10, 0, 0, []int{5,0,5,1,5,4}, 0, "", nil}
	c.initialize()
	c.run()

	if c.strip != "0,1,2," {
		t.Fatalf(`Want strip=0,1,2, when running program=%+v; got A,B,C=%d,%d,%d,strip=%s("")`, c.program, c.A, c.B, c.C, c.strip)
	}
}

func Test3(t *testing.T) {
	c := &Computer{2024, 0, 0, []int{0,1,5,4,3,0}, 0, "", nil}
	c.initialize()
	c.run()

	if c.strip != "4,2,5,6,7,7,7,7,3,1,0," {
		t.Fatalf(`Want strip=4,2,5,6,7,7,7,7,3,1,0, when running program=%+v; got A,B,C=%d,%d,%d,strip=%s("")`, c.program, c.A, c.B, c.C, c.strip)
	}

	if c.A != 0 {
		t.Fatalf(`Want c.A = 0, when running program=%+v; got A,B,C=%d,%d,%d,strip=%s("")`, c.program, c.A, c.B, c.C, c.strip)
	}
}

func Test4(t *testing.T) {
	c := &Computer{0, 29, 0, []int{1,7}, 0, "", nil}
	c.initialize()
	c.run()

	if c.B != 26 {
		t.Fatalf(`Want c.B = 26, when running program=%+v; got A,B,C=%d,%d,%d,strip=%s("")`, c.program, c.A, c.B, c.C, c.strip)
	}
}

func Test5(t *testing.T) {
	c := &Computer{0, 2024, 43690, []int{4,0}, 0, "", nil}
	c.initialize()
	c.run()

	if c.B != 44354 {
		t.Fatalf(`Want c.B = 44354, when running program=%+v; got A,B,C=%d,%d,%d,strip=%s("")`, c.program, c.A, c.B, c.C, c.strip)
	}
}

func Test6(t *testing.T) {
	c := &Computer{729, 0, 0, []int{0,1,5,4,3,0}, 0, "", nil}
	c.initialize()
	c.run()

	if c.strip != "4,6,3,5,6,3,5,2,1,0," {
		t.Fatalf(`Want strip=4,6,3,5,6,3,5,2,1,0 when running program=%+v; got A,B,C=%d,%d,%d,strip=%s("")`, c.program, c.A, c.B, c.C, c.strip)
	}
}