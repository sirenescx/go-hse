package main

import (
	"fmt"
	"math"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
  	   if n % i == 0 {
           fmt.Println(i)
           return
       }
    }
    fmt.Println(n)
}
