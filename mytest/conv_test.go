package mytest

import "testing"

func Test_abc(t *testing.T){
	var response interface{}
	response="[]"
	stringResp := string(response.([]uint8))
	println(stringResp)
	println(111)
}

func Test_nilSlice(t *testing.T){
	var a []int64
	a=nil
	println(len(a))
}