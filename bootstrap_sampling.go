package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Function to perform bootstrap sampling
func bootstrapSampling(data []float64, nBootstrap int) []float64 {
	sampleMeans := make([]float64, nBootstrap)
	var wg sync.WaitGroup
	for i := 0; i < nBootstrap; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sample := make([]float64, len(data))
			for j := range sample {
				sample[j] = data[rand.Intn(len(data))]
			}
			sampleMeans[i] = median(sample)
		}(i)
	}
	wg.Wait()
	return sampleMeans
}

// Function to calculate median
func median(data []float64) float64 {
	middle := len(data) / 2
	if len(data)%2 == 0 {
		return (data[middle-1] + data[middle]) / 2
	}
	return data[middle]
}

// Function to calculate standard error
func standardError(sampleMeans []float64) float64 {
	var sum, mean float64
	for _, value := range sampleMeans {
		sum += value
	}
	mean = sum / float64(len(sampleMeans))
	var varianceSum float64
	for _, value := range sampleMeans {
		varianceSum += (value - mean) * (value - mean)
	}
	variance := varianceSum / float64(len(sampleMeans)-1)
	return sqrt(variance / float64(len(sampleMeans)))
}

// Function to calculate square root
func sqrt(value float64) float64 {
	return math.Sqrt(value)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	data := generateDistributions()
	nBootstrap := 10000

	for _, dist := range data {
		bootstrapResults := bootstrapSampling(dist.values, nBootstrap)
		se := standardError(bootstrapResults)
		fmt.Printf("Distribution: %s, Standard Error of the Median: %f\n", dist.name, se)
	}
}

// Distribution struct
type Distribution struct {
	name   string
	values []float64
}

// Function to generate distributions
func generateDistributions() []Distribution {
	return []Distribution{
		{"Positively Skewed", generatePositivelySkewed(1000)},
		{"Symmetric", generateSymmetric(1000)},
		{"Negatively Skewed", generateNegativelySkewed(1000)},
	}
}

// Function to generate positively skewed distribution
func generatePositivelySkewed(size int) []float64 {
	dist := make([]float64, size)
	for i := range dist {
		dist[i] = rand.ExpFloat64()
	}
	return dist
}

// Function to generate symmetric distribution
func generateSymmetric(size int) []float64 {
	dist := make([]float64, size)
	for i := range dist {
		dist[i] = rand.NormFloat64() * 10 + 50
	}
	return dist
}

// Function to generate negatively skewed distribution
func generateNegativelySkewed(size int) []float64 {
	dist := generatePositivelySkewed(size)
	sort.Sort(sort.Reverse(sort.Float64Slice(dist)))
	return dist
}
