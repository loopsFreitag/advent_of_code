package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
  red = 12
  green = 13
  blue = 14
)

type amout struct {
  color string
  amout int
}

func main () {
  var total int

  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Printf("unable to read file")
  }
  defer file.Close()
 
  scanner := bufio.NewScanner(file)

  for i := 1; scanner.Scan(); i++ {
    text := formatString(scanner.Text())
    out:
     for x ,j := range text {
      switch j.color {
      case "red":
        if j.amout > red {
          break out
        }
      case "blue":
        if j.amout > blue {
          break out
        }
      case "green":
        if j.amout > green {
          break out
        }
      }
      if x+1 == len(text) {
        // fmt.Println(len(text))
        // fmt.Println(x+1)
        total = total + i
      }
    }
  } 
  fmt.Println(total)
}

func formatString(text string) []amout {
  var line []amout
	colonIndex := strings.Index(text, ":")
   if colonIndex != -1 {
		text = text[colonIndex+1:]
	}

  re := regexp.MustCompile(`(\d+)\s([a-zA-Z]+)`)

	formatted := re.ReplaceAllString(text, `"$1","$2"`)

	formatted = strings.ReplaceAll(formatted, ";", ",")
	split := strings.Split(formatted, ",")

	for i := range split {
		split[i] = strings.TrimSpace(split[i])
	}

	for i := 0; i < len(split); i += 2 {
    var amoutLine amout
    v, _ := strconv.Atoi(strings.Trim(split[i], `"`))
    amoutLine.amout = v
    amoutLine.color = strings.Trim(split[i+1], `"`) 

    line = append(line, amoutLine)
	}
  return line 
}
