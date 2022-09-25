use aoc9::part1::solve_part_1;
use aoc9::part2::solve_part_2;
use aoc9::utils::is_test;

fn main() {
    let test = is_test().unwrap();
    solve_part_1(test);
    solve_part_2(test);
}
