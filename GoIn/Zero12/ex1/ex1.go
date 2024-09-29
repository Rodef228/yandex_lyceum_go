package ex1

import (
	"fmt"
)

func ConcatStringsAndInt(str1, str2 string, num int) string {
	return str1 + str2 + fmt.Sprintf("%d", num)
}
