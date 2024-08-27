# GoAsync Package

`goasync` is a simple Go package that helps manage goroutines and ensures that all goroutines complete before the program exits. It provides an easy-to-use API for running functions asynchronously and waiting for their completion.

## Features

-   **Run functions asynchronously**: The `Go` function runs a callback function in a goroutine and tracks its completion.
-   **Wait for all goroutines to finish**: The `WaitAll` function blocks until all previously started goroutines have finished.
-   **Automatic management of goroutines in `main`**: The `RunMain` function wraps the `main` function to ensure all goroutines complete before the program exits.

## Installation

To install the package, run:

```bash
go get github.com/kooqix/goasync
```

Then, import it in your Go files:

```go
import "github.com/kooqix/goasync"
```

## Usage

### 1. Running Functions Asynchronously

The `Go` function allows you to run a function in a goroutine. It increments an internal `sync.WaitGroup` counter, ensuring that the program waits for the goroutine to finish before exiting, provided you call `goasync.WaitAll` or use `goasync.RunMain`.

```go
package main

import (
	"fmt"
	"github.com/kooqix/goasync"
)

func myFunction() {
	// Create a goroutine to run this function
	goasync.Go(func() {
		fmt.Println("Hello, World!")
	})

	// Other code here
}
```

### 2. Waiting for All Goroutines to Complete

The `WaitAll` function blocks until all goroutines started with `goasync.Go` have completed. This is essential for ensuring that your program does not exit prematurely, leaving unfinished tasks.

```go
package main

import (
	"github.com/kooqix/goasync"
)

func main() {
	// Create a goroutine to run some async code
	goasync.Go(func() {
		// Your async code here
	})
	goasync.WaitAll() // Waits for all async tasks to complete before exiting
}
```

### 3. Automatic Waiting with `RunMain`

`RunMain` is a convenient wrapper for your `main` function. It ensures that all goroutines complete before the program exits, without requiring you to manually call `WaitAll`.

```go
package main

import "github.com/kooqix/goasync"

func main() {
	goasync.RunMain(func() {
		goasync.Go(func() {
			// Your async code here
		})

		// Other code here
	})
}
```

### 4. Example Use Case

Consider a scenario where you have multiple independent tasks that need to run concurrently. Using the `goasync` package, you can easily manage these tasks:

```go
package main

import (
	"fmt"
	"time"
	"github.com/kooqix/goasync"
)

func main() {
	goasync.RunMain(func() {
		goasync.Go(func() {
			time.Sleep(2 * time.Second)
			fmt.Println("Task 1 complete")
		})

		goasync.Go(func() {
			time.Sleep(1 * time.Second)
			fmt.Println("Task 2 complete")
		})

		fmt.Println("Main function complete")
	})
}
```

Output:

```plaintext
Main function complete
Task 2 complete
Task 1 complete
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with improvements or bug fixes.

## Contact

For any questions or suggestions, feel free to open an issue or contact [kooqix.dev@gmail.com](mailto:kooqix.dev@gmail.com).
