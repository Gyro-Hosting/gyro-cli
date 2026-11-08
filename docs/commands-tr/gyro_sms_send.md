## gyro sms send

Toplu Gönderim

### Synopsis

Toplu Gönderim

**Para harcar.** Mesajı gönderir ve tutarı cüzdandan keser. — cüzdandan keser

API: POST /sms/send (operationId coreSendSms, scope Sms/SendSms)

```
gyro sms send [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for send
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --message string           Mesaj metni. Parça tavanını aşan metin kırpılır ve kırpıldığı bildirilir.
      --numbers strings          Alıcı numaralar. Uluslararası biçimde; tek bir metin olarak gönderilirse satır, virgül ve noktalı virgülden bölünür.
      --sender string            Gönderen başlığı. Hesabın etkin başlıklarından biri olmalıdır.
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro sms](gyro_sms.md)	 - SMS mesajları

