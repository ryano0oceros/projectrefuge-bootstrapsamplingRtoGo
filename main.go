package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// Function to perform bootstrap sampling
func bootstrapSampling(data []float64, nBootstrap int) ([]float64, time.Duration) {
	start := time.Now()
	sampleMeans := make([]float64, nBootstrap)
	var wg sync.WaitGroup
	wg.Add(nBootstrap)
	for i := 0; i < nBootstrap; i++ {
		go func(i int) {
			defer wg.Done()
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			sample := make([]float64, len(data))
			for j := range sample {
				sample[j] = data[r.Intn(len(data))]
			}
			sampleMeans[i] = median(sample)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	return sampleMeans, elapsed
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
	data := generateDistributions()
	nBootstrap := 10000
	fmt.Println("Standard Error of the Median for n_bootstrap= ", nBootstrap, ":")

	for _, dist := range data {
		bootstrapResults, elapsedTime := bootstrapSampling(dist.values, nBootstrap)
		se := standardError(bootstrapResults)
		fmt.Printf("Distribution: %s, Standard Error of the Median: %f\n", dist.name, se)
		fmt.Printf("Time taken for bootstrap sampling: %v\n", elapsedTime)
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
		dist[i] = rand.NormFloat64()*10 + 50
	}
	return dist
}

// Function to generate negatively skewed distribution
func generateNegativelySkewed(size int) []float64 {
	dist := generatePositivelySkewed(size)
	sort.Sort(sort.Reverse(sort.Float64Slice(dist)))
	return dist
}
