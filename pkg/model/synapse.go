package model

// 突触结构体
type Synapse struct {
	from   *Neuron
	to     *Neuron
	weight float64
}

// 创建突出链接
func NewSynapse(from *Neuron, to *Neuron, weight float64) *Synapse {
	return &Synapse{
		from:   from,
		to:     to,
		weight: weight,
	}
}
