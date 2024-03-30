package morenumbers

import "fmt"

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func (b ByteSize) String() string {
	switch {
	case b >= PB:
		return fmt.Sprintf("%.2f PB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2f TB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2f GB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2f MB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2f KB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

// `<<` == "times 2"; `>>` == "divided by 2"
// `n << x` == "n times 2^x"
// `n >> x` == "n divided 2^y"
func LeftShitf(n, x int) {
	leftShift := n << x
	fmt.Printf("%d << %d == %d(%d^%d) == %d\n", n, x, n, 2, x, leftShift)
}

func RightShift(n, x int) {
	rightShift := n >> x
	fmt.Printf("%d << %d == %d/(%d^%d) == %d\n", n, x, n, 2, x, rightShift)
}
