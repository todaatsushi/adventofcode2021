use aoc10::utils;
use std::collections::{HashMap, HashSet};

fn build_closing_bracket_mapper() -> HashMap<&'static str, &'static str> {
    let mut map = HashMap::new();

    map.insert("{", "}");
    map.insert("[", "]");
    map.insert("(", ")");
    map.insert("<", ">");
    map
}

fn build_score_mapper() -> HashMap<&'static str, i32> {
    let mut map = HashMap::new();

    map.insert(")", 3);
    map.insert("]", 57);
    map.insert("}", 1197);
    map.insert(">", 25137);
    map
}

fn main() {
    let test = utils::is_test().unwrap();
    let input = utils::parse_input(test);

    let closing_mapper = build_closing_bracket_mapper();
    let score_mapper = build_score_mapper();
    let opening_brackets = HashSet::from(["{", "(", "<", "["]);

    let mut expected: &str;
    let mut total = 0;

    for line in input {
        let mut bracket_stack: Vec<&str> = Vec::new();
        let mut first_illegal_character: Option<&str> = None;
        let mut subtotal = 0;

        for bracket in line {
            if opening_brackets.contains(bracket) {
                bracket_stack.push(closing_mapper.get(bracket).unwrap());
            } else {
                expected = match bracket_stack.pop() {
                    Some(b) => b,
                    None => panic!("bracket_stack empty"),
                };
                if expected != bracket {
                    first_illegal_character = match first_illegal_character {
                        Some(value) => Some(value),
                        None => Some(bracket),
                    };

                    if first_illegal_character.unwrap() == bracket {
                        subtotal += score_mapper.get(bracket).unwrap()
                    }
                }
            }
        }

        total += subtotal;
    }
    println!("Total: {}", total);
}
