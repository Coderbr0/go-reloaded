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
	//num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //for (low, number); (up, number); (cap, number)
	fmt.Printf("\n")
	var newWords []string
	for i, word := range strArr { //this can also be written as i := range strArr {}; value can be omitted but not vice versa (index has to be omitted with underscore _ )
								  //without word value defined we would use strArr[i] in if statements; word is equivalent to strArr[i] 
		if word == "(cap)" {
			newWords[len(newWords)-1] = strings.Title(strArr[i-1]) // result is newWords[len(newWords)-1] as we want to capitalize word of previous index and assign to new string (newWords); we don't want to iterate over and modify current string (strArr) as this would be bad practice
			continue                                               // continue skips "(cap)" string and moves on to the next word, therefore omitting "(cap)"
		} else if word == "(up)" {
			newWords[len(newWords)-1] = strings.ToUpper(strArr[i-1])  //newWords[i-1] wouldn't take into consideration shortening of total elements in the slice after each operation e.g. when "(cap)" is omitted in previous operation
			continue												  //so we use newWords[len(newWords)-1]
		} else if word == "(low)" {
			newWords[len(newWords)-1] = strings.ToLower(strArr[i-1])
			continue
		} else if word == "(cap," && strArr[i+1] == "6)" {
			newWords[len(newWords)-6] = strings.Title(strArr[i-6])  //[i-1] for previous index; [i-6] for previous 6 in index
			newWords[len(newWords)-5] = strings.Title(strArr[i-5])
			newWords[len(newWords)-4] = strings.Title(strArr[i-4])
			newWords[len(newWords)-3] = strings.Title(strArr[i-3])
			newWords[len(newWords)-2] = strings.Title(strArr[i-2])
			newWords[len(newWords)-1] = strings.Title(strArr[i-1])
			continue
		} else if word == "(low," && strArr[i+1] == "3)" {
			newWords[len(newWords)-3] = strings.ToLower(strArr[i-3])
			newWords[len(newWords)-2] = strings.ToLower(strArr[i-2])
			newWords[len(newWords)-1] = strings.ToLower(strArr[i-1])
			continue
		} else if len(word) > 1 && word[1] == ')' {
			continue
		} else if word == "," {
			newWords[len(newWords)-1] += ","	//concatenates (adds) the comma to end of string; can also be written as newWords[len(newWords)-1] = newWords[len(newWords)-1] + ","
			continue							//continue skips the extra "," that is read by the computer as a separate word due to the space between text and "," and moves on to the next word, therefore omitting the extra ","
		} else if word == "(up," {
			numWords := strArr[i+1][0]-48
			wordsToModify := newWords[len(newWords)-int(numWords):] //alternative is strArr[i-int(numWords):i]    //when we have square brackets [] we are telling the compiler in Golang to look at what is before the square brackets and get the element at the stated index e.g. [i], [i-1] etc
			newWords = newWords[:len(newWords)-int(numWords)] //to delete the words not required
			newWords = append(newWords, strings.Split(strings.ToUpper(strings.Join(wordsToModify, " ")), " ")...) //strings.Join to turn slice into string; strings.Split to split string into elements (slice) to append later one by one which is done for consistency
			continue 	//int(numWords):i we cast to int as without it we have bytes. i is at the end as we want uppercase up to the index and not after so that we don't uppercase other words
		}
		
		// else if word == "." {
		// 	newWords[len(newWords)-1] += "."
		// 	continue
		// } else if word == "!" {
		// 	newWords[len(newWords)-1] += "!"
		// 	continue
		// } else if word == "?" {
		// 	newWords[len(newWords)-1] += "?"
		// 	continue
		// } else if word == ":" {
		// 	newWords[len(newWords)-1] += ":"
		// 	continue
		// } else if word == ";" {
		// 	newWords[len(newWords)-1] += ";"
		// 	continue
		// }
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
