# CheersCLI

Randomly generates greeting messages.

Also provides random emojis.

```bash
$ ./cheers
hola!

$ ./cheers emoji
💯
```


# TODO

- cross-compiling
    ```bash
    # For Intel Mac
    $ GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin_amd64/cheers main.go
    ```
  - apple,   amd64 (done)
  - apple,   arm64
  - linux,   amd64
  - windows, amd64
