package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"unicode"
)

func main () {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Printf("unable to read file")
  }
  defer file.Close()
 
  var password []int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text := scanner.Text()

    number := extractDigit(text)

    text = reverse(text)
  
    number += extractDigit(text)

    i, _ := strconv.Atoi(number)
    password= append(password, i)
  }
  
  var sum int
  for _, value := range password {
    sum += value 
  }

  fmt.Println(sum)
  fmt.Println("---------------------------------")
  partTwo()
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func extractDigit(text string) string {
    for _, char := range text {
      if unicode.IsDigit(char) {
        return string(char)
      }
    }
  return ""
}

