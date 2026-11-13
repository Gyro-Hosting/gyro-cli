## gyro server storage-mount list

Sunucuya bağlı kovalar

### Synopsis

Sunucuya bağlı kovalar

Nesne Depolama hizmetinizdeki bir kovayı Linux sunucunuza (systemd) klasör olarak bağlar (rclone FUSE, yerel yazma önbelleği, açılışta başlar, kopunca yeniden bağlanır). Her bağlantının yalnız o kovaya yetkili (salt okuma ya da okuma/yazma) ayrı bir erişim anahtarı vardır; anahtar sınırınızdan düşmez, bağlantıyla birlikte silinir. Gyro Bulut İstanbul (tr-ist-1) sunucularında otomatik uygulanır (`automatic: true`); diğer sunucularda `install_command` bir kez çalıştırılır. Ücretsizdir; S3 trafiği depolama hizmetinden sayılır. Korumalı Depolama kovaları: varsayılan salt okuma (2.20.0).

API: GET /cloud/servers/{id}/storage-mounts (operationId listStorageMounts, scope cloud_read)

```
gyro server storage-mount list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

