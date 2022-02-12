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

	checkArg := os.Args[1]
	if checkArg == "args.txt" {
		file, err := os.Open("args.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		str := bufio.NewScanner(file)
		for str.Scan() {
			number, err := strconv.ParseFloat(str.Text(), 64)
			if err != nil {
				panic(err)
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
		Print("a = ")
		Scan(&a)
		Print("b = ")
		Scan(&b)
		Print("c = ")
		Scan(&c)
	}

	Printf("Equation is: (%.1f) x^2 + (%.1f) x + (%.1f) = 0\n", a, b, c)
	equation(a, b, c)

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
