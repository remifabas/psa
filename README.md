# PSA CLI for CMM project

CLI used for interact with AWS infrastructure of PSA-CMM Project
This CLI is powered by [cobra](https://github.com/spf13/cobra)  

## How to Install this CLI
You need golang at least go1.13.5.  
To find your go version...  
```
go version
```
[GO INSTALL](https://golang.org/doc/install)  

Don't forget to add to your path your gopath (in my zshrc or bashrc)
for me :
```
export PATH=$PATH:/home/remi/go/bin
```

Clone this repository then run in psa folder
```
go install
# THEN
psa -h
```
You have now access to psa cli.

list of [commands](commands.md)

## AWS Credentials
In order to perform some basic commands you need your aws credentials set up  
Credentials from the shared credentials file. (~/.aws/credentials).
If not, ask them to an AWS admin, then run aws configure

## Develop
First install cobra cli on you system [cobra](https://github.com/spf13/cobra) 
```
# First run some test
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out

# Add command : 
cobra add youcommand
```