package main

import (
	"fmt"
)

func isPalindrome(input string) bool {
    for i := 0; i < len(input)/2; i++ {
        if input[i] != input[len(input)-i-1] {
            return false
        }
    }
    return true
}


func main() {
    var my_word string
    fmt.Println("Enter a word\n")
    fmt.Scanf("%s", my_word)
    if isPalindrome(my_word)==true {
        fmt.Println("Yes this word is palindrom")
    } else {
        fmt.Println("No this word is not palindrom")
    }
	
}
