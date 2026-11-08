## gyro key

Alt hesap API anahtarları

```
gyro key [flags]
```

### Options

```
  -h, --help   help for key
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

* [gyro](gyro.md)	 - Gyro Hosting komut satırı aracı
* [gyro key create](gyro_key_create.md)	 - Alt hesap anahtarı oluştur (token bir kez döner)
* [gyro key delete](gyro_key_delete.md)	 - Alt hesap anahtarını sil
* [gyro key list](gyro_key_list.md)	 - Alt hesap anahtarlarını listele (token göstermez)

