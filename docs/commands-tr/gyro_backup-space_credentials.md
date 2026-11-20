## gyro backup-space credentials

Parolayla SFTP bağlantı bilgisi

### Synopsis

Parolayla SFTP bağlantı bilgisi

Adres, port, kullanıcı, dizin ve SFTP parolası (`Cache-Control: no-store`). Gizli değer gösterdiği için `services_write` ister. Alan kurulurken 409 `not_ready`.

API: GET /cloud/backup-spaces/{id}/credentials (operationId getBackupSpaceCredentials, scope services_write)

```
gyro backup-space credentials <ID> [flags]
```

### Options

```
  -h, --help   help for credentials
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

