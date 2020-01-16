module github.com/rvagg/go-fil-commp

go 1.13

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/filecoin-project/filecoin-ffi v0.0.0-20191219131535-bb699517a590
	github.com/filecoin-project/go-sectorbuilder v0.0.1
)

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi
