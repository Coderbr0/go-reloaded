package main

import (
	"bufio"
	"fmt"
	"os"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func ReadFile() []string {
	var fileInput []string
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Invalid Input")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fileInput = append(fileInput, scanner.Text())
	}
	return fileInput
}

func main() {
	fmt.Println(ReadFile())
	strArr := ReadFile()
	strArr = remove(strArr, 1)
	strArr = remove(strArr, 1)
	strArr = remove(strArr, 1)
	fmt.Println(strArr)
}
