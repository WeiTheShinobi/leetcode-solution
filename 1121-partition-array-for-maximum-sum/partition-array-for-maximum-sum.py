class Solution:
    def maxSumAfterPartitioning(self, arr: List[int], k: int) -> int:
        dp = [0] * (len(arr) + 1)
        for i in range(len(arr)):
            mx = 0
            for j in range(i, max(i - k, -1), -1):
                mx = max(mx, arr[j])
                dp[i+1] = max(dp[i+1], dp[j]+(i-j+1)*mx)

        return dp[len(arr)]