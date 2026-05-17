package main

import (
	"fmt"
	"strconv"
)

func main() {
	resultBool, err := strconv.ParseBool("kjas")
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println(resultBool)
	}

	resultInt, err := strconv.Atoi("10")
	//resultInt, err := strconv.ParseInt("10", 10, 64)
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	var stringInt string = strconv.Itoa(1000)
	fmt.Println(stringInt)
}
