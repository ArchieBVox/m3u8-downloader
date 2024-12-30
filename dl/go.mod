module github.com/ArchieBVox/m3u8-downloader/dl

go 1.23.4

require (
	github.com/ArchieBVox/m3u8-downloader/parse v0.0.0-20241231005606-d5ca78353c21
	github.com/ArchieBVox/m3u8-downloader/tool v0.0.0-20241231003836-2f48bae42ae0
)

replace github.com/http-live-streaming/m3u8-downloader/tool => ../tool
