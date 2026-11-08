## gyro domain whois-privacy set

Update whois privacy

### Synopsis

Update whois privacy

Core WISECP client endpoint. Reference: https://dev.wisecp.com/tr/alan-adi-kisileri-ve-gizlilik

API: PUT /domains/{domain}/whois-privacy (operationId coreUpdateWhoisPrivacy, scope Domains/UpdateWhoisPrivacy)

```
gyro domain whois-privacy set <DOMAIN> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --enabled                  Gizlilik açık mı. Gerçek bir mantıksal değer olmalıdır.
  -h, --help                     help for set
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

* [gyro domain whois-privacy](gyro_domain_whois-privacy.md)	 - Update whois privacy

