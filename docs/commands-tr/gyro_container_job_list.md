## gyro container job list

Tek seferlik görevleri listele

### Synopsis

Tek seferlik görevleri listele

Son 50 görev (yeniden eskiye), `active` (şu an çalışan) ve `limit` (eşzamanlı görev sınırı). `?service=<ad>` süzer.

API: GET /cloud/containers/{id}/jobs (operationId listContainerJobs, scope services_read)

```
gyro container job list <ID> [flags]
```

### Options

```
  -h, --help             help for list
      --service string   only the jobs of this service
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

* [gyro container job](gyro_container_job.md)	 - Tek seferlik görevleri listele

