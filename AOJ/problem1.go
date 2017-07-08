package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
  sc.Split(bufio.ScanWords)
  m, n := nextInt(), nextInt()
  amari := make([]int, m)
  m_i := make([]int, m)
  for i := 0; i < m; i++ {
    m_i[i] = nextLine()
  }
  fmt.Println(amari)
  fmt.Println(m)
  fmt.Println(n)
  fmt.Println(m_i)
}

func nextInt() int {
  sc.Scan()
  i, e := strconv.Atoi(sc.Text())
  if e != nil {
    panic(e)
  }
  return i
}

func nextLine() int {
  sc.Scan()
  i, _ := strconv.Atoi(sc.Text())
  return i
}
