package arrays

type IAnyArray struct {
	// Appends an item to the array. Returns the new array
	Append func(value any) []any
	// Returns the value specified by the index.
	At func(idx int) any
	// The Concat() method is used to merge the given array into this array. This method does not change the existing arrays, but instead returns a new array.
	Concat func(array []any) []any
	// The CopyWithin() method copies the value specified by the given index and appends it at the end.
	CopyWithin func(idx int) []any
	// Searches through the array for an item based on the given value.
	FindIndex func(value any)
}

func AnyArray(base_value any) *IAnyArray {
	var (
		store    = []any{base_value}
		appender = func(value any) []any {
			var newStore = store
			newStore = append(newStore, value)

			return newStore
		}
	)

	return &IAnyArray{
		Append: appender,
		At: func(idx int) any {
			return store[idx]
		},
		Concat: func(array []any) []any {
			var newArr = store

			for _, v := range array {
				newArr = append(newArr, v)
			}

			return newArr
		},
		CopyWithin: func(idx int) []any {
			var item = store[idx]

			return appender(item)
		},
		// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/entries
	}
}
