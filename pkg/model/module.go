package model

import (
	"github.com/yubing744/hebbian-net/pkg/utils"
)

// 模块结构体
type Module struct {
	neurons []*Neuron
}

// 创建模块
func NewModule(numNeurons, inputSize int) *Module {
	module := &Module{}

	for i := 0; i < numNeurons; i++ {
		module.neurons = append(module.neurons, NewNeuron(inputSize))
	}

	return module
}

// 前向传播
func (module *Module) forwardPass(input []float64) {
	for _, neuron := range module.neurons {
		neuron.value = 1.0 / (1.0 + utils.EuclideanDistance(neuron.weights, input))
	}
}

// 竞争学习
func (module *Module) competitiveLearning(input []float64, learningRate float64) {
	// 找到最接近输入的神经元（胜者）
	var winner *Neuron
	minDist := utils.EuclideanDistance(module.neurons[0].weights, input)
	winner = module.neurons[0]

	for _, neuron := range module.neurons[1:] {
		dist := utils.EuclideanDistance(neuron.weights, input)
		if dist < minDist {
			minDist = dist
			winner = neuron
		}
	}

	// 根据Hebbian学习规则更新胜者的权重
	if winner != nil {
		winner.HebbianLearning(input, learningRate)
	}
}
