class Solution:
    def climbStairs(self, n: int) -> int:
        @cache
        def f(curr: int) -> int:
          if curr == 1:
            return 1
          if curr == 2:
            return 2
          return f(curr-2) + f(curr-1)

        return f(n)