package pkg

// Hebbian学习规则
func hebbianLearning(neuron *Neuron, input []float64, learningRate float64) {
	for i := range neuron.weights {
		neuron.weights[i] += learningRate * input[i] * neuron.value
	}
}

// 前向传播
func forwardPass(module *Module, input []float64) {
	for _, neuron := range module.neurons {
		neuron.value = 1.0 / (1.0 + euclideanDistance(neuron.weights, input))
	}
}

// 竞争学习
func competitiveLearning(module *Module, input []float64, learningRate float64) {
	// 找到最接近输入的神经元（胜者）
	var winner *Neuron
	minDist := euclideanDistance(module.neurons[0].weights, input)
	winner = module.neurons[0]
	for _, neuron := range module.neurons[1:] {
		dist := euclideanDistance(neuron.weights, input)
		if dist < minDist {
			minDist = dist
			winner = neuron
		}
	}
	// 根据Hebbian学习规则更新胜者的权重
	if winner != nil {
		hebbianLearning(winner, input, learningRate)
	}
}

// 处理输入数据并更新全局记忆
func processInput(network *NeuralNetwork, inputs [][]float64, learningRate float64) {
	for _, input := range inputs {
		for _, module := range network.modules {
			forwardPass(module, input)
			competitiveLearning(module, input, learningRate)
			updateMemory(network, module)
		}
	}
}

// 更新全局记忆
func updateMemory(network *NeuralNetwork, module *Module) {
	for i, neuron := range module.neurons {
		network.memory[i] += neuron.value
	}
}

// 信息传递与协作
func collaborate(network *NeuralNetwork) {
	for _, module := range network.modules {
		forwardPass(module, network.memory)
	}
}

// 获取输出
func getOutput(network *NeuralNetwork) []float64 {
	outputValues := make([]float64, len(network.outputs.neurons))
	for i, neuron := range network.outputs.neurons {
		sum := 0.0
		for _, prevNeuron := range network.modules[len(network.modules)-1].neurons {
			sum += prevNeuron.value * neuron.weights[i]
		}
		outputValues[i] = sigmoid(sum)
	}
	return outputValues
}
