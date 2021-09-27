package printmetrix

import (
	"github.com/pkg/errors"
)

func PrintMetrix(inputMetrix [][]int64) (err error) {
	if inputMetrix == nil {
		return errors.New("metrix is nil")
	}
	var (
		left  = 0
		right = len(inputMetrix) - 1
	)
	for {
		if left > right {
			break
		}
		for i := left; i <= right; i++ {
			println(inputMetrix[left][i])
		}
		for i := left + 1; i <= right; i++ {
			println(inputMetrix[i][right])
		}
		for i := right - 1; i >= left; i-- {
			println(inputMetrix[right][i])
		}
		for i := right - 1; i > left+1; i-- {
			println(inputMetrix[i][left])
		}
		left++
		right--
	}
	return nil
}
