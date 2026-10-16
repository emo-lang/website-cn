package home

import _ "embed"

// The runnable landing-page samples are real Emo projects under samples/,
// each with its own package.emo. Keeping the source in .emo files means the
// code shown on the page and the WebAssembly the browser runs come from one
// source; scripts/build-demo-wasm.sh compiles each project into
// app/assets/public/demo/<name>.wasm.
//
//go:embed samples/hero/main.emo
var heroSample string

//go:embed samples/pipeline/main.emo
var concurrencySample string

//go:embed samples/greeter/main.emo
var interfaceSample string
