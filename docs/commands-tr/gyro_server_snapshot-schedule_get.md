## gyro server snapshot-schedule get

Sunucunun zamanlanmış snapshot planı

### Synopsis

Sunucunun zamanlanmış snapshot planı

Yalnız Gyro Bulut İstanbul (tr-ist-1) sunucuları (diğerleri: 422 `not_supported`). Her gün ya da her hafta `hour`:`minute` Türkiye saatinde `auto-YYYYMMDD-HHMM` adlı snapshot alınır; öncesinde en eski otomatik snapshot'lar silinerek en çok `keep` tanesi kalır. Elle aldığınız snapshot'lar asla silinmez ama sunucunun snapshot sınırına (`snapshot_limit`, elle + otomatik) sayılır: sınır doluysa çalışma atlanır (`last_status` skipped, `last_error` snapshot_limit). Ön ödemeli sunucuda ücretsiz hakkın üstündeki snapshot'lar elle alınanlar gibi saatlik ücretlenir. Sunucuda snapshot varken ek disk büyütülemez (409 `has_snapshots`): önce planı kapatıp snapshot'ları silin.

API: GET /cloud/servers/{id}/snapshot-schedule (operationId getSnapshotSchedule, scope cloud_read)

```
gyro server snapshot-schedule get <ID> [flags]
```

### Options

```
  -h, --help   help for get
```

### Options inherited from parent commands

```
      --account gyro account list   panel kullanıcısı olduğunuz başka bir hesapta işlem yap (gyro account list id'si; "self" = kendi hesabınız; ortam GYRO_ACCOUNT, profil varsayılanı: `gyro account use`)
      --api-url string              API taban adresi (ortam GYRO_API_URL)
      --debug                       HTTP isteklerini stderr'e yaz (sırlar maskeli)
      --lang string                 mesaj dili: tr ya da en (varsayılan yerel ayar)
      --no-color                    renkleri kapat (NO_COLOR da olur)
      --no-header                   tablo başlığını yazma
  -o, --output string               çıktı biçimi: table, json, yaml, text
  -p, --profile string              kullanılacak profil (ortam GYRO_PROFILE)
      --query string                sonuca uygulanacak JMESPath sorgusu
  -q, --quiet                       yalnız kimlikleri yaz
      --timeout duration            toplam zaman aşımı, ör. 30s, 10m
      --wide                        tüm tablo sütunlarını göster
  -y, --yes                         onay sorma
```

### SEE ALSO

* [gyro server snapshot-schedule](gyro_server_snapshot-schedule.md)	 - Sunucunun zamanlanmış snapshot planı

