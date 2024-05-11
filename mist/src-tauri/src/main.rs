// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

use hello_world::greeter_client::GreeterClient;
use hello_world::{HelloReply, HelloRequest};

pub mod hello_world {
    tonic::include_proto!("helloworld");
}

use tauri::Manager;
use tokio;

// Learn more about Tauri commands at https://tauri.app/v1/guides/features/command
#[tauri::command]
fn greet(name: &str) -> String {
    format!("Hello, {}! You've been  from Rust!", name)
}

fn main() {
    tauri::Builder::default()
        .setup(|app| {
            let app_handle = app.handle();

            tauri::async_runtime::spawn(async move {
                loop {
                    tokio::time::sleep(tokio::time::Duration::from_millis(1000)).await;
                    let mut client = GreeterClient::connect("http://[::1]:50051").await.unwrap();

                    let request = tonic::Request::new(HelloRequest {
                        name: "Tonic".into(),
                    });

                    let response = client.say_hello(request).await.unwrap();

                    app_handle
                        .emit_all("notification", response.into_inner().message)
                        .unwrap();
                }
            });

            Ok(())
        })
        .invoke_handler(tauri::generate_handler![greet])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
