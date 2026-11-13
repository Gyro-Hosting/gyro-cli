## gyro service list

Tüm ürün hatlarındaki hizmetleri listele

### Synopsis

Tüm ürün hatlarındaki hizmetleri listele

API: GET /cloud/services (operationId listServices, scope services_read)

```
gyro service list [flags]
```

### Options

```
      --all                         tüm sayfaları getir
      --customer-id int             Reseller main key only
  -h, --help                        help for list
      --limit int                   sayfa başına kayıt (en çok 100)
      --page int                    sayfa numarası (default 1)
      --product-line string          (cloud_server|vds|dedicated_server|mail|monitor|storage|license|app)
      --project project_not_found   Only services in this project (0 = default project). Unknown project: 404 project_not_found.
      --status string               
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

* [gyro service](gyro_service.md)	 - Tüm ürün hatlarındaki hizmetler (genel köprü)

