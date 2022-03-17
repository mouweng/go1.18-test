package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	s := "abcdefghijklmn"
	s1 := s[:4]
	s2 := strings.Clone(s[:4])

	sHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("sHeader Len : %d, sHeader Data : %d\n", sHeader.Len, sHeader.Data)
	s1Header := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	fmt.Printf("s1Header Len : %d, s1Header Data : %d\n", s1Header.Len, s1Header.Data)
	s2Header := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Printf("s2Header Len : %d, s2Header Data : %d\n", s2Header.Len, s2Header.Data)
}

/*
[Running] go run "./main.go"
sHeader Len : 14, sHeader Data : 17448912
s1Header Len : 4, s1Header Data : 17448912
s2Header Len : 4, s2Header Data : 824633811092
*/
