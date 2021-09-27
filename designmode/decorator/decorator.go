package decorator

type ihuman interface {
	puton()
	putoff()
}

type human struct{
}

func (h *human) puton(){
	println(1)
}

func (h *human) putoff(){
	println(2)
}

type closedhuman struct {
	ihuman
}

func (h *closedhuman) puton(){
	println("a")
	h.puton()
}

func (h *closedhuman) putoff(){
	println("b")
	h.putoff()
}

func NewClosedHuman(h ihuman)ihuman{
	return &closedhuman{ihuman:h}
}




