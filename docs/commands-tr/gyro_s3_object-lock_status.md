## gyro s3 object-lock status

Korumalı Depolama: nesne kilidi özeti

### Synopsis

Korumalı Depolama: nesne kilidi özeti

Kilitli kova hakkı (Nesne Kilidi ek hizmeti) ve kullanımı, kovaların varsayılan saklaması ve son kilit taraması (COMPLIANCE / GOVERNANCE / legal hold sürüm sayısı, en geç bitiş). Yalnız Korumalı Depolama: standart Nesne Depolama hizmeti 409 `feature_not_available` döner.

API: GET /cloud/storage/{id}/object-lock (operationId getStorageObjectLock, scope services_read)

```
gyro s3 object-lock status <ID> [flags]
```

### Options

```
  -h, --help   help for status
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

