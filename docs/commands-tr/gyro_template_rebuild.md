## gyro template rebuild

Sunucuyu şablonla yeniden kur

### Synopsis

Sunucuyu şablonla yeniden kur

tr-ist-1 sunucunuzun kök diskini şablonla değiştirir (disk en az `min_disk_gb` olmalı); sunucunun SSH anahtarları yeniden kurulur, yeni root parolası döner (Idempotency-Key gönderilirse dönmez). `"confirm": true` gerekir.

API: POST /cloud/templates/{id}/rebuild (operationId rebuildFromTemplate, scope cloud_write)

```
gyro template rebuild <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for rebuild
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --server-id int            server_id
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --user-data string         user_data
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

* [gyro template](gyro_template.md)	 - Şablonlarınızı listele

