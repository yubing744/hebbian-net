package types

// 突触结构体
type Synapse struct {
	From   *Neuron
	To     *Neuron
	Weight float64
}

// 创建突出链接
func NewSynapse(from *Neuron, to *Neuron, weight float64) *Synapse {
	return &Synapse{
		From:   from,
		To:     to,
		Weight: weight,
	}
}
