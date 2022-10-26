package sd

type Reward struct {
	Type int `json:"类型"`
	Num  int `json:"数量"`
}

func (sd Reward) Clone() (nsd Reward) {
	nsd.Type = sd.Type
	nsd.Num = sd.Num
	return
}

type Item [2]int

func (sd Item) Clone() (nsd Item) {
	nsd[0] = sd[0]
	nsd[1] = sd[1]
	return
}
