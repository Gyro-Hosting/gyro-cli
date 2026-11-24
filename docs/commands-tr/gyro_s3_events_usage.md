## gyro s3 events usage

Bu ayın olay hakkı

### Synopsis

Bu ayın olay hakkı

Ücretsiz hak (kova başına × kova), bu ay alınan olay paketleri, kullanılan ve düşen olaylar (hak / hız), teslim / başarısız sayıları.

API: GET /cloud/storage/{id}/events/usage (operationId getStorageEventUsage, scope services_read)

```
gyro s3 events usage <ID> [flags]
```

### Options

```
  -h, --help   help for usage
```

### Options inherited from parent commands

```
      --account gyro account list   panel kullanıcısı olduğunuz başka bir hesapta işlem yap (gyro account list id'si; "self" = kendi hesabınız; ortam GYRO_ACCOUNT, profil varsayılanı: `gyro account use`)
      --api-url string              API taban adresi (ortam GYRO_API_URL)
      --debug                       HTTP isteklerini stderr'e yaz (sırlar maskeli)
      --lang string                 mesaj dili: tr ya da en (varsayılan yerel ayar)
      --no-color                    renkleri kapat (NO_COLOR da olur)
      --no-header                   tablo başlığını yazma
  -o, --output string               çıktı biçimi: table, json, yaml, text
  -p, --profile string              kullanılacak profil (ortam GYRO_PROFILE)
      --query string                sonuca uygulanacak JMESPath sorgusu
  -q, --quiet                       yalnız kimlikleri yaz
      --timeout duration            toplam zaman aşımı, ör. 30s, 10m
      --wide                        tüm tablo sütunlarını göster
  -y, --yes                         onay sorma
```

### SEE ALSO

* [gyro s3 events](gyro_s3_events.md)	 - Kovanın olay webhook'larını listele

