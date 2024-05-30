package model

import (
	"github.com/yubing744/hebbian-net/pkg/utils"
)

// 神经网络结构体
type NeuralNetwork struct {
	inputs  []*Module
	modules []*Module
	outputs []*Module

	learningRate float64
}

// 初始化神经网络
func NewNetwork(numModules, numNeurons, inputSize int, learningRate float64) *NeuralNetwork {
	network := &NeuralNetwork{
		learningRate: learningRate,
	}

	for i := 0; i < numModules; i++ {
		network.modules = append(network.modules, NewModule(numNeurons, inputSize))
	}

	return network
}

func (network *NeuralNetwork) Learn(input [][]float64, output bool) []float64 {
	// 处理输入数据并更新全局记忆
	network.processInput(input, network.learningRate)

	if output {
		return network.GetOutput(0)
	}

	return []float64{}
}

// 处理输入数据并更新全局记忆
func (network *NeuralNetwork) processInput(inputs [][]float64, learningRate float64) {
	for _, input := range inputs {
		for _, module := range network.modules {
			module.forwardPass(input)
			module.competitiveLearning(input, learningRate)
		}
	}
}

// 获取输出
func (network *NeuralNetwork) GetOutput(index int) []float64 {
	module := network.outputs[index]
	outputValues := make([]float64, len(module.neurons))

	for i, neuron := range module.neurons {
		sum := 0.0
		for _, prevNeuron := range network.modules[len(network.modules)-1].neurons {
			sum += prevNeuron.value * neuron.weights[i]
		}
		outputValues[i] = utils.Sigmoid(sum)
	}

	return outputValues
}
