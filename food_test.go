package main

import "testing"

func TestRandomFood(t *testing.T) {
	food := RandomFood()
	if len(food) <= 0 {
		t.Errorf("RandomFood returned an empty string")
	}
}
