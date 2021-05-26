package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

const (
	certPath = "../out/certs/localhost.crt"
)

func main() {
	reg, err := name.NewRegistry("localhost:5000")
	if err != nil {
		os.Exit(1)
	}

	transport, err := getTransport(certPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	catalog, err := remote.Catalog(context.Background(),
		reg,
		remote.WithAuthFromKeychain(authn.DefaultKeychain),
		remote.WithTransport(transport))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(catalog)
}

func getTransport(path string) (*http.Transport, error) {
	pool, err := x509.SystemCertPool()
	if err != nil {
		pool = x509.NewCertPool()
	}

	if crt, err := ioutil.ReadFile(path); err != nil {
		return nil, err
	} else if ok := pool.AppendCertsFromPEM(crt); !ok {
		return nil, errors.New("failed to append homemade cert to pool")
	}

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    pool,
	}

	return transport, nil
}
