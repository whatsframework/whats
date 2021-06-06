package utils

import (
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	for i := 0; i < 3; i++ {
		fmt.Println(RandString(32))
	}
}
