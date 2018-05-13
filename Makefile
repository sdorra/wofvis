VERSION=0.1.0

.PHONY=build
build: wofvis

${GOPATH}/bin/go-bindata:
	@go get github.com/jteeuwen/go-bindata/...

${GOPATH}/bin/go-bindata-assetfs: ${GOPATH}/bin/go-bindata
	@go get github.com/elazarl/go-bindata-assetfs/...

pkg/bindata.go: ${GOPATH}/bin/go-bindata-assetfs
	@cd pkg; go generate

client/node_modules:
	@cd client; yarn install

client/build: client/node_modules
	@cd client; yarn build

.PHONY=assets
assets: pkg/bindata.go client/build

${GOPATH}/bin/dep:
	@go get -u github.com/golang/dep/cmd/dep

vendor: ${GOPATH}/bin/dep
	@dep ensure

${GOPATH}/bin/gox:
	@go get github.com/mitchellh/gox

wofvis: assets ${GOPATH}/bin/gox
	gox -arch amd64 -os "darwin linux windows" -ldflags "-X main.Version=${VERSION}" -output dist/wofvis_{{.OS}}_{{.Arch}}

.PHONY=clean
clean:
	@rm -f pkg/bindata.go
	@rm -rf client/build
	@rm -f wofvis
	@rm -rf dist

.PHONY=distclean
distclean: clean
	@rm -rf client/node_modules
	@rm -rf vendor


