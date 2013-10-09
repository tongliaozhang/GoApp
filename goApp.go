package main

import (
	"GoApp/fileDeal"
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := initAppConsole()

	chooseMethod(input)

	//fileDeal.ReadFileAndPrint("test.txt")
}

func initAppConsole() string {
	fmt.Println("*****************************")
	fmt.Println("*  Welcom GoApp ......      *")
	fmt.Println("*  1 FileDeal               *")
	fmt.Println("*  0 exit                   *")
	fmt.Println("*****************************")
	fmt.Println("Please choose function code : ")

	reader := bufio.NewReader(os.Stdin)

	input, _, _ := reader.ReadLine()

	return string(input)
}

func chooseMethod(input string) {
	switch input {
	case "1":
		flag := fileDeal.ReadFileAndPrint()

		if flag {
			sr := initAppConsole()
			chooseMethod(sr)
		} else {
			sr := initAppConsole()
			chooseMethod(sr)
		}
	case "0":
		os.Exit(0)
	default:
		fmt.Println("input error,please try again!")

		reader := bufio.NewReader(os.Stdin)
		tmp, _, _ := reader.ReadLine()

		chooseMethod(string(tmp))
	}
}
