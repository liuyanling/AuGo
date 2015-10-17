package sort

type MP struct {
	value int
}

/**
 * if p.value less than b.value return true
 * else return false
 */
func (p *MP) lessThan(b MP) bool {
	return p.value < b.value
}