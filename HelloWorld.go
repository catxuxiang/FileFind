package main

import (
	"container/list"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CheckErr(err error) {
	if nil != err {
		panic(err)
	}
}

func GetFullPath(path string) string {
	absolutePath, _ := filepath.Abs(path)
	return absolutePath
}

func PrintFilesName(pre string, path string, listStr *list.List) {
	fullPath := GetFullPath(path)

	dir, _ := ioutil.ReadDir(fullPath)
	for _, fi := range dir {
		PthSep := string(os.PathSeparator)
		if fi.IsDir() { // 忽略目录
			PrintFilesName(pre+PthSep+fi.Name(), path+PthSep+fi.Name(), listStr)
		}
		//fmt.Println(fi.Name())
		if strings.HasSuffix(strings.ToUpper(fi.Name()), ".H") {
			//if strings.HasSuffix(strings.ToUpper(fi.Name()), ".CPP") || strings.HasSuffix(strings.ToUpper(fi.Name()), ".CC") {
			listStr.PushBack(pre + PthSep + fi.Name())
		}
	}
}

func PrintDirName(pre string, path string, listStr *list.List) {
	PthSep := string(os.PathSeparator)
	listStr.PushBack(pre)

	fullPath := GetFullPath(path)
	dir, _ := ioutil.ReadDir(fullPath)
	for _, fi := range dir {
		if fi.IsDir() {
			PrintDirName(pre+PthSep+fi.Name(), path+PthSep+fi.Name(), listStr)
		}
	}
}

func ConvertToSlice(listStr1 *list.List) []string {
	sli := []string{}
	for el := listStr1.Front(); nil != el; el = el.Next() {
		sli = append(sli, el.Value.(string))
	}

	return sli
}

func OutputFilesName(listStr *list.List) {
	files := ConvertToSlice(listStr)
	//sort.StringSlice(files).Sort()// sort

	f, err := os.Create("Output.txt")
	//f, err := os.OpenFile(outputFileName, os.O_APPEND | os.O_CREATE, os.ModeAppend)
	CheckErr(err)
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(f)

	length := len(files)
	for i := 0; i < length; i++ {
		writer.Write([]string{files[i]})
	}

	writer.Flush()
}

func main() {
	path := "/home/sky/cocos2d-x-3.8.1/tests/SocketClient/Classes"

	listStr := list.New()
	//PrintDirName("  Classes", path, listStr)
	PrintFilesName("  Classes", path, listStr)

	for i := listStr.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
