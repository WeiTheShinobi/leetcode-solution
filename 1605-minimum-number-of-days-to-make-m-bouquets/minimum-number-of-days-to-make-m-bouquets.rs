impl Solution {
    pub fn min_days(bloom_day: Vec<i32>, m: i32, k: i32) -> i32 {
        if (bloom_day.len() as i32) < m * k || m * k < 0{
            return -1;
        }

        let max_days = *(bloom_day.iter().max().unwrap());
        let (mut l, mut r) = (-1, max_days+1);

        let check = |days: i32| {
            let mut f = 0;
            let mut bouquet = 0;
            for &i in &bloom_day {
                if days >= i {
                    f += 1;
                } else {
                    f = 0;
                }
                
                if f == k {
                    bouquet += 1;
                    f = 0;
                }
                if bouquet >= m {
                    break;
                }
            }
            bouquet >= m
        };

        while l + 1 != r {
            let mid = (r + l) / 2;
            if check(mid) {
                r = mid
            } else {
                l = mid
            }
        }

        r
    }
}