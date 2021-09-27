package decorator

import "testing"

func Test_decorator(t *testing.T){
	var hh = human{}
	var h= NewClosedHuman(&hh)
	var h1= NewClosedHuman(h)
	h1.puton()
}