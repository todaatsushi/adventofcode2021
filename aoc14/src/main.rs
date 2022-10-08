use aoc14::input::read_input;

fn main() {
    let mut polymer = read_input();
    polymer.run_steps(10);

    let max = polymer.tally.values().max().unwrap();
    let min = polymer.tally.values().min().unwrap();
    println!("Part 1: {}", max - min);
}
