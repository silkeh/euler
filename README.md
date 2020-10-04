Pipelined Euler
===============

[![godoc](https://godoc.org/github.com/silkeh/euler?status.svg)](https://godoc.org/github.com/silkeh/euler)
[![build status](https://travis-ci.org/silkeh/euler.svg?branch=master)](https://travis-ci.org/silkeh/euler)
[![goreportcard](https://goreportcard.com/badge/github.com/silkeh/euler)](https://goreportcard.com/report/github.com/silkeh/euler)

The goal of this project is to solve the problems in [Project Euler] using pipelines.
A pipeline is structured as follows:

```
Generator -- int --> Filter(s) -- int --> Consumer
```

The components are:

- [`Generator`](https://godoc.org/github.com/silkeh/euler/generator):
  generates a sequence of numbers. This may be stopped by an internal condition,
  or by a signal from the Consumer.

- [`Filter`](https://godoc.org/github.com/silkeh/euler/filter):
  produces numbers based on the values of numbers on the input.

- [`Consumer`](https://godoc.org/github.com/silkeh/euler/consumer):
  consumes numbers. This may be done until the incoming channel is closed,
  or until consumption is stopped by an internal condition.
  In the latter case a signal needs to be sent to the Generator.

The goal is to use the composition of these components to solve the problems.

[Project Euler]: https://projecteuler.net/
