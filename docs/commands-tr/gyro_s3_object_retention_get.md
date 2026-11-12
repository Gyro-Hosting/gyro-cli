## gyro s3 object retention get

Nesne sürümünün saklaması

### Synopsis

Nesne sürümünün saklaması

Bir sürümün kilit kipi, bitiş tarihi ve legal hold durumu (`version_id` boş = güncel sürüm).

API: GET /cloud/storage/{id}/buckets/{bucket}/objects/retention (operationId getStorageObjectRetention, scope services_read)

```
gyro s3 object retention get <ID> <BUCKET> [flags]
```

### Options

```
  -h, --help                help for get
      --key string          Object key
      --version-id string   Version id (empty = current)
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

* [gyro s3 object retention](gyro_s3_object_retention.md)	 - Nesne sürümünün saklaması

