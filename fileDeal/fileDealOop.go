package fileDeal

import (
	"bufio"
	"github.com/axgle/mahonia"
	"io"
	"os"
	"strings"
)

type DataFile struct {
	Name string
	Type string
	Info map[string]interface{}
	Data []string
}

func (this DataFile) FileInfoInit(fileName string) (df DataFile) {
	f, _ := os.Open(fileName)
	defer f.Close()

	fnames := strings.Split(fileName, ".")
	df.Name = fnames[0]
	df.Type = fnames[1]

	//mahonia 处理中文
	decoder := mahonia.NewDecoder("GBK")

	reader := bufio.NewReader(decoder.NewReader(f))
	data := make([]string, 0)
	for {
		line, isPrefix, err := reader.ReadLine()

		if isPrefix {

		} else {
			data = append(data, string(line))
		}

		if err == io.EOF {
			break
		}
	}

	df.Data = data

	return
}
