package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now().Format("1999-01-02-15-04-05")
	p(now)
}
