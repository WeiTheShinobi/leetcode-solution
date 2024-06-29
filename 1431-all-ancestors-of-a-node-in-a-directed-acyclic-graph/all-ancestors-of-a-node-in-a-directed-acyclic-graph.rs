use std::collections::VecDeque;

impl Solution {
    pub fn get_ancestors(n: i32, edges: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let n = n as usize;
        let mut graph  = vec![None  ; n];
        let mut result = vec![vec![]; n];
        let mut seen   = vec![false ; n];
        let mut queue  = VecDeque::new();

        for edge in edges {
            let from = edge[0] as usize;
            let to   = edge[1] as usize;
            graph[from].get_or_insert_with(Vec::new).push(to);
        }
        for v in 0..n {
            queue.push_back(v);
            while let Some(u) = queue.pop_front() {
                if let Some(adj) = &graph[u] {
                    for &u in adj {
                        if !seen[u] {
                            seen[u] = true;
                            result[u].push(v as i32);
                            queue.push_back(u);
                        }
                    }
                }
            }
            seen.fill(false);
        }
        result
    }
}