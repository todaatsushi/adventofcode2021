use crate::data::{Polymer, Rule};
use std::env;

#[derive(Debug)]
enum InputError {
    FileDoesntExist,
    Needs2Args,
}

fn read_file() -> &'static str {
    let args: Vec<String> = env::args().collect();
    let content = match args.len() as u8 {
        2 => match args[1].as_str() {
            "test" => Ok(include_str!("../inputs/test.txt")),
            "puzzle" => Ok(include_str!("../inputs/puzzle.txt")),
            _ => Err(InputError::FileDoesntExist),
        },
        _ => Err(InputError::Needs2Args),
    }
    .unwrap();

    content
}

pub fn read_input() -> Polymer {
    let (sequence, rules_str) = read_file().split_once("\n\n").unwrap();
    let rules: Vec<Rule> = rules_str
        .lines()
        .map(str::parse)
        .map(Result::unwrap)
        .collect();
    let sequence = sequence.to_string();

    Polymer::new(sequence, rules)
}
