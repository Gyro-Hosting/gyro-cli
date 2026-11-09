## gyro s3 object list

Kova klasöründeki nesneleri listele

### Synopsis

Kova klasöründeki nesneleri listele

Bir klasörün bir sayfası (200): `folders` ve `files` (anahtar, boyut, tarih). Devam için önceki sayfanın `next` değerini verin.

API: GET /cloud/storage/{id}/buckets/{bucket}/objects (operationId typed_get_storage__id__buckets__bucket__objects, scope services_read)

```
gyro s3 object list <ID> <BUCKET> [flags]
```

### Options

```
  -h, --help            help for list
      --next string     Continuation token
      --prefix string   Folder, ends with / (empty = root)
      --search string   Name prefix inside the folder
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

* [gyro s3 object](gyro_s3_object.md)	 - Kova klasöründeki nesneleri listele

