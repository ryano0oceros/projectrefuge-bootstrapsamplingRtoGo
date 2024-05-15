# Bootstrap Sampling in R

# Function to perform bootstrap sampling
bootstrap_sampling <- function(data, n_bootstrap) {
  sample_means <- numeric(n_bootstrap)
  for (i in 1:n_bootstrap) {
    sample <- sample(data, length(data), replace = TRUE)
    sample_means[i] <- median(sample)
  }
  return(sample_means)
}

# Function to calculate standard error
standard_error <- function(sample_means) {
  return(sd(sample_means) / sqrt(length(sample_means)))
}

# Generating distributions
set.seed(123)
positively_skewed <- rexp(1000, rate = 0.1)
symmetric <- rnorm(1000, mean = 50, sd = 10)
negatively_skewed <- sort(rexp(1000, rate = 0.1), decreasing = TRUE)

# Bootstrap sampling
n_bootstrap <- 10000
bootstrap_positively_skewed <- bootstrap_sampling(positively_skewed, n_bootstrap)
bootstrap_symmetric <- bootstrap_sampling(symmetric, n_bootstrap)
bootstrap_negatively_skewed <- bootstrap_sampling(negatively_skewed, n_bootstrap)

# Calculating and printing standard error of the median
se_positively_skewed <- standard_error(bootstrap_positively_skewed)
se_symmetric <- standard_error(bootstrap_symmetric)
se_negatively_skewed <- standard_error(bootstrap_negatively_skewed)

cat("Standard Error of the Median:\n")
cat("Positively Skewed: ", se_positively_skewed, "\n")
cat("Symmetric: ", se_symmetric, "\n")
cat("Negatively Skewed: ", se_negatively_skewed, "\n")