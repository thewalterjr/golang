func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
        temp := (num & 1)
        if temp == 1 {
            count++
        }
        num = num >> 1
	}
	return count
}

func hammingWeight(num uint32) int {
	aux := 0
	for num != 0 {
		aux += int(num & 1) 
		num >>= 1 
	}
	return aux
}
