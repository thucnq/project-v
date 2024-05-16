package number

func Binomial(n, k int) int {
	if n < 0 || k < 0 {
		panic("Binomial - negative input")
	}
	if n < k {
		panic("Binomial - bad set size")
	}
	// (n,k) = (n, n-k)
	if k > n/2 {
		k = n - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}
