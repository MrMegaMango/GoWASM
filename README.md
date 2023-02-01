# GoWASM
Explore some frontend work

to fix this red herring: could not import syscall/js (no required module provides package "syscall/js")
add this to your vscode setting:

    "gopls": {
        "build.buildFlags": ["-tags=tools"],
        "biuld.env": {
            "GOOS": "js",
            "GOARCH": "wasm"
        }
    },
