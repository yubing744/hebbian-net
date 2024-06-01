package model

import (
	"github.com/yubing744/hebbian-net/pkg/types"
	"github.com/yubing744/hebbian-net/pkg/utils"
)

// 模块结构体
type Module struct {
	input   []*types.Neuron
	neurons []*types.Neuron
	output  []*types.Neuron
}

// 创建模块
func NewModule(numNeurons, inputSize int) *Module {
	module := &Module{
		input:  make([]*types.Neuron, 1),
		output: make([]*types.Neuron, 1),
	}

	for i := 0; i < numNeurons; i++ {
		module.neurons = append(module.neurons, types.NewNeuron(inputSize))
	}

	return module
}

// 前向传播
func (module *Module) forwardPass(input []float64) {
	for _, neuron := range module.neurons {
		neuron.Value = 1.0 / (1.0 + utils.EuclideanDistance(neuron.Weights, input))
	}
}

// 竞争学习
func (module *Module) competitiveLearning(input []float64, learningRate float64) {
	// 找到最接近输入的神经元（胜者）
	var winner *types.Neuron
	minDist := utils.EuclideanDistance(module.neurons[0].Weights, input)
	winner = module.neurons[0]

	for _, neuron := range module.neurons[1:] {
		dist := utils.EuclideanDistance(neuron.Weights, input)
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
