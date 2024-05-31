impl Solution {
    pub fn single_number(nums: Vec<i32>) -> Vec<i32> {
        let mut xor = 0;
        nums.iter().for_each(|&x| xor ^= x);
        
        let right_most = xor & (-xor);
        nums.iter().fold(vec![0,0], |mut acc, x| {
            if x & right_most > 0 {
                acc[0] ^= x;
            } else {
                acc[1] ^= x;
            }
            acc
        })
    }
}