## gyro lb metrics

Yük dengeleyici ölçümleri

### Synopsis

Yük dengeleyici ölçümleri

Açık bağlantı, yeni bağlantı/sn, istek/sn, bant genişliği (Mbit/s).

API: GET /cloud/load-balancers/{id}/metrics (operationId typed_get_load_balancers__id__metrics, scope services_read)

```
gyro lb metrics <ID> [flags]
```

### Options

```
  -h, --help            help for metrics
      --period string   24h | 7d | 30d
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

* [gyro lb](gyro_lb.md)	 - Yük dengeleyicileri listele

