## gyro domain whois set

Update whois contacts (path: {domain})

### Synopsis

Update whois contacts (path: {domain})

Update whois contacts (path: {domain}). Domain operations run against the registrar module of the domain. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/alan-adi-kisileri-ve-gizlilik).

API: PUT /domains/{domain}/whois (operationId coreUpdateWhoisContacts, scope Domains/UpdateWhoisContacts)

```
gyro domain whois set <DOMAIN> [flags]
```

### Options

```
      --address string           Adres satırı.
      --address2 string          İkinci adres satırı (daire, kat, blok). Anahtar gönderilmezse sağlayıcıdaki satır korunur; boş metin satırı siler.
      --city string              Şehir.
      --company string           Kurum adı. Tek isteğe bağlı alandır.
      --contact string           contactYazılacak kişi. (JSON)
      --country string           Ülke kodu. Büyük harfe çevrilir.
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --email string             Geçerli bir e-posta.
      --first-name string        Kişinin adı.
  -h, --help                     help for set
      --idempotency-key string   idempotency key (automatic for money operations)
      --last-name string         Kişinin soyadı.
      --phone string             Telefon.
      --postal-code string       postal_code
      --role string              Hangi rollere yazılacağı. Verilmezse dört rolün hepsine yazılır.
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

* [gyro domain whois](gyro_domain_whois.md)	 - Get whois contacts (path: {domain})

