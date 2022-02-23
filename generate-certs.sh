#!/bin/bash

openssl="/usr/local/opt/openssl@3/bin/openssl"

# MacOS verified only:
ips=$(ifconfig | grep -Eo 'inet [0-9.]*' | sed 's/inet /IP:/g' | tr '\n' '#' | sed 's/#$//' | sed 's/#/, /g')

rm -rf certs
mkdir -p certs

# Generate certificate authority
# -------------------------------

$openssl genpkey -algorithm Ed25519 -out certs/ca.key

$openssl req -x509 \
  -key certs/ca.key \
  -out certs/ca.cert \
  -sha256 \
  -days 3650 \
  -subj "/O=Decred Self-Signed Certificate Authority/CN=decred" \
  -addext "subjectAltName = DNS:localhost, $ips"

# Generate dcrd key
# -------------------------------

$openssl genpkey -algorithm Ed25519 -out certs/dcrd-rpc.key

$openssl req -x509 \
  -CA certs/ca.cert \
  -CAkey certs/ca.key \
  -key certs/dcrd-rpc.key \
  -out certs/dcrd-rpc.cert \
  -sha256 \
  -days 3650 \
  -subj "/O=dcrd autogenerated cert/CN=pop-os" \
  -addext "subjectAltName = DNS:pop-os, DNS:localhost, $ips"

# $openssl x509 -text -noout -in certs/dcrd-rpc.cert

# Generate dcrwallet key
# -------------------------------

$openssl genpkey -algorithm Ed25519 -out certs/dcrwallet-rpc.key

$openssl req -x509 \
  -CA certs/ca.cert \
  -CAkey certs/ca.key \
  -key certs/dcrwallet-rpc.key \
  -out certs/dcrwallet-rpc.cert \
  -sha256 \
  -days 3650 \
  -subj "/O=dcrwallet autogenerated cert/CN=pop-os" \
  -addext "subjectAltName = DNS:pop-os, DNS:localhost, $ips"

# $openssl x509 -text -noout -in certs/dcrwallet-rpc.cert

# Generate dcrwallet clients ca
# -----------------------------------------

#$openssl genpkey -algorithm ed25519 -out certs/dcrwallet-clients-ca.key
#
#$openssl req \
#  -new \
#  -x509 \
#  -sha256 \
#  -key certs/dcrwallet-clients-ca.key \
#  -out certs/dcrwallet-clients.pem \
#  -subj "/O=dcrwallet clients CA/CN=pop-os" \
#  -addext "subjectAltName = DNS:pop-os, DNS:localhost, $ips"

# Generate dcrctl client cert
# ------------------------------

#$openssl genpkey -algorithm Ed25519 -out certs/dcrctl-client-key.pem
#
#$openssl req -new -sha256 \
#  -key certs/dcrctl-client-key.pem \
#  -out certs/dcrctl-client.csr \
#  -subj "/O=dcrwallet client/CN=pop-os"

#$openssl x509 -req \
#  -in certs/dcrctl-client.csr \
#  -CA certs/dcrwallet-clients.pem \
#  -CAkey certs/dcrwallet-clients-ca.key \
#  -CAcreateserial \
#  -out certs/dcrctl-client.pem \
#  -days 3650 \
#  -sha256
