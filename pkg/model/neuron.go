package model

import "math/rand"

// 神经元结构体
type Neuron struct {
	weights []float64
	value   float64
}

// 创建神经元
func NewNeuron(inputSize int) *Neuron {
	weights := make([]float64, inputSize)
	for i := range weights {
		weights[i] = rand.Float64()
	}
	return &Neuron{weights: weights}
}

// Hebbian学习规则
func (neuron *Neuron) HebbianLearning(input []float64, learningRate float64) {
	for i := range neuron.weights {
		neuron.weights[i] += learningRate * input[i] * neuron.value
	}
}
