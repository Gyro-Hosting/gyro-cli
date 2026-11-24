## gyro backup-space

Yedekleme Alanı: sunucunuza bağlı SFTP yedekleme alanı (tr-ist-1)

```
gyro backup-space [flags]
```

### Options

```
  -h, --help   help for backup-space
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

* [gyro](gyro.md)	 - Gyro Hosting komut satırı aracı
* [gyro backup-space attach-server](gyro_backup-space_attach-server.md)	 - Alanı başka sunucuya bağla
* [gyro backup-space create](gyro_backup-space_create.md)	 - Yedekleme alanı sipariş et
* [gyro backup-space credentials](gyro_backup-space_credentials.md)	 - Parolayla SFTP bağlantı bilgisi
* [gyro backup-space delete](gyro_backup-space_delete.md)	 - Yedekleme alanını iptal et (sil)
* [gyro backup-space get](gyro_backup-space_get.md)	 - Yedekleme alanı ayrıntısı
* [gyro backup-space list](gyro_backup-space_list.md)	 - Yedekleme alanlarını listele
* [gyro backup-space plans](gyro_backup-space_plans.md)	 - Yedekleme alanı paketlerini listele
* [gyro backup-space reset-password](gyro_backup-space_reset-password.md)	 - SFTP parolasını sıfırla

