# projectrefuge-bootstrapsamplingRtoGo
Bootstrap sampling implementations in R and Go


## R

```bash
projectrefuge-bootstrapsamplingRtoGo % Rscript bootstrap_sampling.R
Standard Error of the Median for n_bootstrap= 10000 :
Positively Skewed:  0.003944053 
Symmetric:  0.004468373 
Negatively Skewed:  0.003742951 
Runtimes:
Generation time:  0.0004620552  seconds
Bootstrap sampling time:  1.239963  seconds
Calculation time:  0.00112915  seconds
ryano0oceros@Ryans-MacBook-Pro-2 projectrefuge-bootstrapsamplingRtoGo % Rscript bootstrap_sampling.R
Standard Error of the Median for n_bootstrap = 1e+05 :
Positively Skewed:  0.001252871 
Symmetric:  0.001414359 
Negatively Skewed:  0.00117999 
Runtimes:
Generation time:  0.0004689693  seconds
Bootstrap sampling time:  12.37505  seconds
Calculation time:  0.002007961  seconds
```

## Go

Sampling n = 10000
```bash
go run main.go
Standard Error of the Median for n_bootstrap=  10000 :
Distribution: Positively Skewed, Standard Error of the Median: 0.007382
Time taken for bootstrap sampling: 72.419416ms
Distribution: Symmetric, Standard Error of the Median: 0.071418
Time taken for bootstrap sampling: 80.96275ms
Distribution: Negatively Skewed, Standard Error of the Median: 0.007357
Time taken for bootstrap sampling: 77.609917ms
```

Sampling n = 100000
```bash
go run main.go
Standard Error of the Median for n_bootstrap=  100000 :
Distribution: Positively Skewed, Standard Error of the Median: 0.002343
Time taken for bootstrap sampling: 696.786792ms
Distribution: Symmetric, Standard Error of the Median: 0.022206
Time taken for bootstrap sampling: 650.631417ms
Distribution: Negatively Skewed, Standard Error of the Median: 0.002218
Time taken for bootstrap sampling: 642.173209ms
```

