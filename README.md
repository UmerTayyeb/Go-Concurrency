# Go-Concurrency

This Go program demonstrates the difference in execution time between running tasks sequentially using a simple function and running them concurrently using goroutines. It simulates work with a delay and compares the time taken by each approach. The program includes:

- A simpleFunction that runs sequentially for a fixed number of iterations with a delay.
- A goroutineFunction that runs concurrently using goroutines, leveraging Go's concurrency model.
- Time measurements using Go's time package to compare the performance of both approaches.
- This example is ideal for understanding the impact of goroutines in improving concurrency and reducing the time taken for parallelizable tasks.