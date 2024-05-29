package main

import (
	"fmt"
	"math/rand"

	"github.com/yubing744/hebbian-net/pkg"
)

func main() {
	numModules := 3
	numNeurons := 10
	inputSize := 5
	learningRate := 0.1
	numIterations := 100

	// 初始化神经网络
	network := pkg.NewNetwork(numModules, numNeurons, inputSize)

	// 无监督学习过程
	for iter := 0; iter < numIterations; iter++ {
		// 随机生成输入数据
		inputs := make([][]float64, numModules)
		for i := range inputs {
			input := make([]float64, inputSize)
			for j := range input {
				input[j] = rand.Float64()
			}
			inputs[i] = input
		}

		pkg.Learning(network, inputs, learningRate, false)
	}

	// 预测
	output := pkg.Predict(network)
	fmt.Printf("%v", output)
}
