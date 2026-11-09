## gyro s3 folder delete

Klasörü içindekilerle sil

### Synopsis

Klasörü içindekilerle sil

`prefix` (sorgu) altındaki tüm nesneleri siler. Büyük klasörler parça parça silinir: `more` true olduğu sürece tekrarlayın. `"confirm": true` gerekir.

API: DELETE /cloud/storage/{id}/buckets/{bucket}/folders (operationId typed_delete_storage__id__buckets__bucket__folders, scope services_write)

```
gyro s3 folder delete <ID> <BUCKET> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for delete
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --prefix string            Folder, ends with /
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

* [gyro s3 folder](gyro_s3_folder.md)	 - Klasördeki nesneleri say

