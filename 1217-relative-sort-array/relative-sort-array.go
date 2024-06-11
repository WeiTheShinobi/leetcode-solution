func relativeSortArray(arr1 []int, arr2 []int) []int {
	result := make([]int, 0)

	for _, v := range arr2 {
		for i := len(arr1) - 1; i >= 0; i-- {
			if v == arr1[i] {
				result = append(result, v)
				arr1 = append(arr1[:i], arr1[i+1:]...)
			}
		}
	}

	sort.Ints(arr1)
	return append(result, arr1...)
}