## gyro configure

API anahtarını ve tercihleri adlı bir profile kaydet

### Synopsis

'aws configure' gibi etkileşimli kurulum. Profiller ~/.config/gyro/config.toml dosyasındadır (izin 0600).
Birden çok hesap için --profile AD; --keyring anahtarı dosya yerine işletim sistemi anahtar zincirinde tutar.
Etkileşimsiz: echo "$ANAHTAR" | gyro configure --token-stdin --profile ci

```
gyro configure [flags]
```

### Options

```
  -h, --help                 help for configure
      --keyring              anahtarı işletim sistemi anahtar zincirinde sakla
      --no-verify            anahtarı /whoami ile doğrulama
      --set-api-url string   profile yazılacak API adresi
      --set-lang string      kaydedilecek dil (tr/en)
      --set-output string    kaydedilecek varsayılan çıktı
      --token-stdin          anahtarı stdin'den oku (etkileşimsiz)
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
* [gyro configure delete](gyro_configure_delete.md)	 - Profili sil
* [gyro configure get](gyro_configure_get.md)	 - Profil ayarını yaz (api_url, output, lang, ssh_user, ssh_identity)
* [gyro configure list](gyro_configure_list.md)	 - Profilleri listele
* [gyro configure set](gyro_configure_set.md)	 - Profil ayarını değiştir (API anahtarı için 'gyro configure --token-stdin')

