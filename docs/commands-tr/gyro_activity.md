## gyro activity

Hesap etkinlikleri (girişler ve güvenlik olayları)

```
gyro activity [flags]
```

### Options

```
      --all                tüm sayfaları getir
      --from string        Bu günden (YYYY-AA-GG, tüm gün, Türkiye saati) ya da andan (ISO 8601) itibaren, dahil
  -h, --help               help for activity
      --limit int          sayfa başına kayıt (en çok 100)
      --page int           sayfa numarası (default 1)
      --param-lang title   title dili (varsayılan: hesap dili) (tr|en)
      --to string          Bu güne (YYYY-AA-GG, tüm gün) ya da ana (ISO 8601) kadar, dahil
      --type string        login = yalnız girişler, action = yalnız hesap olayları (all|login|action)
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
* [gyro activity get](gyro_activity_get.md)	 - Hesap etkinlikleri (girişler ve güvenlik olayları)
* [gyro activity list](gyro_activity_list.md)	 - API anahtarlarınızla yapılan yazma çağrılarını listele

