fn parse_input(test: bool) -> Vec<Vec<u16>> {
    let file_contents = match test {
        true => include_str!("../inputs/test.txt").lines(),
        false => include_str!("../inputs/puzzle.txt").lines(),
    };
    let mut points: Vec<Vec<u16>> = Vec::new();

    for line in file_contents {
        let line_points = line
            .split("")
            .filter(|&height| !height.is_empty())
            .map(|x| x.trim().parse::<u16>().unwrap())
            .collect();

        let _ = &points.push(line_points);
    }

    points
}
fn main() {
    let input = parse_input(true);
    println!("{:?}", input);
}
