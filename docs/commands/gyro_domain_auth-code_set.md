## gyro domain auth-code set

Save auth code

### Synopsis

Save auth code

Core WISECP client endpoint. Reference: https://dev.wisecp.com/tr/alan-adi-edinme

API: PUT /domains/{domain}/auth-code (operationId coreSaveAuthCode, scope Domains/SaveAuthCode)

```
gyro domain auth-code set <DOMAIN> [flags]
```

### Options

```
      --code string              Yeni transfer kodu. İşaretleme ayıklanır ve boşluklar kırpılır.
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
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

* [gyro domain auth-code](gyro_domain_auth-code.md)	 - Send auth code

