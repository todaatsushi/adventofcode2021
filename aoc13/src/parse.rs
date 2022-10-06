use std::env;

#[derive(Debug)]
enum InputError {
    FileDoesntExist,
    Needs2Args,
}

fn read_file() -> &'static str {
    let args: Vec<String> = env::args().collect();
    let lines = match args.len() as u8 {
        2 => match args[1].as_str() {
            "test" => Ok(include_str!("../inputs/test.txt")),
            "puzzle" => Ok(include_str!("../inputs/puzzle.txt")),
            _ => Err(InputError::FileDoesntExist),
        },
        _ => Err(InputError::Needs2Args),
    }
    .unwrap();

    lines
}
