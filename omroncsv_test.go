package omroncsv

import (
	"os"
	"testing"
	"time"
)

func TestReadFile(t *testing.T) {
	// Create a temporary CSV file
	file, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	// Write test data to the file
	data := `Dec 20 2025,06:00 am,110,70,80
Jan 2 2026,03:04 pm,120,80,70
Jan 3 2026,04:05 am,130,85,75
Jan 4 2026,00:01 am,140,90,80
`
	if _, err := file.WriteString(data); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	file.Close()

	// Define the time range for the test
	start, _ := time.Parse("Jan 2 2006", "Jan 1 2026")
	end, _ := time.Parse("Jan 2 2006", "Jan 4 2026")

	// Call the function under test
	bps, err := ReadFile(file.Name(), start, end)
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	// Check the results
	if len(bps) != 2 {
		t.Fatalf("expected 2 records, got %d", len(bps))
	}

	expected := []BP{
		{Stamp: time.Date(2026, time.January, 2, 15, 4, 0, 0, time.UTC), Systolic: 120, Diastolic: 80, Pulse: 70},
		{Stamp: time.Date(2026, time.January, 3, 4, 5, 0, 0, time.UTC), Systolic: 130, Diastolic: 85, Pulse: 75},
	}

	for i, bp := range bps {
		if bp != expected[i] {
			t.Errorf("expected %v, got %v", expected[i], bp)
		}
	}
}

func TestToTime(t *testing.T) {
	dateString := "Jan 2 2026"
	timeString := "03:04 pm"
	expected := time.Date(2026, time.January, 2, 15, 4, 0, 0, time.UTC)

	stamp, err := toTime(dateString, timeString)
	if err != nil {
		t.Fatalf("ToTime failed: %v", err)
	}

	if !stamp.Equal(expected) {
		t.Errorf("expected %v, got %v", expected, stamp)
	}
}
