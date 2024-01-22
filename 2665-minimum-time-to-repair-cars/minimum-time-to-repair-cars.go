func repairCars(ranks []int, cars int) int64 {
	greatest := 101
	for _, rank := range ranks {
		greatest = min(greatest, rank)
	}

	return int64(sort.Search(greatest*cars*cars, func(i int) bool {
		s := 0
		for _, rank := range ranks {
			s += int(math.Sqrt(float64(i / rank)))
		}
		return s >= cars
	}))
}