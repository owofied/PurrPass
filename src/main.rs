use axum::Router;
use log::info;

mod config;

#[tokio::main]
async fn main() {
    env_logger::init();

    let app = Router::new();

    info!("Starting server");
    let listener = tokio::net::TcpListener::bind("127.0.0.1:3000")
        .await
        .unwrap();
    axum::serve(listener, app).await.unwrap()
}
