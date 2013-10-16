package oop

type MyFile struct {
	Id       int32
	FileName string
}

func (f MyFile) InitMyFile() MyFile {
	return MyFile{Id: 1, FileName: "zhangql.txt"}
}
