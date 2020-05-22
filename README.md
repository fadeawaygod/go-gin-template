# go-iris-template

Setup:
1. install go 1.14 or later
2. copy .env.template to .env and configurate
3. go build

Go mod
1. go env -w GO111MODULE=on(1.12 or earlier)
2. go mod init go-iris-template(if not exist)

test

mock
- generate mock data file:
  - mockgen -source=ping_handler.go -destination=ping_handler_mock.go