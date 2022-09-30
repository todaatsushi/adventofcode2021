use std::collections::HashMap;
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

pub fn parse_input(test: bool) -> Vec<Vec<&'static str>> {
    let file_contents = match test {
        true => include_str!("../inputs/test.txt").lines(),
        false => include_str!("../inputs/puzzle.txt").lines(),
    };
    let mut lines: Vec<Vec<&str>> = Vec::new();

    for line in file_contents {
        let l = line
            .split("")
            .map(|l| l.trim())
            .filter(|&c| !c.is_empty())
            .collect::<Vec<&str>>();
        lines.push(l);
    }
    lines
}

pub fn build_closing_bracket_mapper() -> HashMap<&'static str, &'static str> {
    let mut map = HashMap::new();

    map.insert("{", "}");
    map.insert("[", "]");
    map.insert("(", ")");
    map.insert("<", ">");
    map
}
