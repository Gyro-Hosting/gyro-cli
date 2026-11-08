## gyro dns record add

Create dns record

### Synopsis

Create dns record

Core WISECP client endpoint. Reference: https://dev.wisecp.com/tr/ad-sunuculari-ve-kayitlar

API: POST /domains/{domain}/dns-records (operationId coreCreateDnsRecord, scope Domains/CreateDnsRecord)

```
gyro dns record add <DOMAIN> [flags]
```

### Examples

```
  gyro dns record add example.com --type A --name www --value 203.0.113.10 --ttl 3600
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for add
      --idempotency-key string   idempotency key (automatic for money operations)
      --identity string          identity
      --name string              Kaydın yanıtlayacağı ad. Kabul edilen biçim sağlayıcıya göre değişir.
      --priority int             Öncelik. Posta ve hizmet kayıtları dışında yok sayılır.
      --set stringArray          extra body field key=value (repeatable)
      --ttl int                  Yaşam süresi. Verilmezse sağlayıcının varsayılanı kullanılır.
      --type string              Kayıt türü.
      --value string             Kaydın içeriği.
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

* [gyro dns record](gyro_dns_record.md)	 - List, add, change and delete DNS records

