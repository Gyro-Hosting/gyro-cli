## gyro account card list

Kayıtlı kartları listele

### Synopsis

Kayıtlı kartları listele

Yalnız maskeli kart (marka, son 4, SKT). Kartı banka saklar (VakıfBank PanCode); Gyro kart numarası tutmaz. `add_url` kartın 3D Secure ile eklendiği panel sayfasıdır.

API: GET /cloud/account/cards (operationId listSavedCards, scope billing_read)

```
gyro account card list [flags]
```

### Options

```
  -h, --help   help for list
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

* [gyro account card](gyro_account_card.md)	 - Kayıtlı kartları listele

