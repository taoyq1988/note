cvt: tools/cvt
	cd tools/cvt && go build -o cvt
	cd tools/cvt && mv cvt ${GOPATH}/bin

.PHONY: cvt