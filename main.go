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
	fmt.Printf("\n")
	var newWords []string
	for _, word := range strArr {
		if word == "[" || word == "(cap)" || word == "(up)" || word == "(cap," || word == "6)" || word == "(low," || word == "3)" || word == "]" {
			continue
		}
		newWords = append(newWords, word)
	}
	fmt.Println(newWords)
}

//func ToUpper(str string) string

//func ToLower(str string) string

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
