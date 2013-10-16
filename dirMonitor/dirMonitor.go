// dirMonitor project dirMonitor.go
package dirMonitor

import (
	//"bytes"
	"fmt"
	"os"
	"path/filepath"
	//"strconv"
	"time"
)

type MyFile struct {
	FileName string
	Size     int64
	FilePath string
}

/*
Dir Info struct
*/
type DirInfo struct {
	Name string //文件夹名称
	//Path     string    //路径
	FileNum  int      //文件数量
	FileList []MyFile //文件列表
	DirNum   int      //文件夹数量
	DirList  []MyFile //文件夹信息
}

/*
Dir Info init
*/
//func (this DirInfo) DirInfoScan(dirName string) (dirInfo DirInfo) {

//	f, _ := os.Open(dirName)
//	defer f.Close()

//	dirInfo.Name = dirName

//	infoD, _ := f.Readdirnames(0)
//	dirNum := 0
//	fileNum := 0
//	fileList := make([]MyFile, 0)
//	dirList := make([]DirInfo, 0)

//	for _, v := range infoD {
//		sb := bytes.Buffer{}
//		sb.WriteString(dirName)
//		sb.WriteString("//")
//		sb.WriteString(v)
//		cpath := sb.String()

//		cf, _ := os.Stat(cpath)
//		if cf.IsDir() {
//			cdir := DirInfo{}

//			chirdDirInfo := cdir.DirInfoScan(cpath)

//			dirList = append(dirList, chirdDirInfo)
//			dirNum++
//		} else {
//			mf := MyFile{}
//			mf.FileName = v

//			fileList = append(fileList, mf)
//			fileNum++
//		}
//	}

//	dirInfo.FileNum = fileNum
//	dirInfo.DirNum = dirNum
//	dirInfo.FileList = fileList
//	dirInfo.DirList = dirList

//	return
//}

//func (this DirInfo) DirMonitor(dirName string, ch chan DirInfo) {
//	for {
//		ch <- this.DirInfoScan(dirName)
//		time.Sleep(5 * time.Second)
//	}
//}

//func DirInfoCompare(dirA DirInfo, dirB DirInfo) bool {
//	if dirA.Name != dirB.Name {
//		fmt.Println("DirName differenct")
//		return false
//	}
//	if dirA.FileNum != dirB.FileNum {
//		fmt.Println("FileNum differenct")
//		return false
//	} else {
//		//比对文件
//		fileFlag := true

//		for i := 0; i < len(dirA.FileList); i++ {
//			afile := dirA.FileList[i]
//			tmpFlag := true
//			for j := 0; j < len(dirB.FileList); j++ {
//				bfile := dirB.FileList[j]

//				if afile.FileName != bfile.FileName {
//					tmpFlag = false
//				} else {
//					tmpFlag = true
//					break
//				}
//			}
//			if tmpFlag == false {
//				fileFlag = false
//				break
//			}
//		}

//		if fileFlag {
//			return fileFlag
//		} else {
//			fmt.Println("File differenct")
//			return fileFlag
//		}

//	}
//	if dirA.DirNum != dirB.DirNum {
//		fmt.Println("ChiidDirName differenct")
//		return false
//	} else {
//		//循环比较文件夹
//		dirFlag := true

//		for i := 0; i < len(dirA.DirList); i++ {
//			adir := dirA.DirList[i]
//			tmpFlag := true
//			for j := 0; j < len(dirB.DirList); j++ {
//				bdir := dirB.DirList[j]

//				if DirInfoCompare(adir, bdir) {
//					tmpFlag = false
//				} else {
//					tmpFlag = true
//					break
//				}
//			}
//			if tmpFlag == false {
//				dirFlag = false
//				break
//			}
//		}

//		if dirFlag {
//			return dirFlag
//		} else {
//			fmt.Println("Dir differenct")
//			return dirFlag
//		}
//	}
//	return true
//}

func (this DirInfo) DirMonitor(dirName string, ch chan DirInfo) {
	for {
		ch <- this.PathWalk(dirName)
		time.Sleep(10 * time.Second)
	}
}

const (
	CREATE_NEW = "New Create"
	DELETE     = "Delete"
)

func (this DirInfo) PathWalk(path string) (res DirInfo) {
	dirNum := 0
	fileNum := 0
	fileList := make([]MyFile, 0)
	dirList := make([]MyFile, 0)

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		//p, _ := filepath.Abs(info.Name())
		//fmt.Printf("%s %s\n", filepath.Dir(p), info.Name())

		if info.IsDir() {
			dirNum++
			mf := MyFile{}
			mf.FileName = info.Name()
			mf.FilePath = path
			dirList = append(dirList, mf)
		} else {
			mf := MyFile{}
			mf.FileName = info.Name()
			mf.Size = info.Size()
			mf.FilePath = path
			fileNum++
			fileList = append(fileList, mf)
		}
		return nil
	})

	res.Name = path
	res.DirNum = dirNum
	res.DirList = dirList
	res.FileNum = fileNum
	res.FileList = fileList
	return
}

