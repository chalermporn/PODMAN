use warp::{Filter, reply};

#[tokio::main]
async fn main() {
    let hello = warp::path::end()
        .map(|| reply::html("Hello, World from Warp!"));

    warp::serve(hello)
        .run(([0, 0, 0, 0], 3000))
        .await;
}