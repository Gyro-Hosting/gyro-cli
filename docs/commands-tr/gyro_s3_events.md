## gyro s3 events

Kovanın olay webhook'larını listele

```
gyro s3 events [flags]
```

### Options

```
  -h, --help   help for events
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
* [gyro s3 events add](gyro_s3_events_add.md)	 - Kovaya olay webhook'u ekle
* [gyro s3 events delete](gyro_s3_events_delete.md)	 - Olay webhook'unu sil
* [gyro s3 events deliveries](gyro_s3_events_deliveries.md)	 - Webhook'un son teslimleri
* [gyro s3 events get](gyro_s3_events_get.md)	 - Olay webhook'unu getir
* [gyro s3 events list](gyro_s3_events_list.md)	 - Kovanın olay webhook'larını listele
* [gyro s3 events list-all](gyro_s3_events_list-all.md)	 - Hizmetin tüm olay webhook'larını listele
* [gyro s3 events rotate-secret](gyro_s3_events_rotate-secret.md)	 - İmza gizli anahtarını yenile
* [gyro s3 events test](gyro_s3_events_test.md)	 - Test olayı gönder
* [gyro s3 events update](gyro_s3_events_update.md)	 - Olay webhook'unu değiştir
* [gyro s3 events usage](gyro_s3_events_usage.md)	 - Bu ayın olay hakkı

