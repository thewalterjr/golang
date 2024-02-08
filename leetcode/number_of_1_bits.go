func hammingWeight(num uint32) int {
	aux := 0
	for num != 0 {
		aux += int(num & 1) 
		num >>= 1 
	}
	return aux
}
