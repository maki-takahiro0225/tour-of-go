package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func For() {
	sum := 0
	for i := 0; i < 20; i++ {
		sum += i
	}
	fmt.Println("For:", sum)
}

func ForContinuedAndWhile() {
	sum := 1
	/**
	 * 省略せずに書くと以下のようになる
	 * for ; sum < 1000; {
	 */
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("For continued and while:", sum)
}

func If() {
	var sqrt func(x float64) string
	sqrt = func(x float64) string {
		if x < 0 {
			return sqrt(-x) + "i"
		}
		return fmt.Sprint(math.Sqrt(x))
	}
	fmt.Println("If:", sqrt(2), sqrt(-4))
}

func IfWithAShortStatement() {
	pow := func(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		}
		return lim
	}
	fmt.Println("If with a short statement:", pow(3, 2, 10), pow(3, 3, 20))
}

func IfAndElse() {
	pow := func(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		} else {
			fmt.Printf("If and else: %g >= %g\n", v, lim)
		}
		return lim
	}
	fmt.Println("If and else:", pow(3, 2, 10), pow(3, 3, 20))
}

func LoopsAndFunctions() {
	Sqrt := func(x float64) float64 {
		if x == 0 {
			return 0
		} else if x < 0 {
			return math.NaN()
		} else {
			z := 1.0
			for i := 0; i < 10; i++ {
				z -= (z*z - x) / (2 * z)
			}
			return z
		}
	}
	x := float64(-1)
	fmt.Println("LoopsAndFunctions:", Sqrt(x), math.Sqrt(x))
}

func Switch() {
	fmt.Print("Switch: Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func SwitchEvaluationOrder() {
	fmt.Println("SwitchEvaluationOrder: When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("SwitchEvaluationOrder: Today.")
	case today + 1:
		fmt.Println("SwitchEvaluationOrder: Tomorrow.")
	case today + 2:
		fmt.Println("SwitchEvaluationOrder: In two days.")
	default:
		fmt.Println("SwitchEvaluationOrder: Too far away.")
	}
}

func SwitchWithNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("SwitchWithNoCondition: Good morning!")
	case t.Hour() < 17:
		fmt.Println("SwitchWithNoCondition: Good afternoon.")
	default:
		fmt.Println("SwitchWithNoCondition: Good evening.")
	}
}

func Defer() {
	defer fmt.Println("world")
	fmt.Print("Defer: hello ")
}

func StackingDefers() {
	fmt.Println("StackingDefers: counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println("StackingDefers:", i)
	}
	fmt.Println("StackingDefers: done")
}

func main() {
	fmt.Println("-----------Flowcontrol Tours-----------")
	For()
	ForContinuedAndWhile()
	If()
	IfWithAShortStatement()
	IfAndElse()
	LoopsAndFunctions()
	Switch()
	SwitchEvaluationOrder()
	SwitchWithNoCondition()
	Defer()
	StackingDefers()
}
