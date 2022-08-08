package arrays

type IStringArray struct {
	Append func(s string) string
}

func StringArray(str string) *IStringArray {
	return &IStringArray{
		Append: func(s string) string {
			str = str + s

			return str
		},
	}
}