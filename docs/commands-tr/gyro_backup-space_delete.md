## gyro backup-space delete

Yedekleme alanını iptal et (sil)

### Synopsis

Yedekleme alanını iptal et (sil)

`confirm=true` zorunlu. `immediate=true` hemen iptal eder: SFTP erişimi durur ve alan İÇİNDEKİ TÜM DOSYALARLA yaklaşık bir dakikada kalıcı olarak silinir (snapshot yok, bekleme süresi yok; geri alınamaz). `immediate` olmadan iptal talebi açılır, işlenene kadar alan çalışır. Anında silme operatörce kapatıldıysa 403 `immediate_disabled`.

API: DELETE /cloud/backup-spaces/{id} (operationId deleteBackupSpace, scope cloud_delete)

```
gyro backup-space delete <ID> [flags]
```

### Options

```
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for delete
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --immediate                alanı ve tüm dosyaları şimdi sil (geri alınamaz)
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

* [gyro backup-space](gyro_backup-space.md)	 - Yedekleme Alanı: sunucunuza bağlı SFTP yedekleme alanı (tr-ist-1)

