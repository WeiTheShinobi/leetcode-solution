func subarraysDivByK(nums []int, k int) int {
    record := map[int]int{0: 1}
    sum, ans := 0, 0
    for _, elem := range nums {
        sum += elem
        modulus := (sum % k + k) % k
        ans += record[modulus]
        record[modulus]++
    } 
    return ans
}
