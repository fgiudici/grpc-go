module github.com/fgiudici/grpc-go/examples

go 1.11

require (
	cloud.google.com/go v0.63.0 // indirect
	github.com/golang/protobuf v1.4.2
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98
	github.com/fgiudici/grpc-go v1.31.0
)

replace github.com/fgiudici/grpc-go => ../
