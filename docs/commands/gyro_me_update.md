## gyro me update

Update me

### Synopsis

Update me

Update me. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/anahtarin-hesabi).

API: PATCH /me (operationId coreUpdateMe, scope Account/UpdateMe)

```
gyro me update [flags]
```

### Options

```
      --avatar string            Profil görseli. Yalnız kodlanmış içerik kabul edilir; adres reddedilir.
      --birthday string          Doğum tarihi.
      --company string           Şirket bilgileri. Hesabın kurumsal olmasını gerektirir. (JSON)
      --country string           Ülke kodu.
      --currency string          Görüntü para birimi kodu.
      --custom-fields string     Ek alanların değerleri. Kutu tipli alanlar dizi alır. (JSON)
      --data string              request body as JSON, @file or @- (stdin)
      --date-format string       Tarih biçimi. Boş göndermek hesabın kendi seçimini siler.
      --dry-run                  show the price quote / request without changing anything
      --email string             email
  -h, --help                     help for update
      --idempotency-key string   idempotency key (automatic for money operations)
      --identity string          Kimlik numarası.
      --kind string              Hesap türü. Bireysele geçmek şirket bilgilerini temizler.
      --landline-phone string    Sabit telefon.
      --language string          Dil kodu. Kurulumda bulunmalıdır.
      --name string              Adı. Boşaltılamaz.
      --phone string             phone
      --set stringArray          extra body field key=value (repeatable)
      --surname string           Soyadı. Boşaltılamaz.
      --timezone string          Saat dilimi. Boş göndermek hesabın kendi seçimini siler.
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

* [gyro me](gyro_me.md)	 - Get me

