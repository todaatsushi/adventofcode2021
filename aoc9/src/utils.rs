use std::env;

#[derive(Debug)]
pub enum InputError {
    MustHave2Args,
    ArgsMustBetrueOrfalse,
}

pub fn is_test() -> Result<bool, InputError> {
    let args: Vec<String> = env::args().collect();

    match args.len() as i32 {
        2 => {
            let test = match args[1].as_str() {
                "true" => Ok("true"),
                "false" => Ok("false"),
                _ => Err(InputError::ArgsMustBetrueOrfalse),
            };
            Ok(test.unwrap() == "true")
        }
        _ => Err(InputError::MustHave2Args),
    }
}

pub fn parse_input(test: bool) -> Vec<Vec<i32>> {
    let file_contents = match test {
        true => include_str!("../inputs/test.txt").lines(),
        false => include_str!("../inputs/puzzle.txt").lines(),
    };
    let mut points: Vec<Vec<i32>> = Vec::new();

    for line in file_contents {
        let line_points = line
            .split("")
            .filter(|&height| !height.is_empty())
            .map(|x| x.trim().parse::<i32>().unwrap())
            .collect();

        let _ = &points.push(line_points);
    }

    points
}
