module github.com/ArchieBVox/m3u8-downloader/dl

go 1.23.4

require (
	github.com/ArchieBVox/m3u8-downloader/parse v0.0.0-20241231005606-d5ca78353c21
	github.com/ArchieBVox/m3u8-downloader/tool v0.0.0-20241231003836-2f48bae42ae0
)

<<<<<<< HEAD
replace github.com/ArchieBVox/m3u8-downloader/tool => ../tool
=======
require (
	github.com/http-live-streaming/m3u8-downloader/tool v0.0.0-20210218104036-d0d4d45ee0bc // indirect
	golang.org/x/net v0.33.0 // indirect
)

replace github.com/http-live-streaming/m3u8-downloader/tool => ../tool
>>>>>>> 1c4f870 (Use local copy of github.com/http-live-streaming/m3u8-downloader/tool)
