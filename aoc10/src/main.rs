use aoc10::utils;
use std::collections::HashSet;

fn main() {
    let test = utils::is_test().unwrap();
    let input = utils::parse_input(test);

    let closing_mapper = utils::build_closing_bracket_mapper();
    let score_mapper = utils::build_score_mapper();
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
                    // Maintain the character if it has already been set, otherwise set
                    first_illegal_character = match first_illegal_character {
                        Some(value) => Some(value),
                        None => Some(bracket),
                    };

                    // Only accrue score if it matches the first illegal char
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
