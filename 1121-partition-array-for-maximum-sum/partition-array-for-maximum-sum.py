class Solution:
    def maxSumAfterPartitioning(self, arr: List[int], k: int) -> int:
        @cache
        def dfs(i: int) -> int:
            res = mx = 0
            for j in range(i, max(i - k, -1), -1):
                mx = max(mx, arr[j])
                res = max(res, dfs(j - 1) + (i - j + 1) * mx)
            return res

        return dfs(len(arr) - 1)
