package main

import (
	"github.com/silkeh/euler/consumer"
	"github.com/silkeh/euler/filter"
	"github.com/silkeh/euler/generator"
)

// Pipeline represents a pipeline consisting of a Generator, some Filters and a Consumer.
type Pipeline struct {
	Generator generator.Generator
	Filters   []filter.Filter
	Consumer  consumer.Consumer
}

// RunPipeline executes a pipeline with the given components and returns the result.
func RunPipeline(g generator.Generator, c consumer.Consumer, fs ...filter.Filter) int {
	return NewPipeline(g, c, fs...).Run()
}

// NewPipeline creates a pipeline with the given components.
func NewPipeline(g generator.Generator, c consumer.Consumer, fs ...filter.Filter) *Pipeline {
	return &Pipeline{
		Generator: g,
		Filters:   fs,
		Consumer:  c,
	}
}

// Run the pipeline.
func (p *Pipeline) Run() int {
	// Create the 'done' channel that indicates success to the generator.
	done := make(chan bool, 1)

	// Create the first output channel
	channel0 := make(chan int)
	channel1 := make(chan int)

	// Start generator
	go p.Generator.Run(channel0, done)
	// Start Filters
	go filter.NewSet(p.Filters...).Run(channel0, channel1)
	// Start the consumer
	p.Consumer.Run(channel1, done)

	// Return the result
	return p.Consumer.Result()
}
