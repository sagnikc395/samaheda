use std::env;
use std::io::{self, Write};
use std::process::{self, Command};

use crate::consts::CMDS;

//defining the type for command functions
pub type CmdFn = fn(Vec<String>) -> io::Result<()>;

pub fn handle_exit(args: Vec<String>) -> io::Result<()> {
    let exit_code = if args.is_empty() {
        0
    } else {
        args[0].parse().unwrap_or(0)
    };

    process::exit(exit_code);
}

pub fn handle_echo(args: Vec<String>) -> io::Result<()> {
    if args.is_empty() {
        println!();
        return Ok(());
    }

    let output = args.join(" ");
    println!("{}", output);
    Ok(())
}

pub fn locate_cmd(cmd: &str) -> Option<String> {
    if let Some(path) = env::var_os("PATH") {
        for dir in env::split_paths(&path) {
            let full_path = dir.join(cmd);
            if full_path.is_file() {
                return full_path.to_str().map(String::from);
            }
        }
    }
    None
}

pub fn handle_type(args: Vec<String>) -> io::Result<()> {
    if args.is_empty() {
        return Ok(());
    }

    let cmd = &args[0];
    //lock mutex before accessing it
    if let Ok(commands) = CMDS.lock() {
        if commands.contains_key(cmd) {
            eprintln!("{} is a shell builtin", cmd);
            return Ok(());
        }
    }

    if let Some(path) = locate_cmd(cmd) {
        println!("{} is {}", cmd, path);
        return Ok(());
    }

    eprintln!("{}: not found", cmd);
    Ok(())
}

pub fn handle_pwd(_args: Vec<String>) -> io::Result<()> {
    let current_dir = env::current_dir()?;
    println!("{}", current_dir.display());
    Ok(())
}

pub fn handle_cd(args: Vec<String>) -> io::Result<()> {
    if args.is_empty() {
        return Ok(());
    }

    let mut dir = args[0].clone();
    if dir == "~" {
        dir = env::var("HOME").unwrap_or_else(|_| String::from("~"));
    }

    if let Err(_) = env::set_current_dir(&dir) {
        eprintln!("cd: {}: No such file or directory", dir);
    }
    Ok(())
}

pub fn get_command(cmd: &str) -> Option<CmdFn> {
    if let Ok(commands) = CMDS.lock() {
        commands.get(cmd).copied()
    } else {
        None
    }
}

pub fn execute_external_command(cmd: &str, args: Vec<String>) -> io::Result<()> {
    match Command::new(cmd).args(args).output() {
        Ok(output) => {
            io::stdout().write_all(&output.stdout)?;
            io::stderr().write_all(&output.stderr)?;
            Ok(())
        }
        Err(err) => {
            eprintln!("{}", err);
            Ok(())
        }
    }
}
