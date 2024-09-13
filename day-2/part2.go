package main

import (
	"bufio"
	"fmt"
  "os"
)

func Part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Unable to read file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var powerSum, powerLine int
	for scanner.Scan() {
		colorAmounts := parseLine(scanner.Text())
    var maxRed, maxBlue, maxGreen int
		for _, entry := range colorAmounts {
			switch entry.color {
			case "red":
				if entry.amount > maxRed {
          maxRed = entry.amount
				}
			case "blue":
				if entry.amount > maxBlue {
          maxBlue = entry.amount
				}
			case "green":
				if entry.amount > maxGreen {
          maxGreen = entry.amount
				}
			}
		}

    powerLine = maxRed * maxGreen * maxBlue
    powerSum = powerSum + powerLine
	}

  fmt.Println(powerSum)
}
