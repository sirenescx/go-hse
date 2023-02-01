package main

import (
  "fmt"
  "math"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)
    squared_root := math.Sqrt(float64(n))
    for i := 1; i <= int(squared_root); i++ {
        if i * i == n {
            fmt.Println(i)
            return
        }
        for j := 1; j <= int(squared_root); j++ {
            if i * i + j * j == n {
                fmt.Println(i, j)
                return
            }
            for k := 1; k <= int(squared_root); k++ {
                if i * i + j * j + k * k == n {
                  fmt.Println(i, j, k)
                  return
                }
                for m := 1; m <= int(squared_root); m++ {
                    if i * i + j * j + k * k + m * m == n {
                      fmt.Println(i, j, k, m)
                      return
                    }
                }
            }
        }
    }
}
