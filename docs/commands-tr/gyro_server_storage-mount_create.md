## gyro server storage-mount create

Kovayı sunucuya bağla

### Synopsis

Kovayı sunucuya bağla

Nesne Depolama hizmetinizdeki bir kovayı Linux sunucunuza (systemd) klasör olarak bağlar (rclone FUSE, yerel yazma önbelleği, açılışta başlar, kopunca yeniden bağlanır). Her bağlantının yalnız o kovaya yetkili (salt okuma ya da okuma/yazma) ayrı bir erişim anahtarı vardır; anahtar sınırınızdan düşmez, bağlantıyla birlikte silinir. Gyro Bulut İstanbul (tr-ist-1) sunucularında otomatik uygulanır (`automatic: true`); diğer sunucularda `install_command` bir kez çalıştırılır. Ücretsizdir; S3 trafiği depolama hizmetinden sayılır. Yanıtta `install_command` döner (tek kullanımlık, 60 dakika). **Korumalı Depolama (2.20.0):** desteklenir; `read_only: false` ve `accept_versioning: true` verilmedikçe salt okuma bağlanır.

API: POST /cloud/servers/{id}/storage-mounts (operationId createStorageMount, scope cloud_write)

```
gyro server storage-mount create <ID> [flags]
```

### Options

```
      --accept-versioning        Protected Storage read/write only: confirms that every write/overwrite creates a new version, deletes only add a delete marker, locked files cannot be deleted and used space/cost grow. Without it: 422 versioning_ack_required.
      --bucket string            bucket name in that service
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --mount-point string       absolute path under /mnt, /srv, /data, /media, /opt, /home or /var/www; default /mnt/<bucket>. Must be empty or missing on the server.
      --read-only                Default: false on Object Storage, **true on Protected Storage**.
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --storage-service-id int   your active Object Storage service
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

