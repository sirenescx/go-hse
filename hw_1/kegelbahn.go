package main

import (
	"fmt"
)

func main() {
    var n, k int
    fmt.Scanln(&n, &k)
    values := make([]bool, n)
    var l, r int
    for i := 0; i < k; i++ {
        fmt.Scanln(&l, &r)
        for j := l - 1; j < r; j++ {
            values[j] = true
        }
    }

    for i := 0; i < n; i++ {
        if values[i] {
            fmt.Printf(".")
        } else {
            fmt.Printf("I")
        }
    }
}
