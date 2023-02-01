package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "regexp"
  "sort"
)

func main() {
    in, _ := os.ReadFile("input.txt")
    strs := string(in)
    re := regexp.MustCompile(`\r?\n`)
    strs = re.ReplaceAllString(strs, "\n")
    s := strings.Split(strs, "\n")
    mapping := make(map[string]int)
    for i := 0; i < len(s); i++ {
        if s[i] == "" {
            continue
        }
        values := strings.Split(s[i], " ");
        votes, _ := strconv.Atoi(values[1])
        if _, ok := mapping[values[0]]; ok {
            mapping[values[0]] += votes
        } else {
            mapping[values[0]] = votes
        }
    }

    keys := make([]string, 0, len(mapping))
    for k := range mapping{
        keys = append(keys, k)
    }
    sort.Strings(keys)
 
    for _, k := range keys {
        fmt.Println(k, mapping[k])
    }
}
