package main

import (
	"fmt"
	"time"

	backend "github.com/lucidmachine/bitclock-go/pkg/backend"
)

func main() {
	now := time.Now()
	bc := backend.NewBitTime(now)
	fmt.Println(bc)
}
