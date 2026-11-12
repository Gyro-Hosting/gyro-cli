## gyro s3 lifecycle noncurrent get

Kovanın eski sürüm silme kuralı

### Synopsis

Kovanın eski sürüm silme kuralı

Eski (güncel olmayan) sürümlerin kaç gün sonra silineceği (yeni kovalarda 30). 0 / null = siz silene kadar durur (kotadan düşer).

API: GET /cloud/storage/{id}/buckets/{bucket}/lifecycle/noncurrent (operationId getStorageNoncurrentExpiration, scope services_read)

```
gyro s3 lifecycle noncurrent get <ID> <BUCKET> [flags]
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

* [gyro s3 lifecycle noncurrent](gyro_s3_lifecycle_noncurrent.md)	 - Kovanın eski sürüm silme kuralı

