package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    // Declare second integer, double, and String variables.
    var first uint64
    var second float64
    var third string
    
    // Read and save an integer, double, and String to your variables.
    fmt.Scan(&first)
    fmt.Scan(&second)
    for scanner.Scan() {
        third = scanner.Text()
    }
    
    // Print the sum of both integer variables on a new line.
    sum := i + first
    fmt.Println(sum)
    
    // Print the sum of the double variables on a new line.
    sumDub := second + d
    fmt.Printf("%.1f \n", sumDub)
    
    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.
    fmt.Println(s + "" + third)
}