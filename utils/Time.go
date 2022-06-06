package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type EpochInt interface {
	uint | uint32 | uint64 | int | int32 | int64
}
type Epoch interface {
	EpochInt | string
}

func EpochToTime[check Epoch](b check) time.Time {
	var fonInt int64
	c := reflect.TypeOf(fonInt)
	fmt.Println(c.Name())

	switch b.(type) {
	case string:
		c, err := strconv.Atoi(b.(string))
		if err != nil {
			fmt.Println("Wrong input string. error:%v", err)

		} else {
			fonInt = int64(c)

		}

	case uint, uint32, uint64, int, int32:
		fonInt = int64(EpochInt(b))
	case int64:
		fonInt = b.(int64)

	}
	fontime := time.Unix(int64(fonInt), 0)
	return fontime

}

func init() {

	v := EpochToTime("Hello")
	fmt.Println(v)

}
