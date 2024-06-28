impl Solution {
    pub fn maximum_importance(n: i32, roads: Vec<Vec<i32>>) -> i64 {
        let mut connections = roads.into_iter().fold(vec![0i64; n as usize], |mut acc, x| { 
            acc[x[0] as usize] += 1; acc[x[1] as usize] += 1; acc 
        });
        connections.sort_unstable();
        connections.iter().enumerate().fold(0, |mut acc, (i, x)| { acc += x * (i as i64 + 1); acc })
    }
}