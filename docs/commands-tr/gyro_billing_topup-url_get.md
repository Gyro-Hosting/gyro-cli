## gyro billing topup-url get

Browser link of the balance top-up page (top-ups are invoiced with VAT; usage is not)

### Synopsis

Browser link of the balance top-up page (top-ups are invoiced with VAT; usage is not)

Prepaid (hourly, from the account balance) services: Gyro Containers (ECS) and Gyro Cloud Istanbul (tr-ist-1). Times are Europe/Istanbul. Usage is never invoiced: invoices are issued for balance top-ups.

API: GET /cloud/billing/topup-url (operationId getPrepaidTopupUrl, scope billing_read)

```
gyro billing topup-url get [flags]
```

### Options

```
  -h, --help   help for get
```

### Options inherited from parent commands

```
      --api-url string     API taban adresi (ortam GYRO_API_URL)
      --debug              HTTP isteklerini stderr'e yaz (sırlar maskeli)
      --lang string        mesaj dili: tr ya da en (varsayılan yerel ayar)
      --no-color           renkleri kapat (NO_COLOR da olur)
      --no-header          tablo başlığını yazma
  -o, --output string      çıktı biçimi: table, json, yaml, text
  -p, --profile string     kullanılacak profil (ortam GYRO_PROFILE)
      --query string       sonuca uygulanacak JMESPath sorgusu
  -q, --quiet              yalnız kimlikleri yaz
      --timeout duration   toplam zaman aşımı, ör. 30s, 10m
      --wide               tüm tablo sütunlarını göster
  -y, --yes                onay sorma
```

### SEE ALSO

* [gyro billing topup-url](gyro_billing_topup-url.md)	 - Browser link of the balance top-up page (top-ups are invoiced with VAT; usage is not)

