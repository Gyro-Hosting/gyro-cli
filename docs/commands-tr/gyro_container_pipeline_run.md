## gyro container pipeline run

Çalıştırma başlat

### Synopsis

Çalıştırma başlat

`ref`: dal (`main`), etiket (`tag:v1.2.0`) ya da tam ref; `sha`: isteğe bağlı 40 karakter commit. İkisi yoksa kuraldaki ilk dal derlenir.

API: POST /cloud/containers/{id}/pipelines/{pipeline}/runs (operationId pipeline_run, scope services_write)

```
gyro container pipeline run <ID> <PIPELINE> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for run
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --ref string               ref
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --sha string               sha
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

