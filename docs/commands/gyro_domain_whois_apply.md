## gyro domain whois apply

Apply whois profile

### Synopsis

Apply whois profile

Core WISECP client endpoint. Reference: https://dev.wisecp.com/tr/alan-adi-kisileri-ve-gizlilik

API: POST /domains/{domain}/whois/apply (operationId coreApplyWhoisProfile, scope Domains/ApplyWhoisProfile)

```
gyro domain whois apply <DOMAIN> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for apply
      --idempotency-key string   idempotency key (automatic for money operations)
      --profile-id int           Uygulanacak profilin numarası.
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

* [gyro domain whois](gyro_domain_whois.md)	 - Get whois contacts

