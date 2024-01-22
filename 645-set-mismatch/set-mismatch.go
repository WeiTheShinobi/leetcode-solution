func findErrorNums(nums []int) []int {
	n := len(nums)
	found := make([]bool,n + 1)
	repeatedNum := 0
	runningSum := 0
	desiredSum := (n * (n + 1)) / 2
	for _,thisNum := range nums {
		if found[thisNum] {
			if repeatedNum == 0 {
				repeatedNum = thisNum
			}
		} else {
			found[thisNum] = true
			runningSum += thisNum
		}
	}
	missingNum := desiredSum - runningSum
	return []int{repeatedNum,missingNum}
}