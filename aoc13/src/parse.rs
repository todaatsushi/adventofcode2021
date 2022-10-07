use crate::data;
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

pub fn read_input() -> (Vec<data::Point>, Vec<data::Fold>) {
    let (points, folds) = read_file().split_once("\n\n").unwrap();

    let points: Vec<data::Point> = points.lines().map(str::parse).map(Result::unwrap).collect();
    let folds: Vec<data::Fold> = folds.lines().map(str::parse).map(Result::unwrap).collect();

    (points, folds)
}
