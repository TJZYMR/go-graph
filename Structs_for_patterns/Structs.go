package Structs_for_patterns

type Patterns interface {
	Pattern() ([]int, []int, int, bool)
}

func Find(P Patterns) ([]int, []int, int, bool) {
	a, b, c, d := P.Pattern()
	return a, b, c, d
} //end of Find and with this adter that dfgfddfdfdf
