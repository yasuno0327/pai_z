package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var ar = make([]string, 0, 100)

func main() {
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		ar = append(ar, sc.Text())
	}
	fmt.Println(ar)
}
