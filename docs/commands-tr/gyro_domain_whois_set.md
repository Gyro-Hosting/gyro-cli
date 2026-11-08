## gyro domain whois set

Alan Adı Kişilerini Yazma

### Synopsis

Alan Adı Kişilerini Yazma

Bir kişiyi seçilen role ya da dört role birden yazar. — sicile yazar

API: PUT /domains/{domain}/whois (operationId coreUpdateWhoisContacts, scope Domains/UpdateWhoisContacts)

```
gyro domain whois set <DOMAIN> [flags]
```

### Options

```
      --address string           Adres satırı.
      --address2 string          İkinci adres satırı (daire, kat, blok). Anahtar gönderilmezse sağlayıcıdaki satır korunur; boş metin satırı siler.
      --city string              Şehir.
      --company string           Kurum adı. Tek isteğe bağlı alandır.
      --contact string           contactYazılacak kişi. (JSON)
      --country string           Ülke kodu. Büyük harfe çevrilir.
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --email string             Geçerli bir e-posta.
      --first-name string        Kişinin adı.
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --last-name string         Kişinin soyadı.
      --phone string             Telefon.
      --postal-code string       postal_code
      --role string              Hangi rollere yazılacağı. Verilmezse dört rolün hepsine yazılır.
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

* [gyro domain whois](gyro_domain_whois.md)	 - Alan Adı Kişilerini Okuma

