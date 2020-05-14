# go-fil-commp

See also https://github.com/rvagg/rust-fil-commp-generate & https://github.com/rvagg/js-fil-utils for similar utilities in Rust and JavaScript.

This repo is older than both of the above repositories and likely doesn't match the current best practice way to generate CommP with Go.

* Build with `make`
* Create a commP for a file with `commp <path to file>`

Borrowed liberally from [Lotus](https://lotu.sh/), specifically:

 * github.com/filecoin-project/lotus/blob/f019b80a/Makefile
 * github.com/filecoin-project/lotus/blob/f019b80a/chain/deals/client_utils.go#L64-L71
 * github.com/filecoin-project/lotus/blob/f019b80a/lib/padreader/padreader.go
