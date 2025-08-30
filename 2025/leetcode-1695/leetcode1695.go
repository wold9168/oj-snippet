package main
import "fmt"
func maximumUniqueSubarray(nums []int) int {
    mp := make(map[int]int)
    ans,sum := 0,0
    l := 0
    for i,_ := range nums{
        num := nums[i]
        ind,exist := mp[num]
        if !exist {
            sum += num
            ans = max(ans,sum)
            mp[num]=i
        } else {
            for j:=l;j<=ind;j++{
                delete(mp,nums[j])
                sum -= nums[j]
            }
            sum += num
            l = ind+1
        }
			fmt.Println(sum,ans,num,i,l,mp)
    }
    return ans
    
}
func main() {
	maximumUniqueSubarray([]int{5,2,1,2,5,2,1,2,5})
}
