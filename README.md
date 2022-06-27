
CGO_CFLAGS="-std=c99 -w" go run .

Current working:
CGO_CFLAGS="-std=c11 -w" go build -a . && ./retro-image-tool