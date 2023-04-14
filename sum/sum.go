package sum


func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers{
		sum += num
	}
	return sum
}
// sum all nums excluding the first num
func SumAllTails(numsToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numsToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
		
	}
	return sums
}
