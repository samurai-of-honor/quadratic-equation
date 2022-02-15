package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var a, b, c float64
	var argumentIndex int

	checkArg := os.Args
	if len(checkArg) > 1 {
		file, err := os.Open(checkArg[1])
		if err != nil {
			panic("file does not exist")
		}
		defer file.Close()

		str := bufio.NewScanner(file)
		for str.Scan() {
			number, err := strconv.ParseFloat(str.Text(), 64)
			if err != nil {
				panic("invalid file format")
			}
			if argumentIndex == 0 && number == 0 {
				panic("Error. a cannot be 0")
			}
			switch {
			case argumentIndex == 0:
				a = number
			case argumentIndex == 1:
				b = number
			case argumentIndex == 2:
				c = number
			}
			argumentIndex++
		}
	} else {
		a = inputArg("a")
		b = inputArg("b")
		c = inputArg("c")
	}
	Printf("Equation is: (%.1f) x^2 + (%.1f) x + (%.1f) = 0\n", a, b, c)
	equation(a, b, c)

	Println("Enter 'q' to quit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		exit := scanner.Text()
		if exit == "q" {
			break
		} else {
			Println("Enter 'q' to quit")
		}
	}
}

func inputArg(arg string) float64 {
	for {
		switch {
		case arg == "a":
			Print("a = ")
		case arg == "b":
			Print("b = ")
		case arg == "c":
			Print("c = ")
		}
		str, _, readerErr := bufio.NewReader(os.Stdin).ReadLine()
		if readerErr != nil {
			panic(readerErr)
		}
		result, convertErr := strconv.ParseFloat(string(str), 64)
		if convertErr != nil {
			Printf("Error. Expected a valid real number, got %s instead\n", str)
		} else {
			return result
		}
	}
}

func equation(a, b, c float64) {
	D := b*b - 4*a*c
	square := math.Sqrt(D)

	switch {
	case D < 0:
		Println("There are 0 roots")
	case D == 0:
		Println("There are 1 roots")
		Println("x1 =", -b/(2*a))
	case D > 0:
		Println("There are 2 roots")
		Println("x1 =", (-b+square)/(2*a))
		Println("x2 =", (-b-square)/(2*a))
	}
}