func DirWaklCompare(dirA DirInfo, dirB DirInfo) (cmpFlag bool) {
	cmpFlag = true
	//比对文件数量
	if dirA.FileNum != dirB.FileNum {
		fmt.Printf("FileNum differenct: %d --> %d \n", dirB.FileNum, dirA.FileNum)
		cmpFlag = false
	}

	//比对文件夹数量
	if dirA.DirNum != dirB.DirNum {
		fmt.Printf("Dir number differenct: %d --> %d \n", dirB.DirNum, dirA.DirNum)
		cmpFlag = false
	}

	//比对文件
	cmpFileFlag := true
	if dirA.FileNum > dirB.FileNum {
		//目录文件增加了
		FileNameCompare(dirA, dirB, CREATE_NEW)
	} else if dirA.FileNum < dirB.FileNum {
		//目录文件少了
		FileNameCompare(dirB, dirA, DELETE)
	} else {
		//文件数一样 但是文件可能变化了
		cmpFileFlag = FileNameCompare(dirA, dirB, CREATE_NEW)
		cmpFileFlag = FileNameCompare(dirB, dirA, DELETE)
	}

	cmpDirFlag := true
	//比对文件夹名称
	if dirA.DirNum > dirB.DirNum {
		//目录增加了
		DirNameCompare(dirA, dirB, CREATE_NEW)
	} else if dirA.DirNum < dirB.DirNum {
		//目录少了
		DirNameCompare(dirB, dirA, DELETE)
	} else {
		//目录数一样 但是目录名可能变化了
		cmpDirFlag = DirNameCompare(dirA, dirB, CREATE_NEW)
		cmpDirFlag = DirNameCompare(dirB, dirA, DELETE)
	}

	if cmpDirFlag == false || cmpFileFlag == false {
		cmpFlag = false
	}

	return cmpFlag
}

func FileNameCompare(dirA DirInfo, dirB DirInfo, changeInfo string) (cmpFlag bool) {
	cmpFlag = true
	fileFlag := true
	fileArray := make([][2]MyFile, 0)
	sizeFlag := false //没改变
	fileExit := false //文件不匹配

	for i := 0; i < len(dirA.FileList); i++ {
		afile := dirA.FileList[i]
		for j := 0; j < len(dirB.FileList); j++ {
			bfile := dirB.FileList[j]

			if afile.FileName == bfile.FileName && afile.FilePath == bfile.FilePath {
				fileExit = true
				if afile.Size != bfile.Size {
					fileArray = append(fileArray, [2]MyFile{afile, bfile})
					sizeFlag = true
					break
				}
			}

			if j == len(dirB.FileList)-1 {
				if fileExit == false {
					//fmt.Println(afile.FileName + " file is " + changeInfo)
					fmt.Printf("%s : %s ,path :(%s)\n", changeInfo, afile.FileName, afile.FilePath)
				}
			}
		}
		if fileExit == false || (fileExit && sizeFlag == true) {
			fileFlag = false
		}

		fileExit = false
	}

	if fileFlag == false {
		if sizeFlag {

			for i := 0; i < len(fileArray); i++ {
				item := [2]MyFile(fileArray[i])
				fmt.Printf("File differenct: Size %s %d --> %d \n", item[0].FileName, item[0].Size, item[1].Size)
			}
		}

		cmpFlag = false
	}

	return
}

func DirNameCompare(dirA DirInfo, dirB DirInfo, changeInfo string) (cmpFlag bool) {
	cmpFlag = true
	dirFlag := true
	dirArray := make([][1]MyFile, 0)
	dirExit := false //文件不匹配

	for i := 0; i < len(dirA.DirList); i++ {
		dira := dirA.DirList[i]
		for j := 0; j < len(dirB.DirList); j++ {
			dirb := dirB.DirList[j]

			if dira.FileName == dirb.FileName && dira.FilePath == dirb.FilePath {
				dirExit = true
				break
			}

			if j == len(dirB.DirList)-1 {
				if dirExit == false {
					dirArray = append(dirArray, [1]MyFile{dira})
				}
			}
		}
		if dirExit == false {
			dirFlag = false
		}

		dirExit = false
	}

	if dirFlag == false {
		for i := 0; i < len(dirArray); i++ {
			item := [1]MyFile(dirArray[i])
			//fmt.Println("Dir differenct: " + item[i] + " is " + changeInfo)
			fmt.Printf("%s : %s , ,path :(%s)\n", changeInfo, item[i].FileName, item[i].FilePath)
		}
		cmpFlag = false
	}

	return
}
