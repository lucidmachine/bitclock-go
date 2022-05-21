package frontend

import (
	"fmt"

	backend "github.com/lucidmachine/bitclock-go/pkg/backend"
)

func render(bt backend.BitTime) {
	fmt.Println(bt)
}
