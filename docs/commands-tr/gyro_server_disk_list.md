## gyro server disk list

Sunucunun ek diskleri

### Synopsis

Sunucunun ek diskleri

İki sınıf: `ssd` (hızlı) ve `hdd` (yedek, arşiv, medya için büyük depolama; veritabanı için kullanmayın). Ek diskler GB başına ön ödemeli bakiyeden saatlik ücretlendirilir (`classes.<sınıf>.per_gb_monthly_eur`, aylık / 672); eklendiği saatten silindiği saate kadar, her disk en az `pricing.minimum_hours` saat. Sunucu başına en çok 4 disk ve sınıf başına `classes.<sınıf>.max_gb` GB; boyutlar `step_gb` katları. Yeni disk boştur (sunucu içinde bölümlendirip biçimlendirin ve bağlayın). Yeniden kurulum ek diskleri siler.

API: GET /cloud/servers/{id}/disks (operationId listServerDisks, scope cloud_read)

```
gyro server disk list <ID> [flags]
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

* [gyro server disk](gyro_server_disk.md)	 - Sunucunun ek diskleri

