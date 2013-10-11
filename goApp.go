package main

import (
	"GoApp/fileDeal"
	"GoApp/oop"
	"bufio"
	"fmt"
	"os"
)

func main() {

	input := initAppConsole()

	chooseMethod(input)
}

func initAppConsole() string {
	fmt.Println("*****************************")
	fmt.Println("*  Welcom GoApp ......      *")
	fmt.Println("*  1 FileDeal               *")
	fmt.Println("*  2 SortInterface          *")
	fmt.Println("*  3 OOP                    *")
	fmt.Println("*  0 exit                   *")
	fmt.Println("*****************************")
	fmt.Println("Please choose function code : ")

	reader := bufio.NewReader(os.Stdin)

	input, _, _ := reader.ReadLine()

	return string(input)
}

type SiginChan struct {
	Id  int
	msg string
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
	case "2":

		c := make(chan interface{})

		go func() {
			ints := oop.Xi{100, 78, 99, 30, 101}
			strings := oop.Si{"111", "333", "abc", "0"}
			oop.Sort(ints)
			fmt.Printf("%v\n", ints)
			oop.Sort(strings)
			fmt.Printf("%v\n", strings)

			c <- 1
			c <- "test"
			c <- &SiginChan{1, "success"}
			c <- &SiginChan{0, "error"}
			close(c)
		}()

		for v := range c {
			fmt.Println("c value :", v)
		}

		sr := initAppConsole()
		chooseMethod(sr)

	//for {
	//	select {
	//	case ch, ok := <-c:
	//		if !ok {
	//			break
	//		} else {
	//			fmt.Println("c value :", ch)
	//		}

	//	default:
	//		break
	//	}

	//}
	case "3":
		mf := oop.MyFile{}
		sf := mf.InitMyFile()
		fmt.Printf("%v\n", sf)
	case "0":
		os.Exit(0)

	default:
		fmt.Println("input error,please try again!")

		reader := bufio.NewReader(os.Stdin)
		tmp, _, _ := reader.ReadLine()

		chooseMethod(string(tmp))
	}
}
