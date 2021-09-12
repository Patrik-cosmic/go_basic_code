package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
)

// Golang Program to Find Largest Element of an Array

// func max(numbers ...int) (int, error) {
// 	if len(numbers) == 0 {
// 		return 0, errors.New("Empty")
// 	}
// 	max := numbers[0]
// 	for _, number := range numbers {
// 		if number > max {
// 			max = number
// 		}
// 	}
// 	return max, nil
// }

// func main() {
// 	/*
// 		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 		result, err := max(numbers...)
// 		result, err := max(1, 2, 3, 4, 5)
// 	*/
// 	fmt.Println("How many numbers? :")
// 	var n int
// 	fmt.Scanf("%d", &n)
// 	if n <= 0 {
// 		panic(fmt.Sprintf("Are you mad or what !"))
// 	}
// 	var numbers [100]int
// 	fmt.Println("Enter the numbers: ")
// 	for i := 0; i < n; i++ {
// 		fmt.Scanf("%d", &numbers[i])
// 	}
// 	result, err := max(numbers[:]...)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("The biggest numer is : ", result)
// 	}

// }

// Golang Program to check Armstrong Number 153 = 1*1*1 + 5*5*5 + 3*3*3

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func isArmstring(number int) bool {
	count, _ := fmt.Print(number)
	CallClear()
	var result int
	var temp_number = number

	for temp_number != 0 {
		r := temp_number % 10
		result += powInt(r, count)
		temp_number /= 10
	}

	if result == number {
		return true
	} else {
		return false
	}
}

func main() {
	number := 153
	fmt.Println(isArmstring(number))
}
