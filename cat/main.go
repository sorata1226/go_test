package main

import (	
	"fmt"
	"io/ioutil"
	"os"
)


func main() {

	ReadFile("main.go")
	

}

func ReadFile(filename string) {
	// �t�@�C�����J��
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// �t�@�C���̓��e�S�Ă�ǂ�
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(string(data))

	// �t�@�C�������
	file.Close()
}
