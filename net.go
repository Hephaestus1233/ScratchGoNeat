package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type NeuralNet struct {
	inputs      int
	connections []*synapse
	neurons     []*neuron
	outputs     []*neuron
}

//NewNet Standard feed-forward network is generated
func NewNet(inputSize int, layerCount int, layerSize int, outputSize int) (NeuralNet, error) {
	rand.Seed(time.Now().UnixNano())
	if inputSize <= 0 || outputSize <= 0 {
		return NeuralNet{}, fmt.Errorf("Invalid input/output size, must be greater than zero")
	}

	var nodes []*neuron
	var connections []*synapse

	var structure [][]*neuron //only used to make life easier when adding all the connections

	//Create required Neurons
	nextId := 0
	for i := 0; i < layerCount; i++ {
		structure = append(structure, []*neuron{})
		for j := 0; j < layerSize; j++ {
			temp := &neuron{
				nextId,
				0,
				i,
				false,
			}
			nodes = append(nodes, temp)
			structure[i] = append(structure[i], temp)
			nextId++
		}
	}

	//Connect each layer
	for i := 0; i < len(nodes); i++ {
		neuron := nodes[i]
		if neuron.layer > 0 {
			for j := 0; j < layerSize; j++ {
				temp := &synapse{
					structure[neuron.layer-1][j],
					neuron,
					rand.Float64() / math.MaxFloat64,
				}
				connections = append(connections, temp)
			}
		}
	}

	//Finish with Outputs
	var out []*neuron
	out = append(out, &neuron{
		-1,
		0,
		-1,
		false,
	})
	out = append(out, &neuron{
		-1,
		0,
		-1,
		false,
	})

	return NeuralNet{
		inputSize,
		connections,
		nodes,
		out,
	}, nil
}
