out: vendor main.go
	go build \
		-gcflags all=-trimpath=/home/noah/documents/go \
		-asmflags all=-trimpath=/home/noah/documents/go \
		-mod=vendor \
		-o out \
	.

vendor: go.mod
	go mod vendor
