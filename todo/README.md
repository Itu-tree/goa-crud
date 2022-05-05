#  コンテナ立ち上げ後の作業
go mod init todo
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
go install  golang.org/x/lint/golint@latest
go install golang.org/x/tools/gopls@latest
go get -u github.com/go-sql-driver/mysql
go get -u  github.com/jmoiron/sqlx 
go install goa.design/goa/v3/cmd/goa@v3 
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest