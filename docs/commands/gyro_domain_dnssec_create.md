## gyro domain dnssec create

Create dnssec record (path: {domain})

### Synopsis

Create dnssec record (path: {domain})

Create dnssec record (path: {domain}). Domain operations run against the registrar module of the domain. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/ad-sunuculari-ve-kayitlar).

API: POST /domains/{domain}/dnssec (operationId coreCreateDnssecRecord, scope Domains/CreateDnssecRecord)

```
gyro domain dnssec create <DOMAIN> [flags]
```

### Options

```
      --algorithm int                    Anahtar algoritmasının numarası. Modülün bildirdiği kümede olmalıdır.
      --data string                      request body as JSON, @file or @- (stdin)
      --digest string                    İmza özeti. DNSKEY isteyen servis sağlayıcı burada açık anahtarı (base64) alır.
      --digest-type meta.dnssec_fields   Özet algoritmasının numarası ya da meta.dnssec_fields'in adlandırdığı anahtar bayrağı veya protokol. Modülün bildirdiği kümede olmalıdır.
      --dry-run                          show the price quote / request without changing anything
  -h, --help                             help for create
      --idempotency-key string           idempotency key (automatic for money operations)
      --key-tag dnskey-keytag-flags      Anahtarın etiketi. dnskey-keytag-flags'te anahtar bayrağını taşır.
      --set stringArray                  extra body field key=value (repeatable)
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

* [gyro domain dnssec](gyro_domain_dnssec.md)	 - List dnssec records (path: {domain})

