package main

import "fmt"

func main() {
	//var batman string = "I am batman"
	//fmt.Println(batman)

	//var superman string
	//superman = "I am superman"
	//fmt.Println(superman)

	//thor := "I am thor"
	//fmt.Println(thor)

	//var defaultValue string
	//fmt.Println(defaultValue)

	//var (
	//	spiderman = "I am spiderman"
	//	age       = 18
	//	powers    = "web slings, spidy sense"
	//)
	//fmt.Println(spiderman, age, powers)

	//fmt.Print("Enter your name: ")
	//var input string
	//fmt.Scanf("%s", &input)

	//fmt.Println(input)

	//for i := 1; i <= 10; i++ {
	//fmt.Println(i)
	//	if i%2 == 0 {
	//even number
	//		fmt.Println(i, "is an even number")
	//	} else {
	//odd
	//		fmt.Println(i, "is an odd number")
	//	}
	//}

	//var x [5]float64
	//x[0] = 98
	//x[1] = 93
	//x[2] = 77
	//x[3] = 82
	//x[4] = 83

	//var total float64 = 0
	//for i := 0; i < len(x); i++ {
	//		total += x[i]
	//		fmt.Println(i)
	//	}
	//	fmt.Println(total / float64(len(x)))

	//slice1 := []int{1, 2, 3}
	//slice2 := make([]int, 2)
	//copy(slice2, slice1)
	//fmt.Println(slice1, slice2)

	x := make(map[string]int)
	x["key"] = 10
	fmt.Println("%v %T", x, x)

}
