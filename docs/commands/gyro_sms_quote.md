## gyro sms quote

Quote sms

### Synopsis

Quote sms

Quote sms. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/mesaj-gonderme).

API: POST /sms/quote (operationId coreQuoteSms, scope Sms/QuoteSms)

```
gyro sms quote [flags]
```

### Options

```
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
  -h, --help                     help for quote
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

