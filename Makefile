VERSION=0.1.0
GITHUB_USER=sdorra
GITHUB_REPO=wofvis

.PHONY=build
build: dist

${GOPATH}/bin/go-bindata:
	@go get github.com/jteeuwen/go-bindata/...

${GOPATH}/bin/go-bindata-assetfs: ${GOPATH}/bin/go-bindata
	@go get github.com/elazarl/go-bindata-assetfs/...

pkg/bindata.go: ${GOPATH}/bin/go-bindata-assetfs client/build
	@cd pkg; go generate

client/node_modules:
	@cd client; yarn install

client/build: client/node_modules
	@cd client; yarn build

.PHONY=assets
assets: pkg/bindata.go

${GOPATH}/bin/dep:
	@go get -u github.com/golang/dep/cmd/dep

vendor: ${GOPATH}/bin/dep
	@dep ensure

${GOPATH}/bin/gox:
	@go get github.com/mitchellh/gox

dist: assets ${GOPATH}/bin/gox
	@gox -arch amd64 -os "darwin linux windows" -ldflags "-X main.Version=${VERSION}" -output dist/wofvis-${VERSION}_{{.OS}}_{{.Arch}}
	@cd dist; shasum -a 256 * > wofvis-${VERSION}.sha256sums
	@cd dist; find . -type f -not -name "*.sha256sums" -exec gpg -armor --detach-sig {} \;

${GOPATH}/bin/github-release:
	@go get github.com/aktau/github-release

release: ${GOPATH}/bin/github-release dist
	@git tag -s -m "release v${VERSION}" v${VERSION}
	@git push origin master --tags
	@github-release release \
  		--user ${GITHUB_USER} \
  		--repo ${GITHUB_REPO} \
  		--tag v${VERSION} \
  		--name v${VERSION} \
  		--description "release version ${VERSION}"
	@cd dist; ls -1 | xargs -n1 -I{} -- \
	 									github-release upload \
																	--user ${GITHUB_USER} \
																	--repo ${GITHUB_REPO} \
																	--tag v${VERSION} \
																	--name {} \
																	--file {}



.PHONY=clean
clean:
	@rm -f pkg/bindata.go
	@rm -rf client/build
	@rm -rf dist

.PHONY=distclean
distclean: clean
	@rm -rf client/node_modules
	@rm -rf vendor


