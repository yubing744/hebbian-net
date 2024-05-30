package types

type Tensor struct {
	data             [][][]float64
	dim1, dim2, dim3 int
}

func NewTensor(data [][][]float64) *Tensor {
	return &Tensor{
		data: data,
		dim1: len(data),
		dim2: len(data[0]),
		dim3: len(data[0][0]),
	}
}
