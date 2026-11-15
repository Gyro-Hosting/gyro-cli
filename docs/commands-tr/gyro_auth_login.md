## gyro auth login

Tarayıcıyla giriş (bu profil için süreli bir API anahtarı oluşturur)

### Synopsis

Gyro panelini tarayıcıda açar; isteği orada onaylarsınız (izinler ve süre o sayfada seçilir) ve yeni anahtar
'gyro configure' ile aynı şekilde profile kaydedilir. Elle anahtar kopyalanmaz.
SSH bağlantısında ya da ekransız ortamda adres yalnız yazdırılır: giriş yapmış olduğunuz herhangi bir cihazda açın.
CI'da anahtar kullanmaya devam edin: echo "$ANAHTAR" | gyro configure --token-stdin

```
gyro auth login [flags]
```

### Examples

```
  gyro auth login
  gyro auth login --profile work --no-browser
  gyro auth login --template readonly
```

### Options

```
  -h, --help              help for login
      --keyring           anahtarı işletim sistemi anahtar zincirinde sakla
      --no-browser        yalnız adresi yaz, tarayıcı açma
      --scopes string     kapsamları önceden seç (virgülle, ör. cloud_read,services_read); tarayıcıda yine değiştirebilirsiniz
      --template string   izin şablonunu önceden seç: readonly, dns, cloud, containers, storage, billing, full
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

* [gyro auth](gyro_auth.md)	 - Tarayıcıyla giriş yap, CLI oturumunu göster ya da kapat

