## gyro server storage-mount command

Yeni kurulum komutu (bağlantı anahtarı yenilenir)

### Synopsis

Yeni kurulum komutu (bağlantı anahtarı yenilenir)

Bağlantı için yeni erişim anahtarı ve yeni tek kullanımlık komut üretir (eski anahtar hemen durur). Sunucu yeniden kurulduğunda ya da komutun süresi dolduğunda kullanın.

API: POST /cloud/servers/{id}/storage-mounts/{mid}/command (operationId renewStorageMountCommand, scope cloud_write)

```
gyro server storage-mount command <ID> <MID> [flags]
```

### Options

```
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for command
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --wait                     sunucu meşgul olmayana kadar bekle
      --wait-timeout duration    --wait için en uzun süre (default 15m0s)
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

* [gyro server storage-mount](gyro_server_storage-mount.md)	 - Sunucuya bağlı kovalar

