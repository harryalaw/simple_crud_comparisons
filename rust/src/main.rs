//! Run with
//!
//! ```not_rust
//! cargo run -p example-hello-world
//! ```

use axum::{
    routing::{get},
    Router,
};
use std::net::SocketAddr;

#[tokio::main]
async fn main() {
    // run it
    let app = Router::new()
        // `GET /` goes to `root`
        .route("/", get(root));

    // run our app with hyper
    // `axum::Server` is a re-export of `hyper::Server`
    let addr = SocketAddr::from(([127, 0, 0, 1], 4321));
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn root() -> &'static str {
    "Hello, World!"
}
