package questions

import (
	"strconv"
	"testing"
)

func TestSafetey(T *testing.T) {
	data := &Q2Data{
		reports: [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 9},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9},
		},
	}
	data.checkSafetey()
	if data.safe != 2 {
		T.Errorf("wrong amount of safe reports. expected 2 but got: %q ", strconv.Itoa(data.safe))
	}
}

func TestSafeteyWithRetry(T *testing.T) {
	data := &Q2Data{
		reports: [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 9},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 1, 2, 3, 4},
		},
	}
	data.checkSafeteyWithRetry()
	if data.safe != 4 {
		T.Errorf("wrong amount of safe reports with retries. expected 4 but got: %q ", strconv.Itoa(data.safe))
	}
}
