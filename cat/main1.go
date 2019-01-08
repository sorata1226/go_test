package main

import (	
	"fmt"
	"io/ioutil"
	"os"
)


func main() {

	for_, f := range os,Args[1:]{
		//fmt.Printf("%s\n")
		ReadFile(f)
	}

	ReadFile("main.go")
	

}

func ReadFile(filename string) {
	// ファイルを開く
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ファイルの内容全てを読む
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(string(data))

	// ファイルを閉じる
	file.Close()
}
