## gyro domain transfer-status get

Get transfer status (path: {domain})

### Synopsis

Get transfer status (path: {domain})

Get transfer status (path: {domain}). Domain operations run against the registrar module of the domain. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/alan-adi-edinme).

API: GET /domains/{domain}/transfer-status (operationId coreGetTransferStatus, scope Domains/GetTransferStatus)

```
gyro domain transfer-status get <DOMAIN> [flags]
```

### Options

```
  -h, --help   help for get
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

* [gyro domain transfer-status](gyro_domain_transfer-status.md)	 - Get transfer status (path: {domain})

