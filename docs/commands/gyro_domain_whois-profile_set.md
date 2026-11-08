## gyro domain whois-profile set

Update whois profile

### Synopsis

Update whois profile

Core WISECP client endpoint. Reference: https://dev.wisecp.com/tr/alan-adi-kisileri-ve-gizlilik

API: PUT /domains/whois-profiles/{pid} (operationId coreUpdateWhoisProfile, scope Domains/UpdateWhoisProfile)

```
gyro domain whois-profile set <PID> [flags]
```

### Options

```
      --address string           Adres satırı.
      --address2 string          İkinci adres satırı. Profil bütün olarak değiştirilir; anahtar yoksa boş satır kaydedilir.
      --city string              Şehir.
      --company string           Kurum adı. Tek isteğe bağlı alandır.
      --contact string           contactTam kişi. Kayıtlı kişi bütünüyle değiştirilir. (JSON)
      --country string           Ülke kodu. Büyük harfe çevrilir.
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --email string             Geçerli bir e-posta.
      --first-name string        Kişinin adı.
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --last-name string         Kişinin soyadı.
      --name string              Yeni etiket.
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

* [gyro domain whois-profile](gyro_domain_whois-profile.md)	 - Get whois profiles

