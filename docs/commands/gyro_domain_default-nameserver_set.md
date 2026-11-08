## gyro domain default-nameserver set

Update default nameservers

### Synopsis

Update default nameservers

Core WISECP client endpoint. Reference: https://dev.wisecp.com/tr/ad-sunuculari-ve-kayitlar

API: PUT /domains/default-nameservers (operationId coreUpdateDefaultNameservers, scope Domains/UpdateDefaultNameservers)

```
gyro domain default-nameserver set [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --nameservers string       Varsayılan set. İkiden dörde kadar geçerli sunucu adı; boşluklar kırpılır ve küçük harfe çevrilir.
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

* [gyro domain default-nameserver](gyro_domain_default-nameserver.md)	 - Get default nameservers

