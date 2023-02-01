# GoWASM
Explore some frontend work

Build WASM: 
GOOS=js GOARCH=wasm go build -o assets/main.wasm cmd/wasm/main.go

Run server:
go run cmd/server/main.go




to fix this red herring: could not import syscall/js (no required module provides package "syscall/js")
add this to your vscode setting:

    "gopls": {
        "build.buildFlags": ["-tags=tools"],
        "biuld.env": {
            "GOOS": "js",
            "GOARCH": "wasm"
        }
    },
