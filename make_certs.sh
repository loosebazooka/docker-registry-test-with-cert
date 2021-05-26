mkdir -p out/certs
openssl req  -newkey rsa:4096 -nodes -sha256 -keyout out/certs/localhost.key -x509 -days 365 -out out/certs/localhost.crt \
    -subj "/C=US/ST=NY/L=New York /O=Test/OU=Test/CN=localhost/emailAddress=test@example.com" \
    -addext "subjectAltName = DNS:localhost"
