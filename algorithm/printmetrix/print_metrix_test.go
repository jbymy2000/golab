package printmetrix

import (
	"testing"
)

func TestPrintMetrix(t *testing.T) {
	PrintMetrix([][]int64{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	})
}

func TestPrintMetrixV1(t *testing.T) {
	PrintMetrixV1([][]int64{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	})
}

