## gyro me update

Profili Güncelleme

### Synopsis

Profili Güncelleme

Gönderdiğiniz profil alanlarını değiştirir. — operatör kapısı

API: PATCH /me (operationId coreUpdateMe, scope Account/UpdateMe)

```
gyro me update [flags]
```

### Options

```
      --avatar string            Profil görseli. Yalnız kodlanmış içerik kabul edilir; adres reddedilir.
      --birthday string          Doğum tarihi.
      --company string           Şirket bilgileri. Hesabın kurumsal olmasını gerektirir. (JSON)
      --country string           Ülke kodu.
      --currency string          Görüntü para birimi kodu.
      --custom-fields string     Ek alanların değerleri. Kutu tipli alanlar dizi alır. (JSON)
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --date-format string       Tarih biçimi. Boş göndermek hesabın kendi seçimini siler.
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --email string             email
  -h, --help                     help for update
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --identity string          Kimlik numarası.
      --kind string              Hesap türü. Bireysele geçmek şirket bilgilerini temizler.
      --landline-phone string    Sabit telefon.
      --language string          Dil kodu. Kurulumda bulunmalıdır.
      --name string              Adı. Boşaltılamaz.
      --phone string             phone
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --surname string           Soyadı. Boşaltılamaz.
      --timezone string          Saat dilimi. Boş göndermek hesabın kendi seçimini siler.
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

* [gyro me](gyro_me.md)	 - Profili Okuma

