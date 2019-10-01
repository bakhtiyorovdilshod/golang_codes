package main

import (
	"fmt"
)

func main() {
    var number int 
    fmt.Println("Enter a positive number\n")
    fmt.Scanf("%d", &number)
    for i:=0; i <=number; i++ {
    if i % 15 == 0 {
        fmt.Println("FizzBuzz")
    }else if i % 3 == 0 {
        fmt.Println("Fizz")
    }else if i % 5 == 0 {
        fmt.Println("Buzz")
    }else {
        fmt.Println(i)
    }

    }
	
}



