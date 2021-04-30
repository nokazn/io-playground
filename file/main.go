package main

import (
	"fmt"
	"os"
)

func openOrCreate(file string, writable bool) (*os.File, error) {
	var f *os.File
	var err error
	if writable {
		f, err = os.Create(file)
	} else {
		f, err = os.Open(file)
	}

	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to open file:", file)
		return nil, err
	}
	return f, nil
}

func read(file *os.File) (string, error) {
	b := make([]byte, 1024)
	count, err := file.Read(b)
	if err != nil && err.Error() != "EOF" {
		fmt.Println(err)
		fmt.Println("Failed to read file:", file.Name())
		return "", err
	}
	fmt.Println(count, "bytes in", file.Name()+":", b[:count])
	fmt.Println("content:", string(b[:count]))
	return string(b[:count]), nil
}

func write(file *os.File, content string) (int, error) {
	b := []byte(content)
	count, err := file.Write(b)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to write file:", file.Name())
	} else {
		fmt.Println(count, "bytes in", file.Name(), b[:count])
		fmt.Println("content:", string(b[:count]))
	}
	return count, err
}

func printDiv() {
	fmt.Println("__________________________________________________\n ")
}

func main() {
	f1, _ := openOrCreate("./file/1.txt", false)
	defer f1.Close()
	read(f1)
	printDiv()

	f2, _ := openOrCreate("./file/2.txt", true)
	defer f2.Close()
	read(f2)
	printDiv()

	f3, _ := openOrCreate("./file/3.txt", true)
	defer f3.Close()
	write(f3, "hey!")
	// "hey!" cannot read
	read(f3)
	f4, _ := openOrCreate("./file/3.txt", false)
	defer f4.close()
	// "hey!" can read
	read(f4)
}
