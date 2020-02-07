package inslice

// String compares an item to a slice of items and return indexes of the first X matched entries.
func String(slice []string, item string, count int) (index []int) {
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

// StringVar compares an item to a slice of item pointers and return indexes of the first X matched entries.
func StringVar(slice []*string, item string, count int) (index []int) {
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

// StringAll acts like String, looking for all occurrences.
func StringAll(slice []string, item string) (index []int) {
	for i, r := range slice {
		if r == item {
			index = append(index, i)
		}
	}
	return
}

// StringVarAll acts like StringVar, looking for all occurrences.
func StringVarAll(slice []*string, item string) (index []int) {
	for i, r := range slice {
		if *r == item {
			index = append(index, i)
		}
	}
	return
}
