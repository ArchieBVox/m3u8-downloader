module github.com/ArchieBVox/m3u8-downloader/dl

go 1.23.4

require (
	github.com/ArchieBVox/m3u8-downloader/parse v0.0.0-20241231005606-d5ca78353c21
	github.com/ArchieBVox/m3u8-downloader/tool v0.0.0-20241231003836-2f48bae42ae0
)

require golang.org/x/net v0.33.0 // indirect

replace github.com/ArchieBVox/m3u8-downloader/tool => ../tool
