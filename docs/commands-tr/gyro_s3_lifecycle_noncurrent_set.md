## gyro s3 lifecycle noncurrent set

Eski sürüm silme süresini ayarla

### Synopsis

Eski sürüm silme süresini ayarla

`days` 1-3650 ya da 0 (eski sürümler durur). Kilitli sürümler saklaması bitince silinir. Kovanın diğer yaşam döngüsü kuralları korunur.

API: PUT /cloud/storage/{id}/buckets/{bucket}/lifecycle/noncurrent (operationId setStorageNoncurrentExpiration, scope services_write)

```
gyro s3 lifecycle noncurrent set <ID> <BUCKET> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --days int                 days
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
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

