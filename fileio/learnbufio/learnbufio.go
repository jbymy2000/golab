package learnbufio

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFile(){
	f, err := os.Open("./src/day1/file_read/1.txt")
	if err != nil{
		fmt.Println(err)
		os.Exit(2)
	}
	defer f.Close()
	bufReader := bufio.NewReader(f)
	//var buf [256]byte
	var i = 0
	for{
		i++
		fmt.Println(i)
		line,err := bufReader.ReadBytes('\n')
		fmt.Println(string(line))
		if err == io.EOF {
			fmt.Println("read the file finished")
			break
		}
		if err != nil{
			fmt.Println(err)
			os.Exit(2)
		}

	}
}

func ReadFileUseScanner() error{
	f, err := os.Open("./test.csv")
	if err != nil{
		fmt.Println(err)
		os.Exit(2)
	}
	defer f.Close()
	bufReader := bufio.NewScanner(f)
	for bufReader.Scan(){
		println(bufReader.Text())
	}

	return bufReader.Err()
}
