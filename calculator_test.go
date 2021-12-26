package main

import "testing"

func TestExpCalculator(t *testing.T) {
	truncatedCalculation := int(expCalculator());
	if truncatedCalculation != 249999994505700352 {
		t.Errorf("calculation yielded unexpected results , got: %d, expected:2499999464254802", truncatedCalculation);
	}
}
