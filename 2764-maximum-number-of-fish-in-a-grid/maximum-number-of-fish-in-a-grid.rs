impl Solution {
    pub fn find_max_fish(grid: Vec<Vec<i32>>) -> i32 {
        let mut result = 0;

        let mut seen = vec![vec![false; grid[0].len()]; grid.len()];
        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                let d = Self::dfs(i, j, &grid, &mut seen);
                result = if result > d {result} else {d};
            }
        }

        result
    }

    fn dfs(i: usize, j: usize, graph: &Vec<Vec<i32>>, seen: &mut Vec<Vec<bool>>) -> i32 {
        if i >= graph.len() || j >= graph[0].len() || i < 0 || j < 0 {return 0;}
        if seen[i][j] {return 0;}
        if graph[i][j] == 0 {return 0;}
        seen[i][j] = true;
        return graph[i][j] +
            Self::dfs(i+1, j, graph, seen) +
            Self::dfs(i-1, j, graph, seen) +
            Self::dfs(i, j+1, graph, seen) +
            Self::dfs(i, j-1, graph, seen);
    }
}