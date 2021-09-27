package printmetrix

import (
	"github.com/pkg/errors"
)

func PrintMetrixV1(inputMetrix [][]int64) (err error) {
	length,err:= MetrixLen(inputMetrix)
	if err!=nil{
		return errors.Wrap(err,"error")
	}
	var left,right int
	left = 0
	right = length-1
	for (left <= right){
		for i:=left;i<=right;i++{
			println(inputMetrix[left][i])
		}
		for j:=left + 1;j <= right;j++{
			println(inputMetrix[j][right])
		}
		for k:=right - 1;k >= left;k--{
			println(inputMetrix[right][k])
		}
		for m:=right -1 ;m >= left + 1;m--{
			println(inputMetrix[m][right])
		}
		left++
		right--
	}
	return errors.New("error")
}

func MetrixLen(inputMetrix [][]int64)(int,error){
	if len(inputMetrix)==0{
		return 0,errors.New("metrix len is 0")
	}
	l:=len(inputMetrix)
	for _,raw:=range inputMetrix{
		if len(raw)!=l{
			return 0,errors.New("metrix len not metch")
		}
	}
	return l,nil
}
