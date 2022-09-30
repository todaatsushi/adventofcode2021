use crate::utils;
use std::collections::{HashMap, HashSet};

fn build_score_mapper() -> HashMap<&'static str, u128> {
    let mut map = HashMap::new();

    map.insert(")", 1);
    map.insert("]", 2);
    map.insert("}", 3);
    map.insert(">", 4);
    map
}

pub fn solve_part_2(input: &Vec<Vec<&'static str>>) {
    let closing_mapper = utils::build_closing_bracket_mapper();
    let opening_brackets = HashSet::from(["{", "(", "<", "["]);

    let mut expected: &str;

    for line in input {
        let mut bracket_stack: Vec<&str> = Vec::new();
        let mut incomplete = true;

        for bracket in line {
            if !incomplete {
                break;
            }

            if opening_brackets.contains(bracket) {
                bracket_stack.push(closing_mapper.get(bracket).unwrap());
            } else {
                expected = match bracket_stack.pop() {
                    Some(b) => b,
                    None => panic!("bracket_stack empty"),
                };
                if expected != *bracket {
                    incomplete = false;
                }
            }
        }

        if incomplete {}
    }
}
