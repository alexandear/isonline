# Is Online

This repo contains a program that checks Internet connectivity of a provider.

## Supported Providers

* [DataGroup](https://datagroup.ua/)

## How to Run

* Download and install [Go](https://go.dev/dl/).
* Run with one command:
```shell
go run cmd/main.go -datagroup_login "<YOUR_DATAGROUP_LOGIN>" -datagroup_password "<YOUR_DATAGROUP_PASSWORD>"
```
or build and run:
```shell
go build -o isonline cmd/main.go
./isonline -datagroup_login "<YOUR_DATAGROUP_LOGIN>" -datagroup_password "<YOUR_DATAGROUP_PASSWORD>"
```

It is possible to use environment variables instead of flags:
```shell
DATAGROUP_LOGIN="<YOUR_DATAGROUP_LOGIN>" DATAGROUP_PASSWORD="<YOUR_DATAGROUP_PASSWORD>" ./isonline
```
 