## gyro s3 versions list

Nesne sürümlerini listele

### Synopsis

Nesne sürümlerini listele

Bir klasördeki sürümler ve silme işaretleri (sayfa 50, ayraç `/`; alt klasörler `folders`). `truncated` true iken önceki sayfanın `key_marker` + `version_id_marker` değerleriyle devam edin. Kilitli kovalarda her sürümde `lock_mode`, `retain_until`, `legal_hold` bulunur.

API: GET /cloud/storage/{id}/buckets/{bucket}/versions (operationId listStorageVersions, scope services_read)

```
gyro s3 versions list <ID> <BUCKET> [flags]
```

### Options

```
  -h, --help                       help for list
      --key-marker string          Continuation: next_key_marker of the previous page
      --prefix string              Folder or key prefix (empty = root)
      --version-id-marker string   Continuation: next_version_id_marker of the previous page
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

