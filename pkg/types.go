package pkg

import "math/rand"

// 神经元结构体
type Neuron struct {
	weights []float64
	value   float64
}

// 突触结构体
type Synapse struct {
	from   *Neuron
	to     *Neuron
	weight float64
}

// 模块结构体
type Module struct {
	neurons []*Neuron
}

// 神经网络结构体
type NeuralNetwork struct {
	modules []*Module
	outputs *Module
	memory  []float64
}

// 创建神经元
func NewNeuron(inputSize int) *Neuron {
	weights := make([]float64, inputSize)
	for i := range weights {
		weights[i] = rand.Float64()
	}
	return &Neuron{weights: weights}
}

// 创建模块
func NewModule(numNeurons, inputSize int) *Module {
	module := &Module{}
	for i := 0; i < numNeurons; i++ {
		module.neurons = append(module.neurons, NewNeuron(inputSize))
	}
	return module
}

// 初始化神经网络
func NewNetwork(numModules, numNeurons, inputSize int) *NeuralNetwork {
	network := &NeuralNetwork{}
	for i := 0; i < numModules; i++ {
		network.modules = append(network.modules, NewModule(numNeurons, inputSize))
	}
	network.memory = make([]float64, inputSize)
	return network
}
