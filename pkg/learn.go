package pkg

func Learning(network *NeuralNetwork, data [][]float64, learningRate float64, output bool) []float64 {
	// 处理输入数据并更新全局记忆
	processInput(network, data, learningRate)

	// 模块之间的协作
	collaborate(network)

	if output {
		return getOutput(network)
	}

	return []float64{}
}

func Output(network *NeuralNetwork) []float64 {
	return getOutput(network)
}
