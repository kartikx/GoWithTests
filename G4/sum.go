package main

// Sum returns the Sum of an Array
func Sum(numbers []int) (sum int) {
	for _, num := range(numbers) {
		sum += num
	}

	return
}

// SumAllTails returns a new slice containing the individual sums of each input slice.
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range(numbersToSum) {
		sum := 0
		
		if len(numbers) > 0 {
			sum = Sum(numbers) - numbers[0]
		}

		sums = append(sums, sum)
	}

	return
}

func main() {

}