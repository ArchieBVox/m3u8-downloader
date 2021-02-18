# m3u8 downloader ![Go](https://github.com/http-live-streaming/m3u8-downloader/workflows/Go/badge.svg?branch=master)
A M3U8 downloader (CLI) written in Golang to download HTTP Live Streaming videos, forked from [oopsguy/m3u8](https://github.com/oopsguy/m3u8) with continuous updates.

You only need to specify the flags(`u`, `o`, `c`) to run, downloader will automatically download all TS files and consolidate them into a single TS file.

## Features

- Download and parse M3U8（VOD）
- Retry on download TS failure
- Parse Master playlist
- Decrypt TS
- Merge TS

## Usage

### source

```bash
go run main.go -u=http://example.com/index.m3u8 -o=/data/example
```

### binary:

Linux & MacOS

```
./m3u8 -u=http://example.com/index.m3u8 -o=/data/example
```

Windows PowerShell

```
.\m3u8.exe -u="http://example.com/index.m3u8" -o="D:\data\example"
```

## Download

[Binary packages](https://github.com/http-live-streaming/m3u8-downloader/releases)

## Screenshots

![Demo](./screenshots/demo.gif)

## References

- [TS科普 2 包头](https://blog.csdn.net/cabbage2008/article/details/49281729)
- [HTTP Live Streaming draft-pantos-http-live-streaming-23](https://tools.ietf.org/html/draft-pantos-http-live-streaming-23#section-4.3.4.2)
- [MPEG transport stream - Wikipedia](https://en.wikipedia.org/wiki/MPEG_transport_stream)


## License

[MIT License](./LICENSE)
