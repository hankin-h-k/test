package main

import("fmt")

func main() {
	//var scores [10]int
	//scores[0] = 339
	// scores := [10]int{100,3300,4005,060}
	// for index,value := range scores {
	// 	print(index, value)
	// 	println("\n")
	// }

	scores := []int{1,4,293,4,9}
	fmt.Println(scores)
	scores2 := make([]int, 10)
	scores2[4] = 200
	fmt.Println(scores2)
	scores3 := make([]int, 0, 10)
	scores3 = append(scores3, 5)
	fmt.Println(scores3)
}
