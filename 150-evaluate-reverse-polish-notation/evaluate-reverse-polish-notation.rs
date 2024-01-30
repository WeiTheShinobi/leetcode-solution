impl Solution {
       pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        let mut stack = Vec::new();

        for token in tokens {
            if let Ok(val) = token.parse::<i32>() {
                stack.push(val);
            } else {
                let a = stack.pop().unwrap();
                let b = stack.pop().unwrap();
                let result = match token.as_str() {
                    "+" => b + a,
                    "-" => b - a,
                    "*" => b * a,
                    "/" => b / a,
                    _ => panic!(),
                };
                stack.push(result);
            }
        }

        stack[0]
    }
}