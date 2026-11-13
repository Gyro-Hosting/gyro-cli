## gyro container pipeline host-key scan

SSH sunucu anahtarı

### Synopsis

SSH sunucu anahtarı

Deponun SSH sunucusuna bağlanıp anahtar parmak izini döner. GitHub / GitLab / Bitbucket anahtarları sabittir. Diğer sunucularda parmak izini kendi kaydınızla karşılaştırıp PUT ile onaylayın.

API: GET /cloud/containers/{id}/pipelines/{pipeline}/host-key (operationId pipeline_host_key_scan, scope services_read)

```
gyro container pipeline host-key scan <ID> <PIPELINE> [flags]
```

### Options

```
  -h, --help   help for scan
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

* [gyro container pipeline host-key](gyro_container_pipeline_host-key.md)	 - SSH sunucu anahtarı

