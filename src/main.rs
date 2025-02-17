use std::io::{self, Write};

use samaheda::{
    cmds::{execute_external_command, get_command, locate_cmd},
    consts::CMDS,
};

fn parse_input(input: &str) -> Vec<String> {
    let mut tokens = Vec::new();
    let mut s = input.trim();

    while !s.is_empty() {
        // Check for quoted strings
        if let Some(first_char) = s.chars().next() {
            if first_char == '"' || first_char == '\'' {
                if let Some(end) = s[1..].find(first_char) {
                    // Add the quoted string without quotes
                    tokens.push(s[1..=end].to_string());
                    s = &s[end + 2..].trim();
                    continue;
                }
            }
        }

        // Handle regular words
        if let Some(end) = s.find(char::is_whitespace) {
            let word = &s[..end];
            if !word.is_empty() {
                tokens.push(word.to_string());
            }
            s = &s[end..].trim();
        } else {
            if !s.is_empty() {
                tokens.push(s.to_string());
            }
            break;
        }
    }

    tokens
}

fn main() -> io::Result<()> {
    let stdin = io::stdin();
    let mut stdout = io::stdout();

    loop {
        // Print prompt
        write!(stdout, "$ ")?;
        stdout.flush()?;

        // Read input
        let mut input = String::new();
        stdin.read_line(&mut input)?;

        // Parse input
        let tokens = parse_input(&input);
        if tokens.is_empty() {
            continue;
        }

        let cmd = tokens[0].to_lowercase();
        let args = tokens[1..].to_vec();

        // Execute command
        if let Some(cmd_fn) = get_command(&cmd) {
            if let Err(err) = cmd_fn(args) {
                eprintln!("{}", err);
            }
        } else if let Some(path) = locate_cmd(&cmd) {
            execute_external_command(&path, args)?;
        } else {
            println!("{}: command not found", cmd);
        }
    }
}
