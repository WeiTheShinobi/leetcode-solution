impl Solution {
    pub fn count_triplets(arr: Vec<i32>) -> i32 {
        let mut result = 0;
        arr.iter().enumerate().take(arr.len()-1).for_each(|x| {
            let mut current = *x.1;
            arr.iter().skip(x.0 + 1).enumerate().for_each(|y| {
                current ^= *y.1;
                if current == 0 {
                    result += y.0 + 1;
                }
            })
        });

        result as i32
    }
}