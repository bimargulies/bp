package omroncsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type BP struct {
	Stamp     time.Time
	Systolic  int
	Diastolic int
	Pulse     int
}

func ReadFile(path string, start time.Time, end time.Time) ([]BP, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s; %v", path, err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	r := []BP{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading %s; %v", path, err)
		}
		dateString := record[0]
		timeString := record[1]
		systolic, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, fmt.Errorf("bad systolic %s for %s in %s", record[2], dateString, path)
		}
		diastolic, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, fmt.Errorf("bad diastolic %s for %s in %s", record[3], dateString, path)
		}
		pulse, err := strconv.Atoi(record[4])
		if err != nil {
			return nil, fmt.Errorf("bad pulse %s for %s in %s", record[4], dateString, path)
		}

		stamp, err := toTime(dateString, timeString)
		if err != nil {
			return nil, err
		}

		r = append(r, BP{Stamp: stamp, Systolic: systolic, Diastolic: diastolic, Pulse: pulse})
	}

	return r, nil
}

func toTime(dateString string, timeString string) (time.Time, error) {
	dateTime := fmt.Sprintf("%s %s", dateString, timeString)
	stamp, err := time.Parse("Jan 2 2026 03:04 PM", dateTime)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse %s as timestamp", dateTime)
	}
	return stamp, nil
}
