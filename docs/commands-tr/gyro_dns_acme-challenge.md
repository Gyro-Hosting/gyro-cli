## gyro dns acme-challenge

ACME doğrulama TXT değeri ekle

```
gyro dns acme-challenge [flags]
```

### Options

```
  -h, --help   help for acme-challenge
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

* [gyro dns](gyro_dns.md)	 - Alan adlarınızın DNS kayıtları
* [gyro dns acme-challenge add](gyro_dns_acme-challenge_add.md)	 - ACME doğrulama TXT değeri ekle
* [gyro dns acme-challenge remove](gyro_dns_acme-challenge_remove.md)	 - ACME doğrulama TXT değerini sil
* [gyro dns acme-challenge status](gyro_dns_acme-challenge_status.md)	 - Doğrulama kaydı yayında mı?

