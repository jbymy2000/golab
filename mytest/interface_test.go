package mytest

import (
	"mygolang/mytest/m1"
	"testing"
)

type a interface {
	m1()
	m2()
}

func Test_interface(t *testing.T) {
	var m m1.M1
	_, ok := interface{}(m).(a)
	println(ok)
	aFunction(m)
}

func aFunction(aa a){
	return
}