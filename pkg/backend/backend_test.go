package backend

import (
	"testing"
	"time"
)

func TestNewBitDigit(t *testing.T) {
	bitDigit := NewBitDigit(9)
	if bitDigit != [4]int{1, 0, 0, 1} {
		t.Errorf("Expected: [1, 0, 0, 1]/nActual: %v", bitDigit)
	}

	bitDigit = NewBitDigit(0)
	if bitDigit != [4]int{0, 0, 0, 0} {
		t.Errorf("Expected: [0, 0, 0, 0]/nActual: %v", bitDigit)
	}

	bitDigit = NewBitDigit(15)
	if bitDigit != [4]int{1, 1, 1, 1} {
		t.Errorf("Expected: [1, 1, 1, 1]/nActual: %v", bitDigit)
	}
}

func TestNewBitTime(t *testing.T) {
	bitTime := NewBitTime(time.Date(2022, 5, 21, 12, 34, 56, 00, time.Local))
	if bitTime != [6]BitDigit{
		[4]int{0, 0, 0, 1},
		[4]int{0, 0, 1, 0},
		[4]int{0, 0, 1, 1},
		[4]int{0, 1, 0, 0},
		[4]int{0, 1, 0, 1},
		[4]int{0, 1, 1, 0}} {
		t.Errorf("Expected: [[0, 0, 0, 1], [0, 0, 1, 0], [0, 0, 1, 1], [0, 1, 0, 0], [0, 1, 0, 1], [0, 1, 1, 0]]\nActual: %v", bitTime)
	}
}

func TestBitTimeString(t *testing.T) {
	bitTime := NewBitTime(time.Date(2022, 05, 21, 12, 34, 56, 00, time.Local))
	bitTimeStr := bitTime.String()
	if bitTimeStr != "000000\n000111\n011001\n101010" {
		t.Errorf("Expected:\n000000\n000111\n011001\n101010\nActual:\n%s", bitTimeStr)
	}
}
