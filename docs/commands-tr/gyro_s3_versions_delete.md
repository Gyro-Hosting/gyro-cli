## gyro s3 versions delete

Sürümü kalıcı sil

### Synopsis

Sürümü kalıcı sil

Bir sürümü ya da silme işaretini kalıcı siler (geri alınamaz). GOVERNANCE / COMPLIANCE saklamasında ya da legal hold'daki sürüm bitiş tarihiyle reddedilir. `confirm=true` gerekir.

API: DELETE /cloud/storage/{id}/buckets/{bucket}/versions (operationId deleteStorageVersion, scope services_write)

```
gyro s3 versions delete <ID> <BUCKET> [flags]
```

### Options

```
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for delete
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --key string               Object key
      --version-id string        Version id
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

* [gyro s3 versions](gyro_s3_versions.md)	 - Nesne sürümlerini listele

