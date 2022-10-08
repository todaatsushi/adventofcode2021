use aoc13::data::Map;
use aoc13::parse;

fn main() {
    let (points, folds) = parse::read_input();
    let mut map = Map::new(points);

    map.fold(&folds[0]);
    map.display(true, true);
}
