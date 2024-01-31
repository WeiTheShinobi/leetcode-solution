impl Solution {
    pub fn daily_temperatures(temperatures: Vec<i32>) -> Vec<i32> {
        let mut result = vec![0; temperatures.len()];
        let mut stk = vec![];

        for (i, &t) in temperatures.iter().enumerate() {
            let i = i as i32;
            if stk.is_empty() {
                stk.push(i);
            } else {
                while !stk.is_empty() && temperatures[stk[stk.len() - 1] as usize] < t {
                    let x = stk.pop().unwrap();
                    result[x as usize] = i - x;
                }
                stk.push(i);
            }
        }

        result
    }
}