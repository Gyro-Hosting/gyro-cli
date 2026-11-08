## gyro domain dns-record create

Create dns record (path: {domain})

### Synopsis

Create dns record (path: {domain})

Create dns record (path: {domain}). Domain operations run against the registrar module of the domain. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/ad-sunuculari-ve-kayitlar).

API: POST /domains/{domain}/dns-records (operationId coreCreateDnsRecord, scope Domains/CreateDnsRecord)

```
gyro domain dns-record create <DOMAIN> [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for create
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

* [gyro domain dns-record](gyro_domain_dns-record.md)	 - List dns records (path: {domain})

