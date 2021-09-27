package channel

import (
	"testing"
	"time"
	"strconv"
	"encoding/json"
	"fmt"
)

func Test_channel(t *testing.T){
	a:=make(chan int, 1)
	go func(){
		time.Sleep(time.Duration(10)*time.Second)
		a<-1
	}()

	println(<-a)
	//go func(){
	//	println(2)
	//	a<-0
	//}()
}
type pair struct{
	key string
	value int64
}

func Test_write2Channels(t *testing.T){
	map1:=map[string]int64{}
	map2:=map[string]int64{}
	map3:=map[string]int64{}

	channel:=make(chan *pair)
	for i :=0;i<=10;i++{
		key:="a."+strconv.Itoa(i)
		value:=i
		map1[key]=int64(value)
	}

	for i :=0;i<=10;i++{
		key:="b."+strconv.Itoa(i)
		value:=i
		map1[key]=int64(value)
	}

	go func(){
		for k,v:=range map1{
			p:=&pair{k,v}
			channel<-p
		}
	}()

	go func(){
		for k,v:=range map2{
			p:=&pair{k,v}
			channel<-p
		}
	}()

	for i:=range channel{
		map3[i.key]=i.value
	}
	v,_:=json.Marshal(map3)
	fmt.Println(string(v))
}