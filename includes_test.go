package esc 

import "testing"

func TestIncludes(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}

	if !Includes(intSlice, 3) {
		t.Errorf("Includes failed: expected true, got false")
	}

	if Includes(intSlice, 6) {
		t.Errorf("Includes failed: expected false, got true")
	}

	stringSlice := []string{"apple", "banana", "orange"}

	if !Includes(stringSlice, "banana") {
		t.Errorf("Includes failed: expected true, got false")
	}

	if Includes(stringSlice, "grape") {
		t.Errorf("Includes failed: expected false, got true")
	}

	// Teste com um slice vazio
	emptySlice := []float64{}

	if Includes(emptySlice, 1.0) {
		t.Errorf("Includes failed: expected false, got true")
	}
}

