## gyro domain dnssec get

İmza Kayıtlarını Listeleme

### Synopsis

İmza Kayıtlarını Listeleme

Alan adının imza doğrulama kayıtlarını döndürür. — sağlayıcıdan okur

API: GET /domains/{domain}/dnssec (operationId coreGetDnssecRecords, scope Domains/GetDnssecRecords)

```
gyro domain dnssec get <DOMAIN> [flags]
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

* [gyro domain dnssec](gyro_domain_dnssec.md)	 - İmza Kayıtlarını Listeleme

