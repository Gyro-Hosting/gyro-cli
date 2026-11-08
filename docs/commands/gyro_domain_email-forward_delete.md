## gyro domain email-forward delete

Delete email forward (path: {domain})

### Synopsis

Delete email forward (path: {domain})

Delete email forward (path: {domain}). Domain operations run against the registrar module of the domain. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/yonlendirme-kurma).

API: DELETE /domains/{domain}/email-forwards (operationId coreDeleteEmailForward, scope Domains/DeleteEmailForward)

```
gyro domain email-forward delete <DOMAIN> [flags]
```

### Options

```
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for delete
      --idempotency-key string   idempotency key (automatic for money operations)
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

* [gyro domain email-forward](gyro_domain_email-forward.md)	 - List email forwards (path: {domain})

