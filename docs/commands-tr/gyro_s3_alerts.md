## gyro s3 alerts

Tüm kovaların doluluk uyarıları

```
gyro s3 alerts [flags]
```

### Options

```
  -h, --help   help for alerts
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
* [gyro s3 alerts delete](gyro_s3_alerts_delete.md)	 - Kovanın doluluk uyarısını kaldır
* [gyro s3 alerts get](gyro_s3_alerts_get.md)	 - Kovanın doluluk uyarısı
* [gyro s3 alerts list](gyro_s3_alerts_list.md)	 - Tüm kovaların doluluk uyarıları
* [gyro s3 alerts set](gyro_s3_alerts_set.md)	 - Kovanın doluluk uyarısını ayarla

