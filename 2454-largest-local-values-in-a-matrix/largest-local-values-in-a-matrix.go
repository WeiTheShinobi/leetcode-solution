func largestLocal(grid [][]int) [][]int {
    res := make([][]int, len(grid) - 2)
    for idx := range res {
        res[idx] = make([]int, len(grid) - 2)
    }
    
    for i:=0;i<len(res);i++{
        for j:=0;j<len(res[0]);j++{
            maxValue := getMaximum(grid, i, j)
            res[i][j] = maxValue
        }
    }
    
    return res
}

func getMaximum(grid [][]int, x, y int) int{
    maxValue := -1
    for i:=x;i<x+3;i++{
        for j:=y;j<y+3;j++{
            if grid[i][j] > maxValue{
                maxValue = grid[i][j]
            }
        }
    }
    return maxValue
}