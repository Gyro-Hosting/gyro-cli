## gyro s3 object legal-hold set

Legal hold koy / kaldır

### Synopsis

Legal hold koy / kaldır

`status` on / off (yalnız kilitli kovalar). Legal hold, saklamadan bağımsız ve bitiş tarihsiz silmeyi engeller; kendiniz kaldırabilirsiniz.

API: PUT /cloud/storage/{id}/buckets/{bucket}/objects/legal-hold (operationId setStorageObjectLegalHold, scope services_write)

```
gyro s3 object legal-hold set <ID> <BUCKET> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --key string               key
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --status string            status (on|off)
      --version-id string        Empty = current version
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

* [gyro s3 object legal-hold](gyro_s3_object_legal-hold.md)	 - Nesne sürümünün legal hold durumu

