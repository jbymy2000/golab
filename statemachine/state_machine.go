package state_machine

const stateName string

const (
	stateName string = "work"
	stateName string = "sleep"
)

type state interface{
	transfers map[string]StateTransfer
	GetCurState()state 
	Execute()state
	GetWord()string
	BeforeEnter()
}

type Work struct{
	stateName
}

func (p *work) GetCurState(){
	return stateName
}

func (p *work) GetWord()string{
	return fmt.Printf("i am %s ing",stateName)
}

func (p *work) BeforeEnter(){
	return
}

type StateTransfer func(currentState state)state {
	s.GetWord()
	s.BeforeEnter()
}

type StateMachine struct {
	state state
	data map[stateName]state
}