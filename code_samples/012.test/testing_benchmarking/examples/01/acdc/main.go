// package acdc asks
package acdc

// sum adds unlimited number of values of type int
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
