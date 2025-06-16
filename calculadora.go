package main

import (
    "fmt"
)

func main() {
    var num1, num2 float64
    var op string

    fmt.Print("Digite o primeiro número: ")
    fmt.Scanln(&num1)

    fmt.Print("Digite o operador (+, -, *, /): ")
    fmt.Scanln(&op)

    fmt.Print("Digite o segundo número: ")
    fmt.Scanln(&num2)

    switch op {
    case "+":
        fmt.Printf("Resultado: %.2f\n", num1+num2)
    case "-":
        fmt.Printf("Resultado: %.2f\n", num1-num2)
    case "*":
        fmt.Printf("Resultado: %.2f\n", num1*num2)
    case "/":
        if num2 != 0 {
            fmt.Printf("Resultado: %.2f\n", num1/num2)
        } else {
            fmt.Println("Erro: Divisão por zero!")
        }
    default:
        fmt.Println("Operador inválido!")
    }
}
