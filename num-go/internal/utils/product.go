package utils

type Number interface {
    int | float64
}

func Product[T Number](numbers []T) T {
	var ans T = 1
	for _, num := range numbers {
		ans = ans*num
	}
	return ans
} 