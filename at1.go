package main

import (
  "bufio"
  "os"
  "fmt"
  "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
  m, n := nextLIne(), nextLine()
}

func nextLine() int {
  sc.Scan()
  i, e := strconv.Atoi(sc.Text())
  if e != nil {
    panic(e)
  }
  return i
}
