package main

import (
  "fmt"
  "os"
  "strings"
  "regexp"
  "sort"
)

func find_depth(child string, relatives map[string]string, depth int) int {
    if _, ok := relatives[child]; ok {
        return find_depth(relatives[child], relatives, depth + 1)
    }
    return depth
}

func main() {
    in, _ := os.ReadFile("input.txt")
    strs := string(in)
    re := regexp.MustCompile(`\r?\n`)
    strs = re.ReplaceAllString(strs, "\n")
    s := strings.Split(strs, "\n")
    counts := make(map[string]int)
    relatives := make(map[string]string)
    for i := 1; i < len(s) - 1; i++ {
        values := strings.Split(s[i], " ")
        relatives[values[0]] = values[1]
        counts[values[0]] = 0
        counts[values[1]] = 0
    }

    for relative := range relatives {
        counts[relative] = find_depth(relative, relatives, 0)
    }
  
    keys := make([]string, 0, len(counts))
    for k := range counts{
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, k := range keys {
        fmt.Println(k, counts[k])
    }
}

