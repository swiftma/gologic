package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("please input password")

	password := 6789
	num := 0
	var err error

	for reader.Scan() {
		num, err = strconv.Atoi(reader.Text())
		if err == nil && num == password {
			break
		}
		if err != nil {
			fmt.Printf("invalid input: %v\n", err)
		} else if num != password {
			fmt.Println("password mismatch")
		}
		fmt.Println("please input password again")
	}
	fmt.Println("correct")
}
