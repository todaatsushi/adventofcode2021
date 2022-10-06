use aoc12::input;
use aoc12::solve;

fn main() {
    let caves = input::parse_file_to_map();
    let num_paths_part1 = solve::get_num_paths(&caves, false);
    let num_paths_part2 = solve::get_num_paths(&caves, true);

    println!("Part 1: {:?}", num_paths_part1);
    println!("Part 2: {:?}", num_paths_part2);
}
