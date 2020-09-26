# nand2tetris HackVM implemented in Go

# Install TinyGO

https://tinygo.org/

## Compile to WebAssembly

```shell
$ tinygo build -o web/main.wasm -target wasm cmd/main.go
```

## Start the server

```shell
$ cd web
$ go run main.go
```

## Voila

![Pong](screenshot.png "Pong running on HackVM")

## References

https://www.nand2tetris.org/
