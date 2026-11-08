## gyro address create

Create address

### Synopsis

Create address

Create address. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/adres-defteri).

API: POST /addresses (operationId coreCreateAddress, scope Addresses/CreateAddress)

```
gyro address create [flags]
```

### Options

```
      --address string           Açık adres.
      --city string              Şehir. Listeden numara, liste boşsa serbest metin.
      --company string           Şirket bilgileri. Kurumsal kişide ad zorunludur. (JSON)
      --country string           Ülke kodu. Başvuru ucundan alınır.
      --data string              request body as JSON, @file or @- (stdin)
      --dry-run                  show the price quote / request without changing anything
      --email string             Geçerli bir e-posta.
  -h, --help                     help for create
      --idempotency-key string   idempotency key (automatic for money operations)
      --identity string          Kimlik numarası.
      --is-default               Bunu varsayılan profil yap.
      --kind string              Kişi türü. Öntanımlı olarak bireysel.
      --label string             Serbest etiket.
      --name string              İlgili kişinin adı.
      --notifications string     Kanal matrisi. Tam değiştirmedir; hiç gönderilmezse hepsi açılır. (JSON)
      --overwrite-invoices       Açık faturaları bu adrese yeniden yaz.
      --phone string             Kişinin telefonu.
      --set stringArray          extra body field key=value (repeatable)
      --state string             İl. Listeden numara, liste boşsa serbest metin.
      --surname string           Soyadı.
      --zipcode string           Posta kodu. En çok yirmi karakter.
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

* [gyro address](gyro_address.md)	 - Billing addresses

