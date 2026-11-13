## gyro container pipeline deploy-key create

SSH deploy key oluştur

### Synopsis

SSH deploy key oluştur

SSH deposu için ed25519 anahtar çifti üretir, yalnız açık anahtarı döner (depoya salt okuma deploy key olarak ekleyin). Özel anahtar şifreli saklanır, hiç dönmez. `"rotate": true` yeniler.

API: PUT /cloud/containers/{id}/pipelines/{pipeline}/deploy-key (operationId pipeline_deploy_key_create, scope services_write)

```
gyro container pipeline deploy-key create <ID> <PIPELINE> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --rotate                   rotate
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

* [gyro container pipeline deploy-key](gyro_container_pipeline_deploy-key.md)	 - SSH deploy key oluştur

