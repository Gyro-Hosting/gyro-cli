## gyro api request

API adresinin altındaki bir yola ham istek gönder

```
gyro api request <METHOD> <PATH> [flags]
```

### Examples

```
  gyro api request GET /cloud/servers -p status=active
  gyro api request POST /cloud/ssh-keys --set name=ci --set public_key=@~/.ssh/id_ed25519.pub
```

### Options

```
      --all                      sayfalamayı izle (GET)
      --data string              JSON gövde, @dosya ya da @-
  -h, --help                     help for request
      --idempotency-key string   Idempotency-Key başlığı gönder
  -P, --param stringArray        sorgu parametresi anahtar=değer
      --set stringArray          gövde alanı anahtar=değer (@dosya dosyadan okur)
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

* [gyro api](gyro_api.md)	 - Herhangi bir API ucunu doğrudan çağır ('gh api' gibi)

