## gyro s3 object-lock get

Kovanın varsayılan saklaması

### Synopsis

Kovanın varsayılan saklaması

`object_lock` (açılışta seçilir, değişmez) ve `default_retention` (null = yok): yeni nesneler bu saklamayı kendiliğinden alır.

API: GET /cloud/storage/{id}/buckets/{bucket}/object-lock (operationId getStorageBucketObjectLock, scope services_read)

```
gyro s3 object-lock get <ID> <BUCKET> [flags]
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

* [gyro s3 object-lock](gyro_s3_object-lock.md)	 - Kovanın varsayılan saklaması

