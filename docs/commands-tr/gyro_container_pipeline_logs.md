## gyro container pipeline logs

Çalıştırma logu

### Synopsis

Çalıştırma logu

Derlenirken canlı, sonra kayıtlı (en çok 512 KB). `tail`: son N satır.

API: GET /cloud/containers/{id}/pipeline-runs/{run}/logs (operationId pipeline_logs, scope services_read)

```
gyro container pipeline logs <ID> <RUN> [flags]
```

### Options

```
  -h, --help       help for logs
      --tail int   last N lines (default 3000)
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

* [gyro container pipeline](gyro_container_pipeline.md)	 - Pipeline'ları listele

