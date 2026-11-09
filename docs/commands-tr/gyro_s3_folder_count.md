## gyro s3 folder count

Klasördeki nesneleri say

### Synopsis

Klasördeki nesneleri say

`prefix` altındaki nesne sayısı (`more` = 100000 üstü ya da süre sınırı).

API: GET /cloud/storage/{id}/buckets/{bucket}/folders/count (operationId typed_get_storage__id__buckets__bucket__folders_count, scope services_read)

```
gyro s3 folder count <ID> <BUCKET> [flags]
```

### Options

```
  -h, --help            help for count
      --prefix string   Folder, ends with /
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

* [gyro s3 folder](gyro_s3_folder.md)	 - Klasördeki nesneleri say

