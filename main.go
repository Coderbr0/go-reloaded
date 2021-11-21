package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func ReadFile() []string {
	var fileInput []string
	file, err := os.Open(os.Args[1]) // We do this first to open the file
	if err != nil {
		fmt.Println("Invalid Input")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file) // scanner := bufio.NewScanner(file) results in scanner.Split(bufio.ScanLines)
	scanner.Split(bufio.ScanWords)    // So we write scanner.Split(bufio.ScanWords) in order to scan words as opposed to scan lines as elements
	for scanner.Scan() {
		fileInput = append(fileInput, scanner.Text()) // Scans text in file
	}
	return fileInput
}

func main() {
	fmt.Println(ReadFile()) // Calling ReadFile() function and printing as output the text in sample.txt
	strArr := ReadFile()
	fmt.Printf("\n")
	var newWords []string
	insideQuotes := true
	for i, word := range strArr { // This can also be written as i := range strArr {}; value can be omitted but not vice versa (index has to be omitted with underscore _ )
								  // Without word value defined we would use strArr[i] in if statements; word is equivalent to strArr[i] 
		if word == "(cap)" {
			newWords[len(newWords)-1] = strings.Title(strArr[i-1]) // Result is newWords[len(newWords)-1] as we want to capitalize word of previous index and assign to new string (newWords); we don't want to iterate over and modify current string (strArr) as this would be bad practice
			continue                                               // continue skips "(cap)" string and moves on to the next word, therefore omitting "(cap)"
		} else if word == "(up)" {
			newWords[len(newWords)-1] = strings.ToUpper(strArr[i-1])  // newWords[i-1] wouldn't take into consideration shortening of total elements in the slice after each operation e.g. when "(cap)" is omitted in previous operation
			continue												  // So we use newWords[len(newWords)-1]
		} else if word == "(low)" {
			newWords[len(newWords)-1] = strings.ToLower(strArr[i-1])  // newWords[len(newWords)-1] represents the last element of the new slice of words
			continue
		} else if word == "(cap," {									// Alternative with specified number: else if word == "(cap," && strArr[i+1] == "6)" { 
			numWords := strArr[i+1][0]-48							// newWords[len(newWords)-6] = strings.Title(strArr[i-6]) 
			wordsToModify := newWords[len(newWords)-int(numWords):] // newWords[len(newWords)-5] = strings.Title(strArr[i-5]) etc 
			newWords = newWords[:len(newWords)-int(numWords)] 		// continue } ; [i-1] for previous index; [i-6] for previous 6 in index
			newWords = append(newWords, strings.Split(strings.Title(strings.Join(wordsToModify, " ")), " ")...) 
			continue 	
		} else if word == "(up," {																			
			numWords := strArr[i+1][0]-48	// strArr[i+1][0] is in bytes so we -48 to get the corresponding value as an integer; strArr[i+1][0] is an ascii value (Char) so we -48 to convert to decimal value (Dec as an integer); for example, if we want to change the last 2 words then Char = 2, Dec = 50, without -48 the last 50 words would change 																	
			wordsToModify := newWords[len(newWords)-int(numWords):] // Alternative is strArr[i-int(numWords):i]; i is at the end as we want uppercase up to the index and not after so that we don't uppercase other words  
			newWords = newWords[:len(newWords)-int(numWords)] // To delete the words not required; when we have square brackets [] we are telling the compiler in Golang to look at what is before the square brackets and get the element at the stated index e.g. [i], [i-1] etc
			newWords = append(newWords, strings.Split(strings.ToUpper(strings.Join(wordsToModify, " ")), " ")...) // strings.Join to turn slice into string; strings.Split to split string into elements (slice) to append later one by one which is done for consistency
			continue 	// int(numWords); we cast to int as without it we have bytes
		} else if word == "(low," {
			numWords := strArr[i+1][0]-48 
			wordsToModify := newWords[len(newWords)-int(numWords):]  
			newWords = newWords[:len(newWords)-int(numWords)]
			newWords = append(newWords, strings.Split(strings.ToLower(strings.Join(wordsToModify, " ")), " ")...)
			continue
		} else if len(word) > 1 && word[1] == ')' { // Specifying len(word) > 1 as some words may be only one character long so word[1] would be out of range; index starts from zero e.g. word[0]
			continue
		} else if word == "," {
			newWords[len(newWords)-1] += ","	// Concatenates (adds) the comma to end of string; can also be written as newWords[len(newWords)-1] = newWords[len(newWords)-1] + ","
			continue							// continue skips the extra "," that is read by the computer as a separate word due to the space between text and "," and moves on to the next word, therefore omitting the extra ","
		} else if word == "." {
			newWords[len(newWords)-1] += "."
			continue
		} else if word == "!" {
			newWords[len(newWords)-1] += "!"
			continue
		} else if word == "?" {
			newWords[len(newWords)-1] += "?"
			continue
		} else if word == ":" {
			newWords[len(newWords)-1] += ":"
			continue
		} else if word == ";" {
			newWords[len(newWords)-1] += ";"
			continue
		} else if word == "..." {
			newWords[len(newWords)-1] += "..."
			continue
		} else if word == "!?" {
			newWords[len(newWords)-1] += "!?"
			continue
		} else if word == "?!" {				// ?! is an alternative way of writing !?
			newWords[len(newWords)-1] += "?!"
			continue
		} else if i > 0 && strArr[i-1] == "'" && len(newWords[len(newWords)-1]) == 1 { // i > 0 as we can't have an index less than 0 to accommodate strArr[i-1] ("index out of range")
			newWords[len(newWords)-1] += word   // For first quotation mark in ' awesome '
			insideQuotes = false
			continue 							// continue used to avoid having the word twice e.g. awesome awesome; finding "'" at the previous index e.g. when the iteration reaches awesome; continue skips appending the current word e.g. the duplication of awesome
		} else if i > 0 && word[0] == '\'' && !insideQuotes {	//newWords[len(newWords)-1][0] means first character of the last element of the newWords slice; len(newWords) by default is set to zero when defined (var newWords []string) so i > 0 required for [len(newWords)-1] to avoid "index out of range" error
		    newWords[len(newWords)-1] += word	// For second quotation mark in ' awesome '
			insideQuotes = true					// insideQuotes set to true by default; we declare different boolean values (true and false) to carry out different operations for first and second quotation marks; !insideQuotes is the same as insideQuotes == false in this case
			continue							// '\'' single quotes is a rune and having backslash escapes it; boolean variable insideQuotes is required to work for strings with multiple words inside quotes e.g. hello ' awesome fish ' nimo
		} else if word == "a" || word == "A" {
			if strArr[i+1][0] == 'a' || strArr[i+1][0] == 'e' || strArr[i+1][0] == 'i' || strArr[i+1][0] == 'o' || strArr[i+1][0] == 'u' || strArr[i+1][0] == 'h' || strArr[i+1][0] == 'A' || strArr[i+1][0] == 'E' || strArr[i+1][0] == 'I' || strArr[i+1][0] == 'O' || strArr[i+1][0] == 'U' || strArr[i+1][0] == 'H' {
				changeLetter := word + "n"
				newWords[len(newWords)-1] += " " + changeLetter
				continue
			} 
		} else if word[0] == ',' {				// For the ",don't" in: Punctuation tests are ... kinda boring ,don't you think !?
			newWords[len(newWords)-1] += "," 
			word = word[1:] 					// Selecting "don't" in ",don't"; omitting word[0] 
		} else if word == "(hex)" {
			if string(strArr[i-1][0]) == "\"" {	// Alternative is strArr[i-1][0] == 34; Decimal value of " is 34; in this case we have casted strArr[i-1][0] to a string
				newWords[len(newWords)-1] = strArr[i-1][1:] // For: "1E (hex) files were added" (including quotation marks); added quotation marks for educational purposes
			}
			numFromHex, _ := strconv.ParseInt(newWords[len(newWords)-1], 16, 64) // ParseInt function returns two values (int64, error) so two variables required (numFromHex, _)
			newWords[len(newWords)-1] = strconv.Itoa(int(numFromHex)) // Casting numFromHex into int so that Itoa can return a string
			//newWords[len(newWords)-1] = "\"" + newWords[len(newWords)-1] // Adding first quotation mark back to the string after conversion; commented out newWords[len(newWords)-1] = "\"" + newWords[len(newWords)-1] to PASS test
			continue
		} else if word == "(bin)" {
			numFromBin, _ := (strconv.ParseInt(newWords[len(newWords)-1], 2, 64))
			newWords[len(newWords)-1] = strconv.Itoa(int(numFromBin))
			continue
		}
		newWords = append(newWords, word)		
	}
	fmt.Println(newWords)
	Error := os.WriteFile("result.txt", []byte(strings.Join(newWords, " ")), 0777)
	fmt.Println(Error)
}
/* Alternative way of doing the project:

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
*/