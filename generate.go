package bzdacs

//go:generate swag init -g ./cmd/backend/main.go -o generated
//go:generate swagger generate client -f ./generated/swagger.yaml -t ./generated
//go:generate npm run --prefix ./ui gen-swag
