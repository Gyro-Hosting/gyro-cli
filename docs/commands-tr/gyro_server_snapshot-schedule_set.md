## gyro server snapshot-schedule set

Zamanlanmış snapshot planını oluştur ya da değiştir

### Synopsis

Zamanlanmış snapshot planını oluştur ya da değiştir

Gönderilmeyen alan değişmez (yeni plan her gün 03:00, keep 7, açık ile başlar). `keep` 1 ile sunucunun snapshot sınırı arası; `weekday` yalnız haftalıkta. Yalnız Gyro Bulut İstanbul (tr-ist-1) sunucuları (diğerleri: 422 `not_supported`). Her gün ya da her hafta `hour`:`minute` Türkiye saatinde `auto-YYYYMMDD-HHMM` adlı snapshot alınır; öncesinde en eski otomatik snapshot'lar silinerek en çok `keep` tanesi kalır. Elle aldığınız snapshot'lar asla silinmez ama sunucunun snapshot sınırına (`snapshot_limit`, elle + otomatik) sayılır: sınır doluysa çalışma atlanır (`last_status` skipped, `last_error` snapshot_limit). Ön ödemeli sunucuda ücretsiz hakkın üstündeki snapshot'lar elle alınanlar gibi saatlik ücretlenir. Sunucuda snapshot varken ek disk büyütülemez (409 `has_snapshots`): önce planı kapatıp snapshot'ları silin.

API: PUT /cloud/servers/{id}/snapshot-schedule (operationId setSnapshotSchedule, scope cloud_write)

```
gyro server snapshot-schedule set <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --enabled                  enabled
      --frequency string         frequency (daily|weekly)
  -h, --help                     help for set
      --hour int                 Türkiye time
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --keep int                 <= snapshot_limit
      --minute int               minute
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --wait                     sunucu meşgul olmayana kadar bekle
      --wait-timeout duration    --wait için en uzun süre (default 15m0s)
      --weekday string           weekly only (0-6 also accepted, 0 = Sunday) (sunday|monday|tuesday|wednesday|thursday|friday|saturday)
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

* [gyro server snapshot-schedule](gyro_server_snapshot-schedule.md)	 - Sunucunun zamanlanmış snapshot planı

