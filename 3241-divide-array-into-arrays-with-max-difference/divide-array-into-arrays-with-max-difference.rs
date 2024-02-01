impl Solution {
    pub fn divide_array(mut nums: Vec<i32>, k: i32) -> Vec<Vec<i32>> {
        let mut res = Vec::with_capacity((nums.len() + 2) / 3);
        nums.sort_unstable();
        nums.chunks_exact(3)
            .take_while(|chunk| chunk.len() == 3)
            .filter(|chunk| chunk[2] - chunk[0] <= k)
            .for_each(|chunk| res.push(chunk.to_vec()));
        
        if res.len() * 3 != nums.len() {
            return Vec::new();
        }
        res
    }
}