use axum::Router;
use log::{error, info};

#[tokio::main]
async fn main() {
    env_logger::init();

    let app = Router::new();

    info!("Starting server");
    let listener = tokio::net::TcpListener::bind("127.0.0.1:3000").await;
    match &listener {
        Ok(_) => {
            info!("Created listener on 127.0.0.1:3000");
        }
        Err(err) => {
            error!("Failed to bind to 127.0.0.1:3000. Error: {:?}", err);
        }
    }

    axum::serve(listener.unwrap(), app).await.unwrap()
}
