package count_integers_with_event_digit_sum

// leetcode 2180
// https://leetcode.com/contest/weekly-contest-281/problems/count-integers-with-even-digit-sum/

func countEven(num int) int {
	var total int
	for i := 1; i <= num; i++ {
		var count int
		for x := i; x > 0; x /= 10 {
			count += x % 10
		}
		// even 0; odd 1
		if (count & 1) == 0 {
			total++
		}
	}
	return total
}
