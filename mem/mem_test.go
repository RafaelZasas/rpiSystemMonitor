package mem

import "testing"

func TestFreeMemRange(t *testing.T) {
	expectedMin, expectedMax := 0.0, 100.0
	actual, err := GetFree()
	if err != nil {
		t.Fatalf("Error getting free memory\nError: %v", err)
	}

	if actual < expectedMin || actual > expectedMax {
		t.Fatalf("Recieved free mem outside expected range\nExpected: [%v, %v]\nactual: %v", expectedMin, expectedMax, actual)
	}
}


func TestUsedMemRange(t *testing.T) {
	expectedMin, expectedMax := 0.0, 100.0
	actual, err := GetUsed()
	if err != nil {
		t.Fatalf("Error getting used memory\nError: %v", err)
	}

	if actual < expectedMin || actual > expectedMax {
		t.Fatalf("Recieved used mem outside expected range\nExpected: [%v, %v]\nactual: %v", expectedMin, expectedMax, actual)
	}
}
