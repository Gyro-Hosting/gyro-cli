## gyro me notification set

Bildirim Tercihlerini Yazma

### Synopsis

Bildirim Tercihlerini Yazma

Bildirim tercihlerinin tamamını yeniden yazar. — tam değiştirme

API: PUT /me/notifications (operationId coreUpdateNotifications, scope Account/UpdateNotifications)

```
gyro me notification set [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --domain string            Alan adı bildirimleri. (JSON)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --general string           general
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --invoices string          Fatura bildirimleri. E-posta ve kısa mesaj kanallarını taşır; eksik kanal kapalı sayılır. (JSON)
      --marketing string         Pazarlama bildirimleri. (JSON)
      --product string           Ürün ve hizmet bildirimleri. (JSON)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --support string           Destek bildirimleri. (JSON)
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

* [gyro me notification](gyro_me_notification.md)	 - Bildirim Tercihlerini Okuma

