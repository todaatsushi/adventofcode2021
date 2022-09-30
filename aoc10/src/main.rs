use aoc10::utils;

fn main() {
    let test = utils::is_test().unwrap();
    let input = utils::parse_input(test);
    println!("{:?}", input);
}
