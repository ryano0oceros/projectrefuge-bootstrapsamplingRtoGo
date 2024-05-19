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
gen_start_time <- Sys.time()
set.seed(123)
positively_skewed <- rexp(1000, rate = 0.1)
symmetric <- rnorm(1000, mean = 50, sd = 10)
negatively_skewed <- sort(rexp(1000, rate = 0.1), decreasing = TRUE)
gen_end_time <- Sys.time()
gen_time <- gen_end_time - gen_start_time

# Bootstrap sampling
n_bootstrap <- 10000
pos_boot_start_time <- Sys.time()
bootstrap_positively_skewed <- bootstrap_sampling(positively_skewed, n_bootstrap)
pos_boot_end_time <- Sys.time()
pos_bootstrap_time <- pos_boot_end_time - pos_boot_start_time

sym_boot_start_time <- Sys.time()
bootstrap_symmetric <- bootstrap_sampling(symmetric, n_bootstrap)
sym_boot_end_time <- Sys.time()
sym_bootstrap_time <- sym_boot_end_time - sym_boot_start_time

neg_boot_start_time <- Sys.time()
bootstrap_negatively_skewed <- bootstrap_sampling(negatively_skewed, n_bootstrap)
neg_boot_end_time <- Sys.time()
neg_bootstrap_time <- neg_boot_end_time - neg_boot_start_time

# Calculating and printing standard error of the median
se_calc_start_time <- Sys.time()
se_positively_skewed <- standard_error(bootstrap_positively_skewed)
se_symmetric <- standard_error(bootstrap_symmetric)
se_negatively_skewed <- standard_error(bootstrap_negatively_skewed)
se_calc_end_time <- Sys.time()
se_calc_time <- se_calc_end_time - se_calc_start_time

cat("Standard Error of the Median for n_bootstrap =", n_bootstrap, ":\n")
cat("Positively Skewed: ", se_positively_skewed, "\n")
cat("Symmetric: ", se_symmetric, "\n")
cat("Negatively Skewed: ", se_negatively_skewed, "\n")

cat("Runtimes:\n")
cat("Generation time: ", gen_time, " seconds\n")
cat("Positively Skewed Bootstrap sampling time: ", pos_bootstrap_time, " seconds\n")
cat("Symmetric Bootstrap sampling time: ", sym_bootstrap_time, " seconds\n")
cat("Negatively Skewed Bootstrap sampling time: ", neg_bootstrap_time, " seconds\n")
cat("Calculation time: ", se_calc_time, " seconds\n")

