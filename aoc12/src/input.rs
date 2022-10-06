use std::{collections::HashMap, env, str::Lines};

#[derive(Debug)]
enum InputError {
    FileDoesntExist,
    Needs2Args,
}

fn read_input() -> Lines<'static> {
    let args: Vec<String> = env::args().collect();
    let lines = match args.len() as u8 {
        2 => match args[1].as_str() {
            "test" => Ok(include_str!("../inputs/test.txt")),
            "test2" => Ok(include_str!("../inputs/test2.txt")),
            "test3" => Ok(include_str!("../inputs/test3.txt")),
            "puzzle" => Ok(include_str!("../inputs/puzzle.txt")),
            _ => Err(InputError::FileDoesntExist),
        },
        _ => Err(InputError::Needs2Args),
    }
    .unwrap()
    .lines();

    lines
}

pub fn parse_file_to_map() -> HashMap<&'static str, Vec<&'static str>> {
    let mut caves_map: HashMap<&str, Vec<&str>> = HashMap::new();
    let lines = read_input();

    for line in lines {
        let caves = line.split("-").collect::<Vec<&str>>();

        let to = caves[0];
        let from = caves[1];

        let to_edges = caves_map.entry(to).or_insert(Vec::new());
        to_edges.push(from);

        let from_edges = caves_map.entry(from).or_insert(Vec::new());
        from_edges.push(to);
    }
    caves_map
}
