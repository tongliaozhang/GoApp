package fileDeal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadFileAndPrint() bool {
	defer func() {
		recover()
	}()

	fmt.Print("Please input fileName:")

	stdRd := bufio.NewReader(os.Stdin)

	fName, _, _ := stdRd.ReadLine()

	file, err := os.Open(string(fName))
	defer file.Close()

	if err != nil {
		fmt.Println("open file failed,please check !")
		panic("Test")
		return false
	}

	readBuffer := bufio.NewReaderSize(file, 4096)

	for {
		line, isPerfix, err := readBuffer.ReadLine()

		if isPerfix {
			fmt.Println(string(line))
		} else if len(line) > 0 {
			fmt.Println(string(line))
		}

		if err == io.EOF {
			break
		}
	}

	return true

}

func checkErr(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}
