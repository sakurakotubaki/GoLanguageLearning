package main

import "fmt"

func main() {
    var num int = 10

    fmt.Println("numのアドレス:", &num)
    fmt.Println("numの値:", num)

    var ptr *int = &num

    fmt.Println("ptrのアドレス:", &ptr)
    fmt.Println("ptrの値:", ptr)
    fmt.Println(num)

    fmt.Println("numの新しい値:", num)
}