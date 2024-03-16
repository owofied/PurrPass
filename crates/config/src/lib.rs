use serde::Serialize;

#[derive(Serialize, Debug)]
struct Config {}

#[derive(Serialize, Debug)]
struct ServerConfig {
    host: String,
    port: u16,
}

impl Config {}
