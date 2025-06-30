package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	// ТИП: Целочисленное число
	var i int = -42
	var i8 int8 = -8
	var i16 int16 = -30000
	var i32 int32 = -100000
	var i64 int64 = -999999999

	var ui uint = 42
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615
	var ptr uintptr = 0xFFFFFFFF

	// ТИП: числа с плавающей точкой
	var f32 float32 = 3.14
	var f64 float64 = 2.718281828

	// ТИП: Комплексные числа
	var c64 complex64 = complex(2, 3)
	var c128 complex128 = cmplx.Sqrt(-1)

	// ТИП: Байт и руна
	var b byte = 'A'
	var r rune = 'Ж'

	// ТИП: Строка
	var s string = "Привет, Go!"

	// ТИП: Булевый тип
	var ok bool = true

	// ВЫВОД:
	fmt.Println("int:", i)
	fmt.Println("int8:", i8)
	fmt.Println("int16:", i16)
	fmt.Println("int32:", i32)
	fmt.Println("int64:", i64)
	fmt.Println("uint:", ui)
	fmt.Println("uint8:", ui8)
	fmt.Println("uint16:", ui16)
	fmt.Println("uint32:", ui32)
	fmt.Println("uint64:", ui64)
	fmt.Println("uintptr:", ptr)

	fmt.Println("float32:", f32)
	fmt.Println("float64:", f64)

	fmt.Println("complex64:", c64)
	fmt.Println("complex128:", c128)

	fmt.Println("byte (ASCII):", b, string(b))
	fmt.Println("rune (Unicode):", r, string(r))

	fmt.Println("string:", s)
	fmt.Println("bool:", ok)
}
