## gyro container notifications get

Bildirim ayarları ve süren uyarılar

### Synopsis

Bildirim ayarları ve süren uyarılar

`settings` (kanallar, olaylar, tekrar aralığı; webhook anahtarı dönmez), `conditions` = şu an süren sorunlar, `history` = son 30 uyarı. Olaylar: down, crashloop, replicas_low, scale_max, capacity, egress (çıkış IP'leri değişti). Olay ~2 dakika sürünce bildirilir; aynı servis + olay tekrar aralığında bir kez (gidip gelenler dahil), düzelince 'düzeldi' mesajı.

API: GET /cloud/containers/{id}/notifications (operationId typed_get_containers__id__notifications, scope services_read)

```
gyro container notifications get <ID> [flags]
```

### Options

```
  -h, --help   help for get
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

