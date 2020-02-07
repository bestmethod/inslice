package inslice

// Uint compares an item to a slice of items and return indexes of the first X matched entries.
func Uint(slice []uint, item uint, count int) (index []int) {
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

// UintVar compares an item to a slice of item pouinters and return indexes of the first X matched entries.
func UintVar(slice []*uint, item uint, count int) (index []int) {
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

// UintAll acts like Uint, looking for all occurrences.
func UintAll(slice []uint, item uint) (index []int) {
	for i, r := range slice {
		if r == item {
			index = append(index, i)
		}
	}
	return
}

// UintVarAll acts like UintVar, looking for all occurrences.
func UintVarAll(slice []*uint, item uint) (index []int) {
	for i, r := range slice {
		if *r == item {
			index = append(index, i)
		}
	}
	return
}
