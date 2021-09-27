package faninfanout

import (
	"strconv"
	"testing"
	"time"
)



func TestProcessTaskWithFaninFanout(t *testing.T) {
	tasks:= []*Task{}
	for i:=0;i<=10000;i++{
		key:=strconv.FormatInt(int64(i),10)
		task:=&Task{map[string]float64{key: float64(i * 1.0)},key,0}
		tasks= append(tasks, task)
	}
	println(ProcessTaskWithFaninFanout(tasks,1))
}

func TestChannel(t *testing.T){
	ch1:=make(chan int,10)
	ch2:=make(chan int,10)
	ch2<-0
	ch2<-1
	ch2<-3

	go func() {
		for i:=range ch1{
			println(i)
		}
	}()

	go func() {
		for i:=range ch2{
			ch1<-i
		}
	}()
	time.Sleep(1000)
}


