## gyro s3 events deliveries

Webhook'un son teslimleri

### Synopsis

Webhook'un son teslimleri

Son 20 teslim kaydı (teslim / başarısız / yeniden denenecek, HTTP kodu, hata, deneme), en yeni önce; 30 gün saklanır.

API: GET /cloud/storage/{id}/events/{event}/deliveries (operationId listStorageEventDeliveries, scope services_read)

```
gyro s3 events deliveries <ID> <EVENT> [flags]
```

### Options

```
  -h, --help   help for deliveries
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

* [gyro s3 events](gyro_s3_events.md)	 - Kovanın olay webhook'larını listele

