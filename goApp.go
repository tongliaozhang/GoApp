package main

import (
	"GoApp/db"
	"GoApp/dirMonitor"
	"GoApp/fileDeal"
	"GoApp/oop"
	"bufio"
	//"bytes"
	//"encoding/json"
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
	fmt.Println("*  4 FileDealToDB           *")
	fmt.Println("*  5 DirMonitor             *")
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
	case "3":
		mf := oop.MyFile{}
		sf := mf.InitMyFile()
		fmt.Print("%v\n", sf)
	case "4":
		fileDb := fileDeal.DataFile{}
		fileInfo := fileDb.FileInfoInit("textFile.csv")

		fmt.Printf("%v\n", fileInfo)

		//buffer := bytes.NewBuffer(nil)
		//encoder := json.NewEncoder(buffer)
		//encoder.Encode(fileInfo)

		//fmt.Printf("%v\n", buffer)

		db.InsertMysql(&fileInfo, "insert into go_insert(C1,C2,C3,C4)VALUES(?,?,?,?)")

	case "5":
		fmt.Println("Please input Dir Path:")
		reader := bufio.NewReader(os.Stdin)
		input, _, _ := reader.ReadLine()
		dirMontor(string(input))

	case "6":
		//dirMonitor.PathWalk("F://GoSpace//src//GoApp")
	case "0":
		os.Exit(0)
	default:
		fmt.Println("input error,please try again!")

		reader := bufio.NewReader(os.Stdin)
		tmp, _, _ := reader.ReadLine()

		chooseMethod(string(tmp))
	}
}

func dirMontor(dirName string) {
	ch := make(chan dirMonitor.DirInfo)
	info1 := dirMonitor.DirInfo{}

	begin := info1.PathWalk(dirName)

	go info1.DirMonitor(dirName, ch)

	for v := range ch {

		if dirMonitor.DirWaklCompare(v, begin) {
			//fmt.Println("==============================")
			//fmt.Println("Dir not change......")
		} else {
			begin = info1.PathWalk(dirName)
			fmt.Println("Dir is change......")
			fmt.Println("==============================")
		}

	}
}
