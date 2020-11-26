package main

import "testing"

func TestInputPayLessThanPriceShouldBeAlert(t *testing.T) {
	v := MoneyChangerLogic(100, 0)
	if len(v) > 0 {
		t.Error("its not for free")
	}
}

func TestInputPriceLessThanPayShouldBeAlert(t *testing.T) {
	v := MoneyChangerLogic(0, 110)
	if len(v) > 0 {
		t.Error("its free")
	}
}
