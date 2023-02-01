package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "regexp"
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
        if len(values) == 4 {
            money, _ := strconv.Atoi(values[3])
            if _, ok := mapping[values[1]]; ok {
                mapping[values[1]] -= money
            } else {
                mapping[values[1]] = -money
            }
            if _, ok := mapping[values[2]]; ok {
                mapping[values[2]] += money
            } else {
                mapping[values[2]] = money
            }
        } else if len(values) == 2 {
            if values[0] == "INCOME" {
                percent, _ := strconv.Atoi(values[1])
                for person := range mapping {
                    if mapping[person] > 0 {
                        mapping[person] += int(float64(percent) / 100 * float64(mapping[person]))
                    }
                }
            } else {
                if _, ok := mapping[values[1]]; ok {
                    fmt.Println(mapping[values[1]]) 
                } else {
                    fmt.Println("ERROR")
                } 
            }
        } else {
            if values[0] == "DEPOSIT" {
                money, _ := strconv.Atoi(values[2])
                if _, ok := mapping[values[1]]; ok {
                    mapping[values[1]] += money
                } else {
                    mapping[values[1]] = money
                } 
            } else {
                money, _ := strconv.Atoi(values[2])
                if _, ok := mapping[values[1]]; ok {
                    mapping[values[1]] -= money
                } else {
                    mapping[values[1]] = -money
                } 
            }
        }
    }
}
