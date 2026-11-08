## gyro domain child-nameserver create

Create child nameserver

### Synopsis

Create child nameserver

Core WISECP client endpoint. Reference: https://dev.wisecp.com/tr/ad-sunuculari-ve-kayitlar

API: POST /domains/{domain}/child-nameservers (operationId coreCreateChildNameserver, scope Domains/CreateChildNameserver)

```
gyro domain child-nameserver create <DOMAIN> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
      --host string              Sunucunun tam adı. Küçük harfe çevrilir.
      --idempotency-key string   idempotency key (automatic for money operations)
      --ip string                Sunucunun adresi. Biçimi doğrulanır.
      --ns string                ns
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

* [gyro domain child-nameserver](gyro_domain_child-nameserver.md)	 - Get child nameservers

