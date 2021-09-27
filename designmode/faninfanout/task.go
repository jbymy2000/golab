package faninfanout

type  Runable interface{
	Run(input string)float64
}

type Task struct {
	Data map[string]float64
	Id string
	Out float64
}

func (t *Task)Run(){
	t.Out=t.Data[t.Id]
}

func ProcessTaskWithFaninFanout(taskQueue []*Task,parallel int16)(res map[string]float64){
	in:=FanIn(taskQueue)
	out:=FanOut(in,parallel)
	res=make(map[string]float64, len(taskQueue))
	for t:=range out{
		t.Run()
	}
	return  res
}

func FanIn(taskQueue []*Task)(in chan *Task){
	in=make(chan *Task,len(taskQueue))
	go func() {
		defer close(in)
		for _,task:= range taskQueue{
			in<-task
		}
	}()
	return in
}

func FanOut(in chan *Task,parallel int16)(out chan *Task){
	out=make(chan *Task,10)
	var i int16
	for i=0;i < parallel;i++{
		go func() {
			t,_:=<-in
			t.Run()
			out<-t
		}()
	}
	return out
}



