package inslice

// Int compares an item to a slice of items and return indexes of the first X matched entries.
func Int(slice []int, item int, count int) (index []int) {
	for i, r := range slice {
		if len(index) == count {
			break
		}
		if r == item {
			index = append(index, i)
		}
	}
	return
}

// IntVar compares an item to a slice of item pointers and return indexes of the first X matched entries.
func IntVar(slice []*int, item int, count int) (index []int) {
	for i, r := range slice {
		if len(index) == count {
			break
		}
		if *r == item {
			index = append(index, i)
		}
	}
	return
}

// IntAll acts like Int, looking for all occurrences.
func IntAll(slice []int, item int) (index []int) {
	for i, r := range slice {
		if r == item {
			index = append(index, i)
		}
	}
	return
}

// IntVarAll acts like IntVar, looking for all occurrences.
func IntVarAll(slice []*int, item int) (index []int) {
	for i, r := range slice {
		if *r == item {
			index = append(index, i)
		}
	}
	return
}
