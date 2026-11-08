## gyro domain tld list

List tlds

### Synopsis

List tlds

List tlds. Domain operations run against the registrar module of the domain. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/alan-adi-edinme).

API: GET /domains/tlds (operationId coreGetTlds, scope Domains/GetTlds)

```
gyro domain tld list [flags]
```

### Options

```
  -h, --help         help for list
      --tld string   Tek bir uzantı. Listeyi daraltır ve süre başına fiyat tablosunu ekler.
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

* [gyro domain tld](gyro_domain_tld.md)	 - List tlds

