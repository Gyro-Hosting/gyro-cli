## gyro ssh-key

Hesabınızdaki SSH açık anahtarları

```
gyro ssh-key [flags]
```

### Options

```
  -h, --help   help for ssh-key
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
* [gyro ssh-key create](gyro_ssh-key_create.md)	 - SSH açık anahtarı yükle
* [gyro ssh-key delete](gyro_ssh-key_delete.md)	 - SSH anahtarını sil
* [gyro ssh-key get](gyro_ssh-key_get.md)	 - SSH anahtarını getir
* [gyro ssh-key list](gyro_ssh-key_list.md)	 - SSH anahtarlarını listele

