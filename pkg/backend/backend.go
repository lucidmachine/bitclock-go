package backend

import (
	"fmt"
	"time"
)

type BitDigit [4]int
type BitTime [4]BitDigit

func NewBitDigit(i int) BitDigit {
	eightsBit := (8 & i) >> 3
	foursBit := (4 & i) >> 2
	twosBit := (2 & i) >> 1
	onesBit := (1 & i)

	return [4]int{eightsBit, foursBit, twosBit, onesBit}
}

func NewBitTime(t time.Time) BitTime {
	hour := t.Local().Hour()
	minute := t.Local().Minute()

	hourTens := NewBitDigit(hour / 10)
	hourOnes := NewBitDigit(hour % 10)
	minuteTens := NewBitDigit(minute / 10)
	minuteOnes := NewBitDigit(minute % 10)

	return [4]BitDigit{hourTens, hourOnes, minuteTens, minuteOnes}
}

func (bt BitTime) String() string {
	return fmt.Sprintf(
		"%d%d%d%d\n%d%d%d%d\n%d%d%d%d\n%d%d%d%d",
		bt[0][0],
		bt[1][0],
		bt[2][0],
		bt[3][0],
		bt[0][1],
		bt[1][1],
		bt[2][1],
		bt[3][1],
		bt[0][2],
		bt[1][2],
		bt[2][2],
		bt[3][2],
		bt[0][3],
		bt[1][3],
		bt[2][3],
		bt[3][3],
	)
}
