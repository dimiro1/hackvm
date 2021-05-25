# nand2tetris HackVM implemented in Go

# Install TinyGO

https://tinygo.org/

## Compile to WebAssembly

```shell
$ tinygo build -o demo/main.wasm -target wasm hackvm/cmd
```

## Start the server

```shell
$ cd demo
$ go run main.go
```

## Voila

![Pong](screenshot.png "Pong running on HackVM")

## Headless mode (Standard Go)

```shell
$ go run hackvm/cmd
```

## References

https://www.nand2tetris.org/
