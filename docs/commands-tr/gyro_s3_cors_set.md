## gyro s3 cors set

Kovanın CORS kurallarını değiştir

### Synopsis

Kovanın CORS kurallarını değiştir

Tüm kuralların yerine geçer. `rules` (dizi), `CORSRules` (AWS `put-bucket-cors` dosyası olduğu gibi: `--data @cors.json`) ya da `cors_json` (metin) gönderin. Alan adları AWS (`AllowedOrigins`) ya da snake_case olabilir. Sınırlar: 100 kural, kural başına 50 köken/başlık, 16 KB, yöntemler GET/PUT/POST/DELETE/HEAD, `MaxAgeSeconds` 0-86400, kökende en çok bir `*`. Boş liste kuralları kaldırır. Kurallar S3 adresinde ve statik web sitesinde (özel alan adları dahil) geçerlidir.

API: PUT /cloud/storage/{id}/buckets/{bucket}/cors (operationId typed_put_storage__id__buckets__bucket__cors, scope services_write)

```
gyro s3 cors set <ID> <BUCKET> [flags]
```

### Options

```
      --cors-json string         Alternatif: kurallar JSON metni olarak
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --rules string             CORS kuralları (JSON dizi) (JSON)
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

* [gyro s3 cors](gyro_s3_cors.md)	 - Kovanın CORS kuralları

