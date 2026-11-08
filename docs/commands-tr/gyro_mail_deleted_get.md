## gyro mail deleted get

deleted-mail (GyroMail)

### Synopsis

deleted-mail (GyroMail)

Genel köprünün kısayolu: araç `deleted-mail`. Gövde alanları aracın kendi alanlarıdır.

API: GET /cloud/mail/{id}/deleted (operationId typed_get_mail__id__deleted, scope services_read)

```
gyro mail deleted get <ID> [flags]
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

* [gyro mail deleted](gyro_mail_deleted.md)	 - deleted-mail (GyroMail)

