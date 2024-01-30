impl Solution {
    pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        let mut stack: Vec<String> = Vec::new();

        for token in tokens {
            match token.as_str() {
                "+" => {
                    let a = stack.pop().unwrap().parse::<i32>().unwrap();
                    let b = stack.pop().unwrap().parse::<i32>().unwrap();
                    stack.push((b+a).to_string())
                }
                "-" => {
                    let a = stack.pop().unwrap().parse::<i32>().unwrap();
                    let b = stack.pop().unwrap().parse::<i32>().unwrap();
                    stack.push((b-a).to_string())
                }
                "*" => {
                    let a = stack.pop().unwrap().parse::<i32>().unwrap();
                    let b = stack.pop().unwrap().parse::<i32>().unwrap();
                    stack.push((b*a).to_string())
                }
                "/" => {
                    let a = stack.pop().unwrap().parse::<i32>().unwrap();
                    let b = stack.pop().unwrap().parse::<i32>().unwrap();
                    stack.push((b/a).to_string())
                }
                _ => stack.push(token)
            }
        }

        stack[0].parse::<i32>().unwrap()
    }
}