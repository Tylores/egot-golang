# sep-golang
Smart Energy Profile using Go

## Setup

### SSL
I found a super handy tool for setting yourself as a CA and generating tls certificates for clients and servers.

* https://smallstep.com/docs/step-cli/reference/

After you have step and step-ca installed simply issue a ca, server, and client sertificate using the very handy tutorial and you are up and ready for tls.

* https://smallstep.com/hello-mtls/doc/combined/go/go

If you follow the basic example settings for the CA setup you need to modify the default certificate duration to be greater than 24 hours. Use the following code, but verify that the provisioner identity is correct for your installation. 

* https://smallstep.com/docs/step-ca/provisioners/#remote-provisioner-management

```shell
step ca provisioner update you@smallstep.com \
   --x509-min-dur=24h \
   --x509-max-dur=8760h \
   --x509-default-dur=8760h
```

### Golang
Go setup is very simple.

 ```shell
go mod init sep-golang
```

## Run

Run the server in the background and then test a client.

```shell
go run server/server.go &
go run client/client.go
```
