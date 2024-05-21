impl Solution {
    pub fn subsets(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut result = vec![];

        fn f(i: usize, v: &mut Vec<i32>, result: &mut Vec<Vec<i32>>, nums: &Vec<i32>) {
            result.push(v.clone());

            for j in i..nums.len() {
                v.push(nums[j]);
                f(j+1, v, result, nums);
                v.pop();
            }
        }

        f(0, &mut vec![], &mut result, &nums);
        result
    }
}