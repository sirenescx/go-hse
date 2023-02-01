package main

import (
  "fmt"
  "os"
  "strings"
  "regexp"
)

func main() {
    in, _ := os.ReadFile("input.txt")
    strs := string(in)
    re := regexp.MustCompile(`\r?\n`)
    strs = re.ReplaceAllString(strs, " ")
    strs = strings.Trim(strs, " ")
    if len(strs) == 0 {
      return
    }
    s := strings.Split(strs, " ")
    mapping := make(map[string]int)
    for i := 0; i < len(s); i++ {
        if s[i] == "" {
            continue
        }
        if _, ok := mapping[s[i]]; ok {
            mapping[s[i]] += 1
        } else {
            mapping[s[i]] = 1
        }
    }
    // fmt.Println(mapping)

    max := -1
    var maxString string
    for key := range mapping {
        if mapping[key] > max {
            maxString = key
            max = mapping[key]
        } 
        if mapping[key] == max {
            if key < maxString {
                maxString = key
            }
        }
    }

    fmt.Println(maxString)
}
