package main

import (
	"fmt"
)



func main() {
	var number int 
    fmt.Println("Enter a positive number\n")
    fmt.Scanf("%d", &number)
    var sum_odd, sum_even int 
    for i:=1; i <=number; i++ {
        if i%2!=0 {
            sum_odd+=i
        }
    }
    for i:=2; i<=number; i++ {
        if i%2==0 {
            sum_even+=i

        }
    }
    fmt.Println("Sum even numbers: ", sum_even)
    fmt.Println("Sum odd numbers: ", sum_odd)
}