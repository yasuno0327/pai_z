package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
  // fmt.Scanを使う方法
/*var a int
  fmt.Println("数値を入力してください")
  fmt.Scan(&a)
  fmt.Printf("入力された数値は%dです\n", a)
*/
  // bufioのScannerを使う
/*s, t := nextLine(), nextLine()
  fmt.Println(s)
  fmt.Println(t)
*/

  //スペース区切りの時
  //Splitを使う
  sc.Split(bufio.ScanWords)
  n := nextInt()
  fmt.Println(n)
  fmt.Println(nextInt())
  fmt.Println(nextInt())
}

//Splitを使ったスペース区切りでの読み込み

func nextInt() int {
  sc.Scan()
  i, e := strconv.Atoi(sc.Text())
  if e != nil {
    panic(e)
  }
  return i
}

// buffioを使った標準入出力関数
// 改行区切りの時
// os.Stdinを使う

func nextLine() string {
  fmt.Println("文字を入力してください")
  sc.Scan()
  return sc.Text()
}
