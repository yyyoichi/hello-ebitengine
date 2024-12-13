dev: build
	go run cmd/server/main.go

server:
	go run cmd/server/main.go

# Ubuntu想定

build:
	env GOOS=js GOARCH=wasm go build -o yourgame.wasm cmd/game/main.go

init:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js .
	echo "<!DOCTYPE html>" > index.html
	echo "<script src=\"wasm_exec.js\"></script>" >> index.html
	echo "<script>" >> index.html
	echo "const go = new Go();" >> index.html
	echo "WebAssembly.instantiateStreaming(fetch(\"yourgame.wasm\"), go.importObject).then(result => {" >> index.html
	echo "    go.run(result.instance);" >> index.html
	echo "});" >> index.html
	echo "</script>" >> index.html
