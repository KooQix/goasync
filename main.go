package goasync

import (
	"sync"
)


var wc sync.WaitGroup

/*
	Go is a function that takes a callback function and runs it in a goroutine. It also increments the WaitGroup counter. This is useful when you want to run a function in a goroutine and wait for it to finish before exiting the program.

	Note that the WaitGroup counter is decremented when the callback function finishes executing.

	If you do not use `goasync.RunMain` function, or if you do not call `goasync.WaitAll` function, `goasync.Go` function will not wait for the goroutine to finish before exiting, therefore behaving like a normal goroutine.

	Example usage:
	
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
	
*/
func Go(callback func()) {
	wc.Add(1)
	go func() {
        defer wc.Done()
        callback()
    }()

}


/*
	WaitAll is a function that waits for all goroutines to finish before exiting the program. This is useful when you want to wait for all goroutines to finish before exiting the program.

	Example usage:

		package main

		import (
			"github.com/kooqix/goasync"
		)

		func main() {
			goasync.Go(func() {
				// Your async code here
			})

			// Wait for all goroutines to finish before exiting
			goasync.WaitAll()
		}

*/
func WaitAll() {
	wc.Wait()
}


/*
	RunMain is a Wrapper for main. That takes the main function and runs it. It also waits for all goroutines to finish before exiting the program. It is the equivalent of calling `goasync.WaitAll` at the end of the main function.

	Example usage: 

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
*/
func RunMain(mainFunc func()) {
	defer WaitAll()
	mainFunc()
}