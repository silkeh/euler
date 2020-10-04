package main

import (
	"github.com/silkeh/euler/consumer"
	"github.com/silkeh/euler/filter"
	"github.com/silkeh/euler/generator"
	"sync"
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
	var wg sync.WaitGroup
	wg.Add(len(p.Filters) + 2)

	// Create the 'done' channel that indicates success to the generator.
	done := make(chan bool, 1)
	defer close(done)

	// Create a list for all channels
	channels := make([]chan int, len(p.Filters)+1)

	// Create the first output channel
	channels[0] = make(chan int)

	// Start generator
	go func() {
		p.Generator.Run(channels[0], done)
		wg.Done()
	}()

	// Create all filters
	for i := range p.Filters {
		// Create the output channel
		channels[i+1] = make(chan int)

		// Copy i to local scope to avoid it changing before the filter starts
		j := i

		// Run the filter in a goroutine
		go func() {
			p.Filters[j].Run(channels[j], channels[j+1])
			wg.Done()
		}()
	}

	// Start the consumer
	go func() {
		p.Consumer.Run(channels[len(channels)-1], done)
		wg.Done()
	}()

	// Wait for everything to exit
	wg.Wait()

	// Return the result
	return p.Consumer.Result()
}
