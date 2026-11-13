## gyro container pipeline redeploy

Bu çalıştırmanın imajını yeniden dağıt (geri dönüş)

### Synopsis

Bu çalıştırmanın imajını yeniden dağıt (geri dönüş)

Önceki bir çalıştırmanın imajını hedef servislere dağıtan yeni çalıştırma başlatır. `"confirm": true` gerekir.

API: POST /cloud/containers/{id}/pipeline-runs/{run}/redeploy (operationId pipeline_redeploy, scope services_write)

```
gyro container pipeline redeploy <ID> <RUN> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for redeploy
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --targets strings          targets
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

