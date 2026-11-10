## gyro lb cert

List certificates

```
gyro lb cert [flags]
```

### Options

```
  -h, --help   help for cert
```

### Options inherited from parent commands

```
      --api-url string     API base URL (env GYRO_API_URL)
      --debug              log HTTP requests to stderr (secrets masked)
      --lang string        message language: tr or en (default from locale)
      --no-color           disable colours (also NO_COLOR)
      --no-header          omit the table header
  -o, --output string      output format: table, json, yaml, text
  -p, --profile string     profile to use (env GYRO_PROFILE)
      --query string       JMESPath query applied to the result
  -q, --quiet              print only ids
      --timeout duration   overall timeout, e.g. 30s, 10m
      --wide               show all table columns
  -y, --yes                do not ask for confirmation
```

### SEE ALSO

* [gyro lb](gyro_lb.md)	 - List load balancers
* [gyro lb cert cancel](gyro_lb_cert_cancel.md)	 - Cancel a pending Let's Encrypt request
* [gyro lb cert delete](gyro_lb_cert_delete.md)	 - Delete a certificate
* [gyro lb cert letsencrypt](gyro_lb_cert_letsencrypt.md)	 - Request a free Let's Encrypt certificate
* [gyro lb cert list](gyro_lb_cert_list.md)	 - List certificates
* [gyro lb cert upload](gyro_lb_cert_upload.md)	 - Upload a certificate
* [gyro lb cert verify](gyro_lb_cert_verify.md)	 - Continue a pending Let's Encrypt request

