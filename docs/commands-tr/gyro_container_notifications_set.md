## gyro container notifications set

Bildirim kanallarını ayarla

### Synopsis

Bildirim kanallarını ayarla

Yalnız değişecek alanları gönderin. Webhook adresi ilk kez girilince imza anahtarı üretilir (POST .../notifications/webhook-secret ile alınır).

API: PUT /cloud/containers/{id}/notifications (operationId typed_put_containers__id__notifications, scope services_write)

```
gyro container notifications set <ID> [flags]
```

### Options

```
      --cooldown int             Minutes between two announcements of the same service + event (5-1440, default 30)
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --email                    Account e-mail (default on)
      --events string            Comma separated: down, crashloop, replicas_low, scale_max, capacity, egress
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --resolved                 Send a 'resolved' message (default on)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --sms                      Account phone
      --telegram string          Telegram chat id (send /start to the platform bot, e.g. @gyrouptime_bot); empty = off
      --webhook-url string       https URL of a public server; JSON POST signed with X-Gyro-Signature: sha256=HMAC(secret, X-Gyro-Timestamp + '.' + body)
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

