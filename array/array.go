package array

func New[T comparable]() *Array[T] {
	return &Array[T]{
		data: []T{},
	}
}

func removeDuplicates[T comparable](arr []T) []T {
	seen := make(map[T]bool)
	result := []T{}

	for _, item := range arr {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// Creates and returns a new array. Using the "existing" parameter as a base/beginning.
func NewWithBase[T comparable](remove_duplicates bool, existing ...T) *Array[T] {
	var new_arr = []T{}

	new_arr = append(new_arr, existing...)
	new_arr = removeDuplicates(new_arr)
	
	return &Array[T]{
		data: new_arr,
	}
}