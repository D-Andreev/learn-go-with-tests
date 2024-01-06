package iteration

const defaultRepeatCount = 5

func Repeat(character string, repeatCount int) string {
	var repeated string

	if repeatCount < 0 {
		repeatCount = defaultRepeatCount
	}

	for i := 0; i < repeatCount; i++ {
		repeated = repeated + character
	}

	return repeated
}
