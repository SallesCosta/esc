package js

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

	emptySlice := []float64{}

	if Includes(emptySlice, 1.0) {
		t.Errorf("Includes failed: expected false, got true")
	}
}

func TestFilter(t *testing.T) {
	words := []string{"spray", "elite", "exuberant", "destruction", "present"}

	result := Filter(words, func(word string) bool {
		return len(word) > 6
	})

	expected := []string{"exuberant", "destruction", "present"}

	if !areEqual(result, expected) {
		t.Errorf("Filter failed: expected %v, got %v", expected, result)
	}

	result = Filter(words, func(word string) bool {
		return len(word) >= 10
	})

	expected = []string{"destruction"}

	if !areEqual(result, expected) {
		t.Errorf("Filter failed: expected %v, got %v", expected, result)
	}
}

func areEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
