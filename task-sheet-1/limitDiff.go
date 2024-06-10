package main

import (
	"fmt"
	"math"
)

func main() {

	var int64Max int64 = math.MaxInt64
	var int64Min int64 = math.MinInt64
	var int32Max int64 = math.MaxInt32
	var int32Min int64 = math.MinInt32

	var float32Max float32 = math.MaxFloat32
	var float32Min float32 = -math.MaxFloat32
	var float64Max float64 = math.MaxFloat64
	var float64Min float64 = -math.MaxFloat64

	fmt.Println("int64 Max:", int64Max)
	fmt.Println("int64 Max:", int64Min)
	fmt.Println("int32 Max:", int32Min)
	fmt.Println("int32 Max:", int32Max)
	fmt.Println("float32 Max:", float32Max)
	fmt.Println("float32 Min:", float32Min)
	fmt.Println("float64 Max:", float64Max)
	fmt.Println("float64 Min:", float64Min)
}
