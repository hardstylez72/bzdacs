package bzdacs

//go:generate swag init -g ./cmd/backend/main.go
//go:generate swagger generate client -f ./docs/swagger.yaml
//go:generate npm run --prefix ./ui gen-swag
