use crate::cave::Cave;
use std::{collections::HashMap, env};

#[derive(Debug)]
pub enum InputError {
    FileDoesntExist,
    NotEnoughArgs,
}

pub fn parse_input() -> HashMap<String, Cave> {
    let args: Vec<String> = env::args().collect();
    let content_lines = match args.len() as u8 {
        2 => match args[1].as_str() {
            "test" => Ok(include_str!("../inputs/test.txt")),
            "test2" => Ok(include_str!("../inputs/test2.txt")),
            "puzzle" => Ok(include_str!("../inputs/puzzle.txt")),
            _ => Err(InputError::FileDoesntExist),
        },
        _ => Err(InputError::NotEnoughArgs),
    }
    .unwrap()
    .lines();

    let mut caves: HashMap<String, Cave> = HashMap::new();
    for line in content_lines {
        let names = line.split("-").collect::<Vec<&str>>();

        let to_name = names[1];
        let from_name = names[0];

        let from_cave = caves.entry(from_name.to_string()).or_insert(Cave {
            name: from_name.to_string(),
            to: Vec::new(),
            from: Vec::new(),
        });
        from_cave.to.push(to_name.to_string());

        let to_cave = caves.entry(to_name.to_string()).or_insert(Cave {
            name: to_name.to_string(),
            to: Vec::new(),
            from: Vec::new(),
        });
        to_cave.from.push(from_name.to_string());
    }
    caves
}
