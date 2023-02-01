package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
  "strconv"
)

type void struct{}

func main() {
    in := bufio.NewReader(os.Stdin)
    nums, _ := in.ReadString('\n')
    set := make(map[float64]void)
    var placeholder void
    nums = strings.Trim(nums, "\n")
    nums = strings.Trim(nums, " ")
    if len(nums) == 0 {
      return
    }
    s := strings.Split(strings.Trim(nums, "\n"), " ")
    for i := 0; i < len(s); i++ {
        value, _ := strconv.ParseFloat(s[i], 64)
        if _, ok := set[value]; ok {
            fmt.Println("YES")
        } else {
            set[value] = placeholder
            fmt.Println("NO")
        }
    }
}
