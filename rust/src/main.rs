//! Run with
//!
//! ```not_rust
//! cargo run -p example-hello-world
//! ```

use axum::{response::Json, routing::get, Router};
use serde_json::{json, Value};
use std::net::SocketAddr;

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/", get(root))
        .route("/json", get(json));

    let addr = SocketAddr::from(([127, 0, 0, 1], 4321));
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn root() -> &'static str {
    "Hello, World!"
}

async fn json() -> Json<Value> {
    Json(
        json!({ "id": 1234, "message": "I'm a JSON value", "array": ["here's", "some", "values", "in", "an", "array", "to", "make", "the", "object", "bigger"]}),
    )
}
