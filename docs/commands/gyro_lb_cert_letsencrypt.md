## gyro lb cert letsencrypt

Request a free Let's Encrypt certificate

### Synopsis

Request a free Let's Encrypt certificate

Up to 10 names whose A record points to this load balancer. Your targets must redirect /.well-known/acme-challenge/<token> to the redirect base shown in GET certificates; answer `acme=waiting` lists the per-name check, fix the redirect and call verify.

API: POST /cloud/load-balancers/{id}/certificates/letsencrypt (operationId typed_post_load_balancers__id__certificates_letsencrypt, scope services_write)

```
gyro lb cert letsencrypt <ID> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --domains string           Comma separated names
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for letsencrypt
      --idempotency-key string   idempotency key (automatic for money operations)
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

