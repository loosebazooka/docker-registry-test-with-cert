docker stop tls_registry && docker rm tls_registry
docker run -d -p 5000:5000 \
  --name tls_registry \
  -v $(pwd)/out/certs:/certs \
  -e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/localhost.crt \
  -e REGISTRY_HTTP_TLS_KEY=/certs/localhost.key \
  registry:2
