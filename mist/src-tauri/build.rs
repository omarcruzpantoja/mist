fn main() {
    tonic_build::compile_protos("protos/helloworld.proto");
    tauri_build::build()
}
