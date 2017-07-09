package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

var sc = bufio.NewScanner(os.Stdin)
func main () {
  sc.Split(bufio.ScanWords)
  fmt.Println("整列したい配列[5]をスペース区切りで入力してください")
  array := make([]int, 5)
  for i := 0; i < 5; i++ {
    array[i] = nextInt()
  }
  for i := 1; i < 5; i++ {
    insert(array, i, array[i])
  }
  fmt.Println(array)
}

func insert(array []int, i int, value int) {
  for j := i - 1; j >= 0 && array[j] > value; j-- {
    array[j+1] = array[j]
    array[j] = value
  }
  fmt.Println(array)
}

func nextInt() int {
  sc.Scan()
  i, _ := strconv.Atoi(sc.Text())
  return i
}
