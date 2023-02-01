package main

import (
	"fmt"
	"math"
)

func min4(a int, b int, c int, d int) int {
  return int(math.Min(math.Min(float64(a), float64(b)), math.Min(float64(c), float64(d))))
}

func main() {
    var a, b, c, d int
    fmt.Scanf("%d", &a)
    fmt.Scanf("%d", &b)
    fmt.Scanf("%d", &c)
    fmt.Scanf("%d", &d)
    fmt.Println(min4(a, b, c, d))
}
