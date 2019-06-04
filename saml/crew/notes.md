# Notes

## GoDocs

<https://godoc.org/github.com/edaniels/go-saml>

## Certs

To generate the key and certificate:

```
openssl req -x509 -newkey rsa:2048 -keyout myservice.key -out myservice.cert -days 365 -nodes -subj "/CN=myservice.example.com"
```

To register with an IdP, run this:

```
curl localhost:8000/saml/metadata > /tmp/digi.xml
```

Then go to here to upload the file:

<https://samltest.id/upload.php>

## Vulnerabilities

<https://github.com/yogisec/VulnerableSAMLApp>
