## gyro domain auth-code set

Gelen Transfer Kodunu Kaydetme

### Synopsis

Gelen Transfer Kodunu Kaydetme

Devam eden bir transfer için kodu günceller. — anahtarın sahibi

API: PUT /domains/{domain}/auth-code (operationId coreSaveAuthCode, scope Domains/SaveAuthCode)

```
gyro domain auth-code set <DOMAIN> [flags]
```

### Options

```
      --code string              Yeni transfer kodu. İşaretleme ayıklanır ve boşluklar kırpılır.
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
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

* [gyro domain auth-code](gyro_domain_auth-code.md)	 - Transfer Kodunu İsteme

