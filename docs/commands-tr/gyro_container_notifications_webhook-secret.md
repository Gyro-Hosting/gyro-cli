## gyro container notifications webhook-secret

Webhook imza anahtarı

### Synopsis

Webhook imza anahtarı

İmza anahtarını döndürür (`rotate: true` yenisini üretir).

API: POST /cloud/containers/{id}/notifications/webhook-secret (operationId typed_post_containers__id__notifications_webhook_secret, scope services_write)

```
gyro container notifications webhook-secret <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for webhook-secret
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

* [gyro container notifications](gyro_container_notifications.md)	 - Bildirim ayarları ve süren uyarılar

