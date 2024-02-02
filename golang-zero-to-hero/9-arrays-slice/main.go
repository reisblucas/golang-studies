package main

import (
	"fmt"
	"reflect"
)

func main() {
	// numbers defined the type is array and not dynamic
	var array1 [5]string
	array1[0] = "Position 1"
	array1[3] = "Position 4"
	array1[2] = "Position 3"
	array1[1] = "Position 2"
	fmt.Println(array1)

	array2 := [5]string{"Position 1", "Position 2", "Position 3", "Position 4", "Position 5"}
	fmt.Println(array2)
	// invalid argument: index 5 out of bounds [0:5]compilerInvalidIndex
	// array2[5] = "Position 6"

	fixedPositionBasedOnArguments := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(fixedPositionBasedOnArguments, len(fixedPositionBasedOnArguments))

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice))

	slice = append(slice, 6, 7)
	fmt.Println(slice, len(slice))

	strSlice := []string{}
	strSlice = append(strSlice, "string")
	fmt.Println(strSlice, len(strSlice))

	fmt.Println(reflect.TypeOf(fixedPositionBasedOnArguments))
	fmt.Println(reflect.TypeOf(slice))

	fmt.Println(fixedPositionBasedOnArguments[2])
	fmt.Println(fixedPositionBasedOnArguments[1:3])
	fmt.Println(fixedPositionBasedOnArguments[:3])
	fmt.Println(fixedPositionBasedOnArguments[:])

	sliceWithMemoryReference := fixedPositionBasedOnArguments[0:2] // pointer to memory position of slice
	fmt.Println(sliceWithMemoryReference)

	fmt.Println("-------")
	fmt.Println("before", fixedPositionBasedOnArguments)
	sliceWithMemoryReference = append(sliceWithMemoryReference, 7, 8, 9, 10)
	fmt.Println("after", fixedPositionBasedOnArguments)

	fixedPositionBasedOnArguments[1] = 999
	fmt.Println(sliceWithMemoryReference)

	fmt.Println("-----No Memory Reference------")
	sliceWithNewMemoryReference := fixedPositionBasedOnArguments[:]
	fmt.Println("before", sliceWithMemoryReference)
	sliceWithNewMemoryReference = append(sliceWithNewMemoryReference, 200, 300, 400)
	fmt.Println("after", sliceWithMemoryReference)
	fixedPositionBasedOnArguments[1] = 0

	fmt.Println("pointed slice", fixedPositionBasedOnArguments)
	fmt.Println("after change pointed slice", sliceWithNewMemoryReference)
}
