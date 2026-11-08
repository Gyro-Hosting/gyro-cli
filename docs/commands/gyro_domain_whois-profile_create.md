## gyro domain whois-profile create

Create whois profile

### Synopsis

Create whois profile

Create whois profile. Domain operations run against the registrar module of the domain. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/alan-adi-kisileri-ve-gizlilik).

API: POST /domains/whois-profiles (operationId coreCreateWhoisProfile, scope Domains/CreateWhoisProfile)

```
gyro domain whois-profile create [flags]
```

### Options

```
      --address string           Adres satırı.
      --address2 string          İkinci adres satırı. İsteğe bağlı; anahtar yoksa boş satır kaydedilir.
      --city string              Şehir.
      --company string           Kurum adı. Tek isteğe bağlı alandır.
      --contact string           contactKaydedilecek kişi. İl burada zorunlu değildir. (JSON)
      --country string           Ülke kodu. Büyük harfe çevrilir.
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --email string             Geçerli bir e-posta.
      --first-name string        Kişinin adı.
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --last-name string         Kişinin soyadı.
      --name string              Profilin etiketi.
      --phone string             Telefon.
      --postal-code string       postal_code
      --set stringArray          extra body field key=value (repeatable)
      --state string             İl ya da eyalet. Alan adı kişilerinde zorunludur, kayıtlı profillerde değildir.
      --zip string               Posta kodu.
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

* [gyro domain whois-profile](gyro_domain_whois-profile.md)	 - List whois profiles

