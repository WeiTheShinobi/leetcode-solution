class Solution
{
public:
    int minSteps(int n)
    {
        if (n == 0 || n == 1)
            return 0;

        std::vector<int> v(n + 1);
        v[0] = 0;
        v[1] = 0;
        v[2] = 2;
        for (int i = 3; i <= n; i++)
        {
            v[i] = i;
            int j = i / 2;
            while (j >= 1)
            {
                if (i % j == 0)
                {
                    v[i] = std::min(v[i] , v[j] + (i / j)) ;
                }
                j--;
            }
        }

        return v[n];
    }
};