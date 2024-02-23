package main

import (
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("期望值=%v,实际值=%v\n", 55, res)
	}
	t.Logf("执行正确")
}

func TestSub(t *testing.T) {
	res := sub(10, 20)
	if res != 30 {
		t.Fatalf("期望值=%v,实际值=%v\n", 30, res)
	}
	t.Logf("执行正确")
}
