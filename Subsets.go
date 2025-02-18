func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	amount := len(nums)
	for amount >= 2 {
        static := make([]int,0,amount)
        for x:=0;x<amount-1;x++{
            static = append(static,x)
        }
        for;;{
            last:= static[len(static)-1] +1
            for last<len(nums){
                temp:= make([]int,0)
                for _,y:=range(static){
				    temp = append(temp, nums[y])
                }
				temp = append(temp, nums[last])
                last++
            result = append(result,temp)
            }
            flag:= false
            for z:=len(static)-1;z>=0;z--{
                if static[z]+1+len(static)-z<len(nums){
                    static[z] += 1
                    for w:=z+1;w<len(static);w++{
                        static[w] = static[z]+w-z
                    }
                    flag = true
                    break
                }
            }
            if flag == false{
                break
            }


        }

		amount--
	}

	for _, x := range nums {
		result = append(result, []int{x})
	}
	result = append(result, []int{})
	return result
}