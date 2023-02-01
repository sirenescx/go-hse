package main

import (
  "fmt"
  "math"
)

func main() {
    x := make([]int16, 8)
    y := make([]int16, 8)
    var a, b int16
    for i := 0; i < 8; i++ {
        fmt.Scanln(&a, &b)
        x[i] = a
        y[i] = b
    }

    for i := 0; i < 8; i++ {
      for j := 0; j < 8; j++ {
        if i == j {
          continue
        }
        if math.Abs(float64(x[i] - x[j])) == math.Abs(float64(y[i] - y[j])) || x[i] == x[j] || y[i] == y[j] {
          fmt.Println("YES")
          return
        }
      }
    }

    fmt.Println("NO")
}


