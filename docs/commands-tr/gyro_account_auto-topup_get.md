## gyro account auto-topup get

Otomatik bakiye yükleme ayarı

### Synopsis

Otomatik bakiye yükleme ayarı

Bakiye `threshold` altına inince `amount` (TL, KDV hariç) kayıtlı karttan çekilir, KDV'li bakiye yükleme faturasıyla yüklenir. Saatte bir kontrol edilir; günlük/aylık sınırlar vardır.

API: GET /cloud/account/auto-topup (operationId getAutoTopup, scope billing_read)

```
gyro account auto-topup get [flags]
```

### Options

```
  -h, --help   help for get
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

* [gyro account auto-topup](gyro_account_auto-topup.md)	 - Otomatik bakiye yükleme ayarı

