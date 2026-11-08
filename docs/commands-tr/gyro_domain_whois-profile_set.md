## gyro domain whois-profile set

Profili Güncelleme

### Synopsis

Profili Güncelleme

Kayıtlı bir profili yeniden yazar. — tam kişi

API: PUT /domains/whois-profiles/{pid} (operationId coreUpdateWhoisProfile, scope Domains/UpdateWhoisProfile)

```
gyro domain whois-profile set <PID> [flags]
```

### Options

```
      --address string           Adres satırı.
      --address2 string          İkinci adres satırı. Profil bütün olarak değiştirilir; anahtar yoksa boş satır kaydedilir.
      --city string              Şehir.
      --company string           Kurum adı. Tek isteğe bağlı alandır.
      --contact string           contactTam kişi. Kayıtlı kişi bütünüyle değiştirilir. (JSON)
      --country string           Ülke kodu. Büyük harfe çevrilir.
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --email string             Geçerli bir e-posta.
      --first-name string        Kişinin adı.
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --last-name string         Kişinin soyadı.
      --name string              Yeni etiket.
      --phone string             Telefon.
      --postal-code string       postal_code
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --state string             İl ya da eyalet. Alan adı kişilerinde zorunludur, kayıtlı profillerde değildir.
      --zip string               Posta kodu.
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

* [gyro domain whois-profile](gyro_domain_whois-profile.md)	 - Kayıtlı Profilleri Listeleme

