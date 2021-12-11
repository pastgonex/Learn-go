/*
This package is used to test the Go-Effective
*/
package test

import (
	"fmt"
	"testing"
)

// TestIfElse test [if else]
func TestIfElse(t *testing.T) {
	x, y := 1, 2
	if x < y {
		fmt.Println("x < y")
	}
	fmt.Println("x >= y")
}

func TestSwitch(t *testing.T) {
	x := 1
	switch x {
	case 1, 2, 3, 4, 5, 6, 7:
		fmt.Println("1 <= x <= 7")
	case 10:
		fmt.Println("x == 10")
	}

	switch {
	case x == 0:
		fmt.Println("x == 0")
	case
		x == 1:
		fmt.Println("x == 1")
	}
}

func TestParallelAssignment(t *testing.T) {
	x, y := 1, 2
	x, y = y, x
	fmt.Println(x, y)
}

func TestRedeclarationAndReassignment(t *testing.T) {
	x := 1
	x, y := 2, 3
	fmt.Println(x, y)
}

func TestBreakLoop(t *testing.T) {
Loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			fmt.Println(i)
			if i == 6 {
				break Loop
			}
		}
	}
}

func TestTypeSwitch(t *testing.T) {
	var x interface{} = -1 + 3i
	switch x := x.(type) {
	case int:
		fmt.Printf("x is %d\n", x)
	case string:
		fmt.Printf("x is %s\n", x)
	default:
		fmt.Printf("x is %T\n", x)
	}

	switch x.(type) {
	case int:
		fmt.Println("x is int")
	case string:
		fmt.Println("x is string")
	case bool:
		fmt.Println("x is bool")
	default:
		fmt.Println("x is unknown")
	}
}

func TestConstructorsAndCompositeLiterals(t *testing.T) {
	var xx = make([]int, 10, 100)
	fmt.Println(xx)
	var x, y, z int = 1, 2, 3
	fmt.Println(x, y, z)

	var a, b, c = true, "foo", 3.14
	fmt.Println(a, b, c)

	var d = []int{1, 2, 3}
	fmt.Println(d)

	var e = map[string]int{"foo": 1, "bar": 2}
	fmt.Println(e)

	var f = struct {
		a int
		b string
	}{1, "foo"}
	fmt.Println(f)
}

func TestPrinting(t *testing.T) {
	fmt.Print(1, 2, "\n")
	fmt.Print("hello", 1)
}

type ByteSlice []byte

func TestMethod(t *testing.T) {
	var b ByteSlice
	_, err := fmt.Fprintf(&b, "This hour has %d days\n", 7)
	if err != nil {
		return
	}
	fmt.Printf("%s", b)
}

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := data
	*p = slice
	return len(data), nil
}
