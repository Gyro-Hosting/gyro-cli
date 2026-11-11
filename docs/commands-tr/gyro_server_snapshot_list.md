## gyro server snapshot list

Snapshot'ları ve kotayı listele

### Synopsis

Snapshot'ları ve kotayı listele

tr-ist-1: `quota` = number of snapshots allowed on the plan (add-on), snapshot ids are stable numbers derived by the API. Prepaid tr-ist-1 servers also return `pricing` {billing, free, extra_monthly_eur, extra_hourly_eur, billable}: the first snapshot(s) are free, each extra snapshot is charged hourly from the balance while it exists.

API: GET /cloud/servers/{id}/snapshots (operationId listSnapshots, scope cloud_read)

```
gyro server snapshot list <ID> [flags]
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

* [gyro server snapshot](gyro_server_snapshot.md)	 - Snapshot'ları ve kotayı listele

