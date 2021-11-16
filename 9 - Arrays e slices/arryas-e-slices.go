package main

import "fmt"

func main() {
	var array1 [5]int8
	array1[0] = 10
	fmt.Println(array1)

	array2 := [5]int8{1, 2, 3, 4, 5}
	fmt.Println(array2)

	
	array3 := [...]int8{1, 2, 3, 4, 5}
	fmt.Println(array3)


	slice := []int8{1, 2, 3, 4, 5}
	fmt.Println(slice)
	
	slice = append(slice, 99)
	fmt.Println(slice)

	slice2 := array3[0:3]
	fmt.Println(slice2)

	// Arratys internos

	slice3 := make([]int8, 5, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 1)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 1)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]int8, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))




}