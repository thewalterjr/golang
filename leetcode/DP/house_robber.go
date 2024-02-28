func rob(nums []int) int {
    rob1 := 0
    rob2 := 0

    for _, v := range nums {
        tmp := int(math.Max(float64(v+rob1), float64(rob2)))
        rob1 = rob2
        rob2 = tmp
    }
    return rob2
}
