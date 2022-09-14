# Stuff found here

## Go commands and Stuff

- build everything - local modules and `main`
  - `go build ./ ...`
- list modules
  - `go list -m`
- update modules
  - `go get -u`
- tidy modules
  - `go mod tidy`

## Command history

37215  rm ../gorm.db
37260  pj go
37262  cd golang-gin-realworld-example-app/
37265  go build
37266  go test

~~build should have been `go build ./ ...` (space between ./ and ...)~~
~~test should have neen `go test ./ ...` (space between ./ and ...)~~

37270  go mod graph
37272  cat go.sum
37273  gd go.sum
37284  go list -m
37285  go list -m -u all
37287  ga go.sum

37298  go get -u
37300  go mod tidy

37339  go mod tidy
37340  go mod
37342  go clean
37343  go build

Reverted some changes to find a silly bug, now redoing

37354  go get -u
37355  go mod tidy
37357  go build

37441  pj go
37442  cd golang-gin-realworld-example-app/
37449  go build

37459  go get "github.com/gin-contrib/logger"
37460  go get "github.com/gin-contrib/requestid"
37461  go get "github.com/rs/zerolog"
37462  go get "github.com/rs/zerolog/log"
37463  go build

How this got here :)

37518  history | grep go | grep ^37 | grep go >> README-SHAZ.md

## Broken stuff that I got wrong - no extra space after ./ in build and test

37477  go build ./ ...
37480  go mod tidy
37481  go build ./ ...

37492  go build ./ ...
37493  go test ./ ...

Added packages as guided by `correct` build command output

37494  go get github.com/kr/text/mc@v0.2.0
37495  go get github.com/rogpeppe/go-internal/testscript@v1.8.0
37496  go get github.com/rs/zerolog/hlog@v1.28.0
37497  go get github.com/rs/zerolog/journald@v1.28.0
37498  go get github.com/rs/zerolog/pkgerrors@v1.28.0
37499  go get github.com/stretchr/testify/mock@v1.8.0
37500  go get go.opentelemetry.io/otel@v1.10.0
37501  go get go.opentelemetry.io/otel/internal/global@v1.10.0
37502  go get golang.org/x/crypto/ssh/terminal@v0.0.0-20220829220503-c86fa9a7ed90
37503  go get golang.org/x/text/cmd/gotext@v0.3.7
37504  go get golang.org/x/text/message/pipeline@v0.3.7
37505  go get golang.org/x/text/cmd/gotext@v0.3.7
37506  go get golang.org/x/text/message/pipeline@v0.3.7
37507  go build ./ ...

This is not present - need to get actual version
37508  go get github.com/coreos/go-systemd/v22/dbus@v22.3.3-0.20220203105225-a9a7ef127534
37509  go build ./ ...

Tried to downgrade above to v22.3.2 which is in github, but that failed

??? Maybe we don't need to `go build ./ ...` after all the error only shows when that command is run
What seems to be being built is - [rs/zerolog](https://github.com/rs/zerolog/blob/master/go.sum)

## Correction

Build command is `go build ./...` - no space after `./`
Test command is also `go test ./...`
