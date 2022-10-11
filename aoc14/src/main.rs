use aoc14::input::read_input;
use std::env;

fn main() {
    let mut polymer = read_input();

    let args: Vec<String> = env::args().collect();
    let num_steps: u64 = str::parse(args[2].as_str()).unwrap();

    polymer.run_steps(num_steps);
    polymer.count_chars();
}
