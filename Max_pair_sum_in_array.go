func findLargest(nums []int) []int {
    biggerDigits:= make([]int,len(nums),len(nums))
	for pos,x := range nums {
		y := strconv.Itoa(x)
        max:= y[0]-'0'
        for _,z:= range(y){
            if z-'0'>max{
                max = z-'0'
            }
        }
        biggerDigits[pos] = max
        
	}
	return biggerDigits
}
func maxSum(nums []int) int {
	bigger:= findLargest(nums)
    fmt.Println(bigger)
	return nums[0]
}