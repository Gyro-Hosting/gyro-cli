## gyro domain forwarding set

Update url forwarding

### Synopsis

Update url forwarding

Core WISECP client endpoint. Reference: https://dev.wisecp.com/tr/yonlendirme-kurma

API: PUT /domains/{domain}/forwarding (operationId coreUpdateUrlForwarding, scope Domains/UpdateUrlForwarding)

```
gyro domain forwarding set <DOMAIN> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --method int               Yönlendirme türü. Kalıcı olan varsayılandır; başka bir değer ona düşer.
      --set stringArray          extra body field key=value (repeatable)
      --url string               Tam hedef adres. Şeması yönlendirmenin şeması olur; yazılmazsa güvensiz şema kullanılır.
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

* [gyro domain forwarding](gyro_domain_forwarding.md)	 - Get url forwarding

