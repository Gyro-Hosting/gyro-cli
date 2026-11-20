## gyro backup-space plans

Yedekleme alanı paketlerini listele

### Synopsis

Yedekleme alanı paketlerini listele

Sipariş edilebilir boyutlar (50 GB - 1 TB), dönem başına fiyat (sabit EUR, KDV hariç; `list_amount` = üstü çizili liste fiyatı) ve `servers`: alanın bağlanabileceği hesap sunucuları. `plans` boşsa bu hesap henüz sipariş veremez. `available: false` = şu an kapasite yok.

API: GET /cloud/backup-spaces/plans (operationId listBackupSpacePlans, scope cloud_read)

```
gyro backup-space plans [flags]
```

### Options

```
  -h, --help   help for plans
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

