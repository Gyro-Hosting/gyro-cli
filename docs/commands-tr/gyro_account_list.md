## gyro account list

--account ile işlem yapabildiğiniz hesaplar (kendi hesabınız + panel kullanıcısı olduğunuz hesaplar)

### Synopsis

Kendi hesabınızı ve sizi Alt Hesaplar altında ekleyen her hesabı oradaki izinlerinizle listeler.
Birinde işlem yapmak için --account <id>, GYRO_ACCOUNT=<id> ya da profil varsayılanı için 'gyro account use <id>'.
İzinler paneldeki izinlerin aynısıdır; yalnız hesap sahibine açık işlemler (API anahtarları, webhook, kullanıcılar, bakiye, kartlar, API etkinlik kaydı) reddedilir.
Sipariş, fatura ödemesi ve alan adı kaydı / transferi yalnız o hesabın bakiyesinden ödenir (kayıtlı kart kullanılmaz).
"Parolaları görüntüleme" izni yoksa gizli alanlar boş döner; konsol, parola sıfırlama, kurtarma, yeniden kurulum gibi araçlar reddedilir.

```
gyro account list [flags]
```

### Examples

```
  gyro account list
  gyro server list --account 12345
  gyro account use 12345
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

* [gyro account](gyro_account.md)	 - İşlem yapabildiğiniz hesaplar, hesap ayarları

