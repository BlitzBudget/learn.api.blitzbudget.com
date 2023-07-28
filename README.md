# golang-lambda-functions
Template Repository for Golang Lambda Functions


## Test Golang AWS Lambdas

Testing out things isn't so complicated.

What we need to do is fire things up.

### Execute Golang Tests

Inorder to Execute golang tests, execute the command below.

`go test -v ./...`

### To Install dependencies

`go get -v all`

### Build project

Set the processor and operating system with powershell

`$Env:GOOS = "linux"; $Env:GOARCH = "amd64";`

Build the package in build/main folder 

`go build -o build/main cmd/main.go`

### We need to Zip binary to be able to upload to AWS lambda (Linux)

`zip -jrm build/main.zip build/main`


# Auto Deploy to Lambda Function

The automatic integration to lambda function is enabled and as such all the deployments are automatic. They are deployed directly to the production.


# Enable Precommit hooks

In order to enable the pre commit hooks execute the following command in Linux OS. (Windows automatically launches it)

`chmod +x pre-commit`