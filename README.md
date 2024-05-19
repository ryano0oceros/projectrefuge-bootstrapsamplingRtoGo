# ProjectRefuge Bootstrap Sampling: R to Go

## Introduction
This repository showcases implementations of bootstrap sampling in both R and Go, comparing their performance and efficiency. The main objective is to determine under which circumstances Go might be more advantageous over R for statistical computations, especially in cloud-based environments.

## Package Selection and Implementation

### Finding R and Go Packages
The search for appropriate R packages led me to an implementation on [stackoverflow](https://stackoverflow.com/questions/68839044/sampling-with-replacement-for-bootstrapping-r) that I leveraged. (It's my first time programming in R and I'm less confident in this implementation.) For Go, the functionality was easier for me to build natively, and I was able to get it to work leveraging Go's native concurrency primitives (`goroutines` and `sync` package) for efficient parallel processing.

### Building the Go Implementation
The Go implementation involves concurrent bootstrap sampling where each bootstrap iteration runs in a separate goroutine, ensuring that sampling tasks are executed in parallel. This setup significantly enhances performance by reducing overall execution time, demonstrated in our benchmarks. Error handling is done on the `bootstrapSampling` function and data manipulation use slices and efficient sorting algorithms from the `sort` package.

### R Implementation Details
In R, the implementation uses vectorized operations and the built-in `sample` function to perform bootstrap sampling. This approach leverages R's strengths in handling statistical computations but lacks the inherent parallel processing capabilities found in Go. The use of `Sys.time()` for logging execution times helps in profiling and benchmarking the script.

## Testing and Benchmarking

### Execution Times
Execution benchmarks are provided in detail within this repository for both R and Go implementations, showing Go's superior performance in terms of execution time, especially when dealing with large datasets and high degrees of parallelism.

## Recommendations to the Research Consultancy

### Use of Go Over R
My analysis suggests that Go is more suitable than R in scenarios requiring:
- High-performance computation with extensive parallel processing capabilities.
- Large-scale data handling with reduced execution times.
- Cost-effective cloud deployment, given Go's efficiency in resource utilization.








# Performance Benchmarking: R vs. Go Bootstrap Sampling

## Setup Details

- **R Version**: 4.4.0 (2024-04-24) -- "Puppy Cup"
- **Go Version**: go1.22.1 darwin/arm64
- **Machine Specifications**: Apple M1 Max, 64GB RAM, SSD
- **Dataset Size**: 1000 observations per distribution
- **Number of Bootstrap Samples**: 10,000 and 100,000


## Summary
When the number of samples in parallel remained fixed at 1,000 for both R and Go, the Go implementation outperformed R by a factor of ~6x (5.57x and 6.17x for `n_bootstrap` 10,000 and 100,000 respectively - Average ratios: 4.92, 5.81, 5.97 and 6.02, 6.49, 6.00).

Interestingly, when the number of samples in parallel increased to 10,000 the performance differences became more stark. The Go implementation was able to handle the increased parallelism more efficiently, resulting in a ~27x improvement in execution time compared to R. This demonstrates the scalability and efficiency of Go's concurrency model for computational tasks like bootstrap sampling. (Average ratios: 25.89, 28.81, 27.55)

Therefore, for large-scale applications or when processing larger datasets, Go is a more efficient choice for bootstrap sampling due to its superior performance and scalability. It's difficult to pin a cost savings as it depends largely on the workload and the cloud provider, but the reduced execution time can lead to substantial compute savings in cloud-based environments.

## Detailed Execution Times

### R

Sampling n_bootstrap = 10000 - with 1000 samples in parallel
```bash
projectrefuge-bootstrapsamplingRtoGo % Rscript bootstrap_sampling.R
Standard Error of the Median for n_bootstrap = 10000 :
Positively Skewed:  0.003944053 
Symmetric:  0.004468373 
Negatively Skewed:  0.003742951 
Runtimes:
Generation time:  0.0004899502  seconds
Positively Skewed Bootstrap sampling time:  0.4178081  seconds
Symmetric Bootstrap sampling time:  0.4278221  seconds
Negatively Skewed Bootstrap sampling time:  0.401186  seconds
Calculation time:  0.001094818  seconds
```

Sampling n_bootstrap = 100000 - with 1000 samples in parallel
```bash
projectrefuge-bootstrapsamplingRtoGo % Rscript bootstrap_sampling.R
Standard Error of the Median for n_bootstrap = 1e+05 :
Positively Skewed:  0.001252871 
Symmetric:  0.001414359 
Negatively Skewed:  0.00117999 
Runtimes:
Generation time:  0.0004599094  seconds
Positively Skewed Bootstrap sampling time:  4.059178  seconds
Symmetric Bootstrap sampling time:  4.03963  seconds
Negatively Skewed Bootstrap sampling time:  4.042154  seconds
Calculation time:  0.003170013  seconds
```

Sampling n_bootstrap = 100000 - with 10000 samples in parallel
```bash
projectrefuge-bootstrapsamplingRtoGo % Rscript bootstrap_sampling.R
Standard Error of the Median for n_bootstrap = 1e+05 :
Positively Skewed:  0.0003368863 
Symmetric:  0.0003873318 
Negatively Skewed:  0.0003041106 
Runtimes:
Generation time:  0.003623009  seconds
Positively Skewed Bootstrap sampling time:  51.66299  seconds
Symmetric Bootstrap sampling time:  52.0984  seconds
Negatively Skewed Bootstrap sampling time:  51.41881  seconds
Calculation time:  0.001960993  seconds
```

### Go

Sampling n_bootstrap = 10000
```bash
Standard Error of the Median for n_bootstrap=  10000 :
Distribution: Positively Skewed, Standard Error of the Median: 0.006815
Time taken for bootstrap sampling: 84.886875ms
Distribution: Symmetric, Standard Error of the Median: 0.073288
Time taken for bootstrap sampling: 73.539583ms
Distribution: Negatively Skewed, Standard Error of the Median: 0.007455
Time taken for bootstrap sampling: 67.208417ms
Total execution time: 226.366209ms
```

Sampling n_bootstrap = 100000
```bash
Standard Error of the Median for n_bootstrap=  100000 :
Distribution: Positively Skewed, Standard Error of the Median: 0.002103
Time taken for bootstrap sampling: 674.106917ms
Distribution: Symmetric, Standard Error of the Median: 0.021844
Time taken for bootstrap sampling: 622.813959ms
Distribution: Negatively Skewed, Standard Error of the Median: 0.002233
Time taken for bootstrap sampling: 673.189334ms
Total execution time: 1.971267542s
```

Sampling n_bootstrap = 100000 - with 10000 samples in parallel
```bash
Standard Error of the Median for n_bootstrap=  100000 :
Distribution: Positively Skewed, Standard Error of the Median: 0.002241
Time taken for bootstrap sampling: 1.995173208s
Distribution: Symmetric, Standard Error of the Median: 0.022355
Time taken for bootstrap sampling: 1.808388459s
Distribution: Negatively Skewed, Standard Error of the Median: 0.002180
Time taken for bootstrap sampling: 1.866448541s
Total execution time: 5.672271167s
