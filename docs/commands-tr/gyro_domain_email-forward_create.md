## gyro domain email-forward create

E-posta Yönlendirmesi Ekleme

### Synopsis

E-posta Yönlendirmesi Ekleme

Alan adına gelen postayı başka bir kutuya taşıyan bir kural ekler. — ek hizmet gerekir

API: POST /domains/{domain}/email-forwards (operationId coreCreateEmailForward, scope Domains/CreateEmailForward)

```
gyro domain email-forward create <DOMAIN> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --prefix string            Alan adındaki yerel kısım. Yönlendirme bu kısım ile alan adının birleşimine yanıt verir.
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --target string            Hedef posta kutusu. Geçerli bir e-posta adresi olmalıdır.
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

* [gyro domain email-forward](gyro_domain_email-forward.md)	 - E-posta Yönlendirmelerini Listeleme

