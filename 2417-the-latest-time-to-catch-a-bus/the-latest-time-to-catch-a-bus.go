func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)

	var result int
	j, c := 0, 0
	for _, t := range buses {
		for c = capacity; c > 0 && j < len(passengers) && passengers[j] <= t; c-- {
			j++
		}
	}
	if c > 0 {
		result = buses[len(buses)-1]
	} else {
		result = passengers[j-1]
	}
	for j--; j >= 0 && result == passengers[j]; j-- {
		result--
	}

	return result
}