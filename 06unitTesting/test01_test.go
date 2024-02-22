package main

import (
	"testing"
)

func Test_addUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("期望值=%v,实际值=%v\n", 55, res)
	}
	t.Logf("执行正确")
}
