package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile() []string {
	var fileInput []string
	file, err := os.Open(os.Args[1]) // We do this first to open the file
	if err != nil {
		fmt.Println("Invalid Input")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file) // scanner := bufio.NewScanner(file) results in scanner.Split(bufio.ScanLines)
	scanner.Split(bufio.ScanWords)    // so we write scanner.Split(bufio.ScanWords) in order to scan words as opposed to scan lines as elements
	for scanner.Scan() {
		fileInput = append(fileInput, scanner.Text()) // scans text in file
	}
	return fileInput
}

func main() {
	fmt.Println(ReadFile()) // calling ReadFile() function and printing as output the text in sample.txt
	strArr := ReadFile()
	fmt.Printf("\n")
	var newWords []string
	for i, word := range strArr {

		if word == "(cap)" {
			newWords[i-1] = strings.Title(strArr[i-1])
			continue
		} else if word == "(up)" {
			newWords[len(newWords)-1] = strings.ToUpper(strArr[i-1])
			continue
		} else if word == "(cap," && strArr[i+1] == "6)" { // word is equivalent to strArr[i]
			newWords[len(newWords)-6] = strings.Title(strArr[i-6])
			newWords[len(newWords)-5] = strings.Title(strArr[i-5])
			newWords[len(newWords)-4] = strings.Title(strArr[i-4])
			newWords[len(newWords)-3] = strings.Title(strArr[i-3])
			newWords[len(newWords)-2] = strings.Title(strArr[i-2])
			newWords[len(newWords)-1] = strings.Title(strArr[i-1])
			continue
		} else if word == "6)" {
			continue
		} else if word == "(low," && strArr[i+1] == "3)" {
			newWords[len(newWords)-3] = strings.ToLower(strArr[i-3])
			newWords[len(newWords)-2] = strings.ToLower(strArr[i-2])
			newWords[len(newWords)-1] = strings.ToLower(strArr[i-1])
			continue
		} else if word == "3)" {
			continue
		}
		newWords = append(newWords, word)
	}
	fmt.Println(newWords)
}

/*
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
*/

// func ToUpper(str string) string

// func ToLower(str string) string

/* hexadecimal to decimal
package main

import (
 "fmt"
 "math"
 "strconv"
 "strings"
)
func hexaNumberToInteger(hexaString string) string {
 // replace 0x or 0X with empty String
 numberStr := strings.Replace(hexaString, "0x", "", -1)
 numberStr = strings.Replace(numberStr, "0X", "", -1)
 return numberStr
}

func main() {
 var hexaNumber string
 fmt.Print("Enter Hexadecimal Number:")
 fmt.Scanln(&hexaNumber)
 output, err := strconv.ParseInt(hexaNumberToInteger(hexaNumber), 16, 64)
 if err != nil {
  fmt.Println(err)
  return
 }
 fmt.Printf("Output %d", output)
}
*/

/* binary to decimal
package main

import (
 "fmt"
 "strconv"
)

func main() {
 var binary string
 fmt.Print("Enter Binary Number:")
 fmt.Scanln(&binary)
 output, err := strconv.ParseInt(binary, 2, 64)
 if err != nil {
  fmt.Println(err)
  return
 }

 fmt.Printf("Output %d", output)

}
*/
/*
package main

import (
    "fmt"
    "strings"
)

func main() {
    res := strings.Title("this is a test sentence")
    fmt.Println(res)
}
*/
