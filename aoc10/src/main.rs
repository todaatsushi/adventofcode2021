use aoc10::part1::solve_part_1;
use aoc10::utils;

fn main() {
    let test = utils::is_test().unwrap();
    let input = utils::parse_input(test);

    solve_part_1(input);
}
