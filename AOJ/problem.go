package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
  sc.Split(bufio.ScanWords)
  a, op, b, _, c := nextString(), nextString(), nextString(), nextString(), nextString()
  var result int
  if a == "x" {
    bi, _ := strconv.Atoi(b)
    ci, _ := strconv.Atoi(c)
    if op == "+" {
      result = ci - bi
    } else if (op == "-") {
      result = ci + bi
    }
  } else if b == "x" {
    ai, _ := strconv.Atoi(a)
    ci, _ := strconv.Atoi(c)
    if op == "+" {
      result = ci - ai
    } else if op == "-" {
      result = ai - ci
    }
  } else if c == "x" {
    ai, _ := strconv.Atoi(a)
    bi, _ := strconv.Atoi(b)
    switch op {
    case "+":
      result = ai + bi
    case "-":
      result = ai - bi
    }
  } else {
    result = 0
  }
  fmt.Println(result)
}

func nextString() string {
  sc.Scan()
  return sc.Text()
}
