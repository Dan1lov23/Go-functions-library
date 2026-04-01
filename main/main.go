package main

import (
	"fmt"
	"func/arr"
	"func/str"
)

func main() {

	fmt.Println(arr.NumsArrSlice([]int{1, 2, 3}, 0, 1))
	fmt.Println(arr.StrArrSlice([]string{"lol", "kek", "haha"}, 0, 1))

	fmt.Println(arr.NumsArrInclude([]int{1, 2, 3}, 2))
	fmt.Println(arr.StrArrInclude([]string{"lol", "kek", "haha"}, "lol"))

	fmt.Println(str.ChangeElement("kotiks", 5, "i"))

	fmt.Println(str.DeleteElement("kotiks", 3))

	fmt.Println(str.ReverseString("kotiks"))

	fmt.Println(str.RepeatElementCounter("kotik", "k"))

}
