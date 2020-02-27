package main

import "fmt"

func Calcular(a int, b int) (result int) {

    result = a + b
    return result
}

func main() {

    fmt.Println("Inicio do teste")
    result := Calcular(5, 5)
    fmt.Println(result)
}  