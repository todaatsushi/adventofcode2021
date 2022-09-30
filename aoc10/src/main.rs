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
    println!("{:?}", input);
}
