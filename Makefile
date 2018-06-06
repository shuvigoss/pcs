all:
	$(MAKE) deps
	$(MAKE) pcs

deps:
	go get github.com/sparrc/gdm
	gdm vendor

pcs:
	GOOS=linux GOARCH=amd64 go build  ./cmd/pcs

go-install:
	go install  ./cmd/pcs

clean:
	rm -f pcs
	rm -rf build

package:
	$(MAKE) clean
	mkdir build
	$(MAKE) pcs
	mv pcs ./build/pcs
	cp etc/pcs.json build/
	cd build && tar -cvzf pcs.tar.gz ./*


.PHONY: deps pcs go-install clean
