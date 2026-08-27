package main

import (
	"fmt"
	"math/rand"

	"github.com/Davidmuthee12/threads/waitgroups-and-barriers/barriers"
)

const matrixsize = 3

// Generate a matrix using random integers
func generateRandMatrix(matrix *[matrixsize][matrixsize]int) {
	for row := 0; row < matrixsize; row++ {
		for col := 0; col < matrixsize; col++ {
			// for every row and column assign a 
			// random number between -5 and 4
			matrix[row][col] = rand.Intn(10) - 5
		}
	}
}

func matrixMultiply(matrixA, matrixB, result *[matrixsize][matrixsize]int) {
	// iterates over every row
	for row := 0; row < matrixsize; row++ {
		// iterates over every column
		for col := 0; col < matrixsize; col++ {
			sum := 0
			for i := 0; i < matrixsize; i++ {
				// Sums up each value of the row from A
				//  multiplied by each value of the column from B
				sum += matrixA[row][i] * matrixB[i][col]
			}
			// Update the result matrix with the sum
			result[row][col] = sum
		}
	}
}

func rowMultiply(matrixA, matrixB, result *[matrixsize] [matrixsize] int, row int, barrier *barriers.Barrier) {
	// Start an infinite loop
	for {
		// Waits on the barrier untill the main()
		// goroutine loads the matrices
		barrier.Wait()
		for col := 0; col < matrixsize; col++ {
			sum := 0
			for i := 0; i < matrixsize; i++ {
				// Calculates the result of the row in this goroutine
				sum += matrixA[row][i]  * matrixB[i][col]
			}
			// Assign the result to the correct row and column
			result[row] [col] = sum
		}
		// waits on the barrier untill every other row has been computed
		barrier.Wait()
	}
}

func main() {
	var matrixA, matrixB, result [matrixsize] [matrixsize] int
	// Creates a new barrier with size of row
	// goroutines + main() goroutine
	barrier := barriers.NewBarrier(matrixsize + 1)
	for row := 0; row < matrixsize; row++ {
		go rowMultiply(&matrixA, &matrixB, &result, row, barrier)
	}

	for i := 0; i < 4; i++ {
		// loads up both matrices by randomly generating them
		generateRandMatrix(&matrixA)
		generateRandMatrix(&matrixB)

		// Releases the barrier so the
		//  goroutines can start their computation
		barrier.Wait()

		// waits until the gotoutines finish their computations
		barrier.Wait()

		for i := 0; i < matrixsize; i++ {
			// outputs results to the console
			fmt.Println(matrixA[i], matrixB[i], result[i])
		} 
		fmt.Println()
	}
}