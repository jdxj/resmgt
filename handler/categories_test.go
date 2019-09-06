package handler

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStrToInt(t *testing.T) {
	num, err := strconv.Atoi("")
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
}
