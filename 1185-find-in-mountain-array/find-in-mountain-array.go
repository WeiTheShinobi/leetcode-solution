func findInMountainArray(target int, mountainArr *MountainArray) int {
	l, r := 0, mountainArr.length()
	for l+1 != r {
		mid := (l + r) / 2
		mV := mountainArr.get(mid)
		mV2 := mountainArr.get(mid + 1)
		if mV > mV2 {
			r = mid
		} else {
			l = mid
		}
	}

	peer := r
	l, r = -1, r+1
	for l+1 != r {
		mid := (l + r) / 2
		if mountainArr.get(mid) <= target {
			l = mid
		} else {
			r = mid
		}
	}

	if l != -1 && mountainArr.get(l) == target {
		return l
	}

	l, r = peer-1, mountainArr.length()
	for l+1 != r {
		mid := (l + r) / 2
		if mountainArr.get(mid) >= target {
			l = mid
		} else {
			r = mid
		}
	}
	
	if l != -1 && mountainArr.get(l) != target {
		return -1
	}
	return l
}
