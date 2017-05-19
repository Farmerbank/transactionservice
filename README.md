# Farmerbank

This project provides an integration with Alexa.


### Build and Run

Install dependencies using [Glide Package Management for Go](https://glide.sh/)

```
$ glide update
$ glide install
```

Release the application for all platforms
```
$ make release
```

```
$ scp release/transactionservice-linux-amd64 ubuntu@farmerbank.nl:/srv/farmerbank-app/transactionservice
```

Optionally you can specify the server address using `$ ./transactionservice`
