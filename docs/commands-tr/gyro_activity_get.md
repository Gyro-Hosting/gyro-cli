## gyro activity get

Hesap etkinlikleri (girişler ve güvenlik olayları)

### Synopsis

Hesap etkinlikleri (girişler ve güvenlik olayları)

Müşteri panelinde **Hesabım > Etkinlikler** (`/hesabim/bilgilerim#activity`) kayıtları: IP, konum, tarayıcı ve platformla girişler ve hesap olayları (parola, iki adımlı doğrulama, geçiş anahtarı, oturum ve güvenilir cihaz değişiklikleri, profil ve adres değişiklikleri ...). En yeni önce. Yalnız kendi hesabınız: hesabı anahtar belirler, hesap parametresi yoktur, alt hesap anahtarları 403 alır. Gizlilik panelle aynı: kendi girişlerinizin tam IP'si; ham tarayıcı kimliği ve olayın saklanan verisi dönmez. `account_read` izni gerekir (anahtar ekranında Hesap / okuma; salt okuma şablonunda var). API anahtarlarınızın yazma kaydı `GET /cloud/activity`'de kalır.

API: GET /activity (operationId getAccountActivity, scope account_read)

```
gyro activity get [flags]
```

### Options

```
      --all                tüm sayfaları getir
      --from string        Bu günden (YYYY-AA-GG, tüm gün, Türkiye saati) ya da andan (ISO 8601) itibaren, dahil
  -h, --help               help for get
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

* [gyro activity](gyro_activity.md)	 - Hesap etkinlikleri (girişler ve güvenlik olayları)

