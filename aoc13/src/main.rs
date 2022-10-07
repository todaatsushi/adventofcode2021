use aoc13::data::Map;
use aoc13::parse;

fn main() {
    let (points, _folds) = parse::read_input();
    let map = Map::new(points);
    map.display();
}
