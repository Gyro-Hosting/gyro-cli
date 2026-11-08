## gyro service update

Hizmet Tercihlerini Değiştirme

### Synopsis

Hizmet Tercihlerini Değiştirme

Otomatik yenilemeyi ve fatura profilini değiştirir. — iki alan

API: PATCH /services/{id} (operationId coreUpdateService, scope Services/UpdateService)

```
gyro service update <ID> [flags]
```

### Options

```
      --auto-renew               Hizmet bazlı otomatik yenileme. Hesap geneli otomatik ödeme açıkken değiştirilemez.
      --billing-profile-id int   Fatura profilinin numarası. Sıfır hesabın varsayılanına döner.
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for update
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

* [gyro service](gyro_service.md)	 - Tüm ürün hatlarındaki hizmetler (genel köprü)

