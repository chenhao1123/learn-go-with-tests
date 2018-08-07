package sum

//Sum will take an array of numbers and return the total.
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

//SumAll will take a varying number of slices, returning a new slice containing the totals for each slice pass in.
func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

//SumAllTails calculates the totals of the "tails" of each slice.
func SumAllTails(numbersToSumTails ...[]int) (sums []int) {
	for _, numbers := range numbersToSumTails {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return
}
