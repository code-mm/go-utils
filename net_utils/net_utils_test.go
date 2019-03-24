package net_utils

import (
	"fmt"
	"testing"
)

func TestNetUtils(t *testing.T) {
	str := Get("https://www.baidu.com/")
	fmt.Println(str)
}
