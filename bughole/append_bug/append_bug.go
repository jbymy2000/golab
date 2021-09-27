package append_bug  
import "encoding/json"

type m  struct {
	L []int64
}

func AppendBug(){
	 ms:= make( []*m,0)
	for i:=0;i<10;i++{
		ms= append(ms,&m{L:[]int64{1,2,3,4}})
	}
	r:=[]int64{5,4,3,2}
	for i:=range  ms {
		ms[i].L = append(r,ms[i].L...)
		v,_:=json.Marshal(ms)
		println(string(v))
	}
}