use solutions::is_valid;
use solutions::two_sum;

mod solutions;

fn main() {
    // 01_two_sum
    println!();
    println!(">>> 01_two_sum <<<");
    println!("{:?}", two_sum(vec![2, 7, 11, 15], 9));
    println!("------------------");

    // 02_valid_parentheses
    println!();
    println!(">>> 02_valid_parentheses <<<");
    println!("{}", is_valid("()".to_string()));
    println!("{}", is_valid("()[]{}".to_string()));
    println!("{}", is_valid("(]".to_string()));
    println!("------------------");
}
