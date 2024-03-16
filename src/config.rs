struct Config {
    pub server: Option<ServerConfig>
}

struct ServerConfig {
    host: String,
    port: i16,
}

