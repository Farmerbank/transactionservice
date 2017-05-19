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


### Start MongoDB on the server

```
$ docker run --name farmerbank -p 127.0.0.1:27017:27017 -v /srv/farmerbank-app/mongo-data:/data/db -d mongo
```