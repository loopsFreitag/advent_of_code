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
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

type colorAmount struct {
	color string
	amount int
}

func main() {
	var validLineSum int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Unable to read file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for lineNum := 1; scanner.Scan(); lineNum++ {
		colorAmounts := parseLine(scanner.Text())

		isValid := true
		for _, entry := range colorAmounts {
			switch entry.color {
			case "red":
				if entry.amount > maxRed {
					isValid = false
					break
				}
			case "blue":
				if entry.amount > maxBlue {
					isValid = false
					break
				}
			case "green":
				if entry.amount > maxGreen {
					isValid = false
					break
				}
			}
		}

		if isValid {
			validLineSum += lineNum
		}
	}

	fmt.Println(validLineSum)
	Part2()
}

func parseLine(text string) []colorAmount {
	var result []colorAmount

	colonIndex := strings.Index(text, ":")
	if colonIndex != -1 {
		text = text[colonIndex+1:]
	}

	re := regexp.MustCompile(`(\d+)\s([a-zA-Z]+)`)
	matches := re.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		amount, _ := strconv.Atoi(match[1])
		color := match[2]
		result = append(result, colorAmount{color: color, amount: amount})
	}

	return result
}

