module github.com/rvagg/go-fil-commp

go 1.13

require (
	github.com/aws/aws-lambda-go v1.13.3
	github.com/aws/aws-sdk-go v1.28.3
	github.com/dustin/go-humanize v1.0.0
	github.com/filecoin-project/filecoin-ffi v0.0.0-20191219131535-bb699517a590
	github.com/filecoin-project/go-fil-markets v0.0.0-20200118022459-68964015978c
	github.com/filecoin-project/go-sectorbuilder v0.0.1
	github.com/ipfs/go-cid v0.0.5 // indirect
)

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi
replace github.com/ipfs/go-cid => ../../ipfs/go-cid
