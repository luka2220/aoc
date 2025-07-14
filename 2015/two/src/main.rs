use std::fs::File;
use std::io::BufReader;
use std::io::{self, BufRead};

// surface area of a box: 2*l*w + 2*w*h + 2*h*l
// extra slack: area of smallest side
// i.e  2x3x4
// 2*6 + 2*12 + 2*8 = 52
// 52 + 6 = 58

fn main() -> io::Result<()> {
    let f = File::open("input.txt")?;
    let reader = BufReader::new(f);

    let mut total = 0;

    for line in reader.lines() {
        let line = line?;
        let dims: Vec<&str> = line.split('x').collect();

        let l: i32 = dims[0].parse().expect("Str val not an int");
        let w: i32 = dims[1].parse().expect("Str val not an int");
        let h: i32 = dims[2].parse().expect("Str val not an int");

        let vals: [i32; 3] = [l * w, w * h, h * l];
        let slack = vals.iter().min().unwrap();

        total += *slack + (2 * vals[0] + 2 * vals[1] + 2 * vals[2]);
    }

    println!("Total amount of surface area for the presents: {total}");

    Ok(())
}
