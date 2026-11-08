## gyro domain dns-record set

Update dns record

### Synopsis

Update dns record

Core WISECP client endpoint. Reference: https://dev.wisecp.com/tr/ad-sunuculari-ve-kayitlar

API: PUT /domains/{domain}/dns-records (operationId coreUpdateDnsRecord, scope Domains/UpdateDnsRecord)

```
gyro domain dns-record set <DOMAIN> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --identity string          Listeden gelen kayıt numarası. Liste bir numara döndürdüyse mutlaka gönderin.
      --name string              Kaydın adı.
      --priority int             Öncelik. Verilmezse kayıt, türü değişmediği sürece mevcut önceliğini korur.
      --set stringArray          extra body field key=value (repeatable)
      --ttl int                  Yaşam süresi. Verilmezse kayıt mevcut yaşam süresini korur.
      --type string              Kayıt türü.
      --value string             Yeni içerik.
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

* [gyro domain dns-record](gyro_domain_dns-record.md)	 - Get dns records

