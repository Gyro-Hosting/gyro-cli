## gyro container pipeline upload

Arşiv derlemesi için yükleme adresi

### Synopsis

Arşiv derlemesi için yükleme adresi

15 dakika geçerli, tam `size` bayt için tek `PUT` adresi (`url`) döner (tar.gz / tgz / zip, en çok 200 MB; ör. `curl -T site.tar.gz "<url>"`). Sonra `POST .../runs {"upload_id": ...}` ile derleyin. Arşivler 24 saatte silinir.

API: POST /cloud/containers/{id}/pipelines/{pipeline}/uploads (operationId pipeline_upload, scope services_write)

```
gyro container pipeline upload <ID> <PIPELINE> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for upload
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --size int                 archive size in bytes (1 - 209715200)
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

