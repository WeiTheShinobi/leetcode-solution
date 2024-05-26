impl Solution {
    const MOD: u32 = (1e9 + 7.) as _;

    pub fn check_record(n: i32) -> i32 {
        let sum = |arr: &[u32]| arr.iter().sum::<u32>() % Self::MOD;
        (0..n)
            .fold([[1, 0, 0], [0, 0, 0]], |prev, _| {
                let [s1, s2] = [sum(&prev[0]), sum(&prev[1])];
                [
                    [s1, prev[0][0], prev[0][1]],
                    [sum(&[s1, s2]), prev[1][0], prev[1][1]],
                ]
            })
            .into_iter()
            .flatten()
            .fold(0, |s, n| (s + n) % Self::MOD) as _
    }
}