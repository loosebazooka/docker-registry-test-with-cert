# docker-registry-test-with-cert
A guide and some scripts to get local go applications to use a self signed cert for testing

### make your certificates
```
$ ./make_certs.sh
```
`make_certs.sh` will create sample test cert and key in `./out/certs`
```
out
└── certs
    ├── localhost.crt
    └── localhost.key
```

### start a registry wth those certs
```
$ ./start_registry.sh
```

Once the registry is up and running, you can try communicating with it

### try with `curl`
```
$ curl -v https://localhost:5000/v2/

ERROR because no valid cert

```
Now use our custom cert
```
$ curl -v --cacert ./out/certs/localhost.crt https://localhost:5000/v2/

SUCCESS
```

### try with golang application
Now to use this cert with an existing go application, you must set the env `SSL_CERT_FILE` before invocation. In this example we use [crane](https://github.com/google/go-containerregistry/blob/main/cmd/crane/README.md)
```
$ crane -v catalog localhost:5000

ERROR because no valid cert
```
Now use our custom cert
```
$ SSL_CERT_FILE=./out/certs/localhost.crt crane -v catalog localhost:5000

SUCCESS
```

### try setting programmatically in golang code

TODO
