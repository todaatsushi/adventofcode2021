use aoc14::input::read_input;
use std::env;

fn main() {
    let mut polymer = read_input();
    let args: Vec<String> = env::args().collect();
    let num_steps: u32 = str::parse(args[2].as_str()).unwrap();

    polymer.run_steps(num_steps);

    let max = polymer.tally.values().max().unwrap();
    let min = polymer.tally.values().min().unwrap();
    println!("Part 1: {}", max - min);
}
