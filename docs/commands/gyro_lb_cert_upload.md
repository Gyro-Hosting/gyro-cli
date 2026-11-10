## gyro lb cert upload

Upload a certificate

### Synopsis

Upload a certificate

PEM certificate chain + unencrypted private key. The key is sent to the load balancer only; it is never stored or logged by the panel.

API: POST /cloud/load-balancers/{id}/certificates (operationId typed_post_load_balancers__id__certificates, scope services_write)

```
gyro lb cert upload <ID> [flags]
```

### Options

```
      --certificate string       PEM chain
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for upload
      --idempotency-key string   idempotency key (automatic for money operations)
      --name string              name
      --private-key string       PEM key (unencrypted)
      --set stringArray          extra body field key=value (repeatable)
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

* [gyro lb cert](gyro_lb_cert.md)	 - List certificates

