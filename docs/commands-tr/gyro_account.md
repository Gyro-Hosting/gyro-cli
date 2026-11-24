## gyro account

İşlem yapabildiğiniz hesaplar, hesap ayarları

```
gyro account [flags]
```

### Options

```
  -h, --help   help for account
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
* [gyro account auto-topup](gyro_account_auto-topup.md)	 - Otomatik bakiye yükleme ayarı
* [gyro account card](gyro_account_card.md)	 - Kayıtlı kartları listele
* [gyro account list](gyro_account_list.md)	 - --account ile işlem yapabildiğiniz hesaplar (kendi hesabınız + panel kullanıcısı olduğunuz hesaplar)
* [gyro account use](gyro_account_use.md)	 - Bir hesabı bu profilin varsayılanı yap (self = kendi hesabınız)

