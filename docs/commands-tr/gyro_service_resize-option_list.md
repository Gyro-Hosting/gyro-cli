## gyro service resize-option list

Hizmetin yükseltilebileceği ürünler ve fiyatları

### Synopsis

Hizmetin yükseltilebileceği ürünler ve fiyatları

Prepaid services answer `billing: prepaid`; each option has `plan`, `current` and one `hourly` price with `monthly_cap`.

API: GET /cloud/services/{id}/resize-options (operationId getResizeOptions, scope services_read)

```
gyro service resize-option list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

* [gyro service resize-option](gyro_service_resize-option.md)	 - Hizmetin yükseltilebileceği ürünler ve fiyatları

