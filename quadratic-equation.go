package main
import . "fmt"
import "math"

func main() {
  var a, b, c float64
  Print("a = ")
  Scan(&a)
  Print("b = ")
  Scan(&b)
  Print("c = ")
  Scan(&c)
  Printf("Equation is: (%.1f) x^2 + (%.1f) x + (%.1f) = 0\n", a, b, c)
  equation(a, b, c)
}

func equation(a, b, c float64) {
  D := b * b - 4 * a * c
  square := math.Sqrt(D)

  switch {
  case D < 0:
    Println("There are 0 roots")
  case D == 0:
    Println("There are 1 roots")
    Println("x1 =", -b / ( 2 * a ))
  case D > 0:
    Println("There are 2 roots")
    Println("x1 =", (-b + square) / ( 2 * a ))
    Println("x2 =", (-b - square) / ( 2 * a ))
  }
}
