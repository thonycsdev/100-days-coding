use std::env;
use std::fs::File;
fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();

    let file_path: &String = &args[0];
    let file_name: &String = &args[1];

    let mut file = File::open(file_name)?;
    let mut buffer = Vec::new();

    Ok(());
}
