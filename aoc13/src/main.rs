use aoc13::data::Map;
use aoc13::parse;

fn main() {
    let (points, folds) = parse::read_input();
    let mut map = Map::new(points);

    for fold in folds {
        map.fold(&fold);
    }
    map.display(true, false);
}
