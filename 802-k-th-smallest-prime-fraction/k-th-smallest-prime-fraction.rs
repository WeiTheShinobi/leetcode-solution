impl Solution {
    pub fn kth_smallest_prime_fraction(arr: Vec<i32>, k: i32) -> Vec<i32> {
        struct Pair {
            x: i32,
            y: i32,
            r: f32,
        }

        let mut r_arr = vec![];
        for i in 0..arr.len() {
            for j in i+1..arr.len() {
                let r = arr[i] as f32 / arr[j] as f32;
                r_arr.push(Pair {
                    x: arr[i],
                    y: arr[j],
                    r,
                });
            }
        }

        r_arr.sort_by(|a, b| a.r.partial_cmp(&b.r).unwrap());

        vec![r_arr[(k-1) as usize].x, r_arr[(k-1) as usize].y]
    }
}