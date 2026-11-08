## gyro sms send

Send sms

### Synopsis

Send sms

**Spends money** (paid from balance or a stored card; prefer the GyroAPI equivalents with confirm + idempotency where available). Core WISECP client endpoint. Reference: https://dev.wisecp.com/tr/mesaj-gonderme

API: POST /sms/send (operationId coreSendSms, scope Sms/SendSms)

```
gyro sms send [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for send
      --idempotency-key string   idempotency key (automatic for money operations)
      --message string           Mesaj metni. Parça tavanını aşan metin kırpılır ve kırpıldığı bildirilir.
      --numbers strings          Alıcı numaralar. Uluslararası biçimde; tek bir metin olarak gönderilirse satır, virgül ve noktalı virgülden bölünür.
      --sender string            Gönderen başlığı. Hesabın etkin başlıklarından biri olmalıdır.
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

* [gyro sms](gyro_sms.md)	 - SMS messages

