## gyro s3 object-lock

Kovanın varsayılan saklaması

```
gyro s3 object-lock [flags]
```

### Options

```
  -h, --help   help for object-lock
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

* [gyro s3](gyro_s3.md)	 - S3 nesne depolama: kovalar ve erişim anahtarları
* [gyro s3 object-lock get](gyro_s3_object-lock_get.md)	 - Kovanın varsayılan saklaması
* [gyro s3 object-lock scan](gyro_s3_object-lock_scan.md)	 - Kilitli sürümleri şimdi tara
* [gyro s3 object-lock set](gyro_s3_object-lock_set.md)	 - Kovanın varsayılan saklamasını ayarla
* [gyro s3 object-lock status](gyro_s3_object-lock_status.md)	 - Korumalı Depolama: nesne kilidi özeti

