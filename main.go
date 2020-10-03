package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var a, b int

	fmt.Fscan(r, &a, &b)

	switch (a * b) % 2 {
	case 0:
		fmt.Println("Even")
	case 1:
		fmt.Println("Odd")
	}
}