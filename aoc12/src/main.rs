use aoc12::input;
use aoc12::solve;

fn main() {
    let caves = input::parse_file_to_map();
    let num_paths = solve::get_num_paths(caves);

    println!("{:?}", num_paths);
}
