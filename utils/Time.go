package utils

import (
	"fmt"
	"time"
)

type Epoch interface {
	int | int64 | string
}

func EpochToTime[check Epoch](check) time.Time {
	var b check
	switch c := b.(type) {
	case string:
		fmt.Printf("string %v", c)
	case int:
		fmt.Printf("string")
	case int64:
		fmt.Printf("string")

	}

	return time.Now()

}

func init() {

	v := EpochToTime("Hello")
	fmt.Println(v)

}
