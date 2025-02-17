use once_cell::sync::Lazy;
use std::{collections::HashMap, sync::Mutex};

use crate::cmds::{handle_cd, handle_echo, handle_exit, handle_pwd, handle_type, CmdFn};
pub static CMDS: Lazy<Mutex<HashMap<String, CmdFn>>> = Lazy::new(|| {
    let mut m = HashMap::new();
    m.insert("exit".to_string(), handle_exit as CmdFn);
    m.insert("echo".to_string(), handle_echo as CmdFn);
    m.insert("type".to_string(), handle_type as CmdFn);
    m.insert("pwd".to_string(), handle_pwd as CmdFn);
    m.insert("cd".to_string(), handle_cd as CmdFn);
    Mutex::new(m)
});
