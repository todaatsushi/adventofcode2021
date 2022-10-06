use std::collections::{HashMap, HashSet, VecDeque};

pub fn get_num_paths(caves: HashMap<&str, Vec<&str>>, allow_doubles: bool) -> u32 {
    let mut num_paths: u32 = 0;
    let mut queue: VecDeque<(&str, HashSet<&str>, bool)> = VecDeque::new();

    let mut init_set = HashSet::new();
    init_set.insert("start");
    queue.push_back(("start", init_set, false));

    while queue.is_empty() == false {
        let (current, seen, visited) = queue.pop_front().unwrap();
        if current == "end" {
            num_paths += 1;
            continue;
        }

        let next_caves = caves.get(&current.clone()).unwrap();
        for cave in next_caves {
            let is_start_or_end = &"start" == cave || &"end" == cave;
            let seen_cave = seen.contains(cave);
            let mut now_seen = seen.clone();

            if seen_cave == false {
                if cave.chars().all(|c| c.is_lowercase()) == true {
                    now_seen.insert(cave);
                }
                queue.push_front((cave, now_seen, visited));
            } else if seen_cave && visited == false && is_start_or_end == false && allow_doubles {
                queue.push_front((cave, now_seen, true));
            }
        }
    }

    num_paths
}
