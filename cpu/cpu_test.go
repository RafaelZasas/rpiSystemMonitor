package cpu

import "testing"

func TestTempRange(t *testing.T) {
	expectedMin, expectedMax := 0.0, 120.0
	actual, err := GetTemp()
	if err != nil {
		t.Fatalf("Error retrieving cpu temp\nError: %v", err)
	}

	if actual < expectedMin || actual > expectedMax {
		t.Fatalf("return temp outside of acceptable range.\nExpected temp within range [%v, %v]\nRecieved %v", expectedMin, expectedMax, actual)
	}
}

func TestFreqRange(t *testing.T) {
	expectedMin, expectedMax := 0.0, 100000.0
	actual, err := GetFreq()
	actual = actual / 1000
	if err != nil {
		t.Fatalf("Error retrieving cpu freq\nError: %v", err)
	}

	if actual < expectedMin || actual > expectedMax {
		t.Fatalf("Return freq outside acceptable range.\nExpected freq within range [%v, %v]\nRecieved %v", expectedMin, expectedMax, actual)
	}
}
