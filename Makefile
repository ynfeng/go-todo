BINARY="todo"

build:
	@go build -o ${BINARY} -tags=jsoniter

test:
	@go test -cpu=1,2,4 -v -tags integration ./...

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: test build
