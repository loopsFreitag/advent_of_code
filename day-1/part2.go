package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partTwo() {
  stringToNumber := map[string]string{
      "one":   "one1one",
      "two":   "two2two",
      "three": "three3three",
      "four":  "four4four",
      "five":  "five5five",
      "six":   "six6six",
      "seven": "seven7seven",
      "eight": "eight8eight",
      "nine":  "nine9nine",
  }

  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Printf("unable to read file")
  }
  defer file.Close()
 
  var password []int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text := scanner.Text()

    for j, i := range stringToNumber {
      text = strings.Replace(text, j, i, -1)
    }

    number := extractDigit(text)

    text = reverse(text)

    number += extractDigit(text)

    i, _ := strconv.Atoi(number)
    password = append(password, i)
  }
  
  var sum int
  for _, value := range password {
    sum += value 
  } 
  fmt.Println(sum)
}


