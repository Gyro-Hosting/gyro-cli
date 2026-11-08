## gyro me notification set

Update notifications

### Synopsis

Update notifications

Core WISECP client endpoint. Reference: https://dev.wisecp.com/tr/anahtarin-hesabi

API: PUT /me/notifications (operationId coreUpdateNotifications, scope Account/UpdateNotifications)

```
gyro me notification set [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --domain string            Alan adı bildirimleri. (JSON)
      --dry-run                  show the price quote / request without changing anything
      --general string           general
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --invoices string          Fatura bildirimleri. E-posta ve kısa mesaj kanallarını taşır; eksik kanal kapalı sayılır. (JSON)
      --marketing string         Pazarlama bildirimleri. (JSON)
      --product string           Ürün ve hizmet bildirimleri. (JSON)
      --set stringArray          extra body field key=value (repeatable)
      --support string           Destek bildirimleri. (JSON)
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

* [gyro me notification](gyro_me_notification.md)	 - Get notifications

