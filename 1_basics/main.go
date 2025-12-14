package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func Packages() {
	fmt.Println("Packages: My favorite number is", rand.Intn(100))
}

func Imports() {
	fmt.Printf("Imports: Now you have %g problems.\n", math.Sqrt(8))
}

func ExportedNames() {
	fmt.Println("ExportedNames:", math.Pi)
}

func Functions() {
	add := func(x int, y int) int {
		return x + y
	}
	fmt.Println("Functions:", add(123, 456))
}

func FunctionsContinued() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println("FunctionsContinued:", add(111, 222))
}

func MultipleResults() {
	swap := func(x, y string) (string, string) {
		return y, x
	}
	a, b := swap("hoge", "fuga")
	fmt.Println("MultipleResults:", a, b)
}

func NamedReturnValues() {
	split := func(sum int) (x, y int) {
		x = sum * 4 / 9
		y = sum - x
		return
	}
	a, b := split(100)
	fmt.Println("NamedReturnValues:", a, b)
}

func Variables() {
	var c, python, java bool
	fn := func() {
		var i int
		fmt.Println("Variables:", i, c, python, java)
	}
	fn()
}

func VariablesWithInitializers() {
	var i, j int = 1, 2
	fn := func() {
		var c, python, java = true, false, "no!"
		fmt.Println("VariablesWithInitializers:", i, j, c, python, java)
	}
	fn()
}

func ShortVariableDeclarations() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println("ShortVariableDeclarations:", i, j, k, c, python, java)
}

func BasicTypes() {
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fn := func() {
		fmt.Printf("BasicTypes: Type: %T Value: %v\n", ToBe, ToBe)
		fmt.Printf("BasicTypes: Type: %T Value: %v\n", MaxInt, MaxInt)
		fmt.Printf("BasicTypes: Type: %T Value: %v\n", z, z)
	}
	fn()
}

func ZeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("ZeroValues: %v %v %v %q\n", i, f, b, s)
}

func TypeConversations() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println("TypeConversations:", x, y, z)
}

func TypeInference() {
	v := 123.345
	fmt.Printf("TypeInference: v is of type %T\n", v)
}

func Constants() {
	const Pi = 3.14
	const World = "世界"
	fmt.Println("Constants: Hello", World)
	fmt.Println("Constants: Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Constants: Go rules?", Truth)
}

func NumericConstants() {
	const (
		Big   = 1 << 10
		Small = Big >> 9
	)
	needInt := func(x int) int { return x*10 + 1 }
	needFloat := func(x float64) float64 {
		return x * 0.1
	}
	fmt.Println("NumericConstants:", needInt(Small))
	fmt.Println("NumericConstants:", needInt(Big))
	fmt.Println("NumericConstants:", needFloat(Small))
	fmt.Println("NumericConstants:", needFloat(Big))
}

func main() {
	fmt.Println("-----------Basic Tours-----------")
	Packages()
	Imports()
	ExportedNames()
	Functions()
	FunctionsContinued()
	MultipleResults()
	NamedReturnValues()
	Variables()
	VariablesWithInitializers()
	ShortVariableDeclarations()
	BasicTypes()
	ZeroValues()
	TypeConversations()
	TypeInference()
	Constants()
	NumericConstants()
}
