impl Solution {
    pub fn max_satisfied(customers: Vec<i32>, grumpy: Vec<i32>, minutes: i32) -> i32 {
        let all_customers = customers.iter().sum();
        let minutes = minutes as usize;
        if minutes >= customers.len() {
            return all_customers;
        }
        let lost_customers = (0..customers.len()).fold(0, |acc, x| acc + customers[x] * grumpy[x]);

        let mut window = (0..minutes).fold(0, |acc, x| acc + customers[x] * grumpy[x]);
        let mut save_customers = window;
        (minutes..customers.len()).for_each(|i| {
            window = window + customers[i] * grumpy[i] - customers[i - minutes] * grumpy[i - minutes];
            save_customers = std::cmp::max(save_customers, window);
        });
        all_customers - lost_customers + save_customers
    }
}