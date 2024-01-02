package main

import "fmt"

func main() {
    x := "apple"
    switch x {
    case "apple", "banana", "peach":
        fmt.Println("x is a fruit")
    case "carrot", "celery", "beet":
        fmt.Println("x is a vegetable")
    default:
        fmt.Println("x is not a fruit or a vegetable")
    }
}