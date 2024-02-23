package store

import "testing"

func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "皮卡丘",
		Age:   1,
		Skill: "十万伏特",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("结果错误")
	}
	t.Logf("结果正确")
}

func TestReStore(t *testing.T) {
	var monster = &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("结果错误")
	}
	t.Logf("结果正确")
}
