## gyro template create

Sunucudan şablon oluştur

### Synopsis

Sunucudan şablon oluştur

Şablonun gerçekten kullandığı alan üzerinden GB başına aylık 0,05 € (ön ödemeli bakiyeden saatlik; ilk şablonla açılan şablon deposu hizmeti üzerinden). Silince ücret durur. En çok 10 şablon; yalnız sizin hesabınızda görünür. Yalnız kök disk alınır. `mode: snapshot` (varsayılan) sunucu açık kalır (geçici anlık görüntü, dosya sistemi kısa süre dondurulur); `mode: shutdown` sunucuyu kapatıp kopyalar ve yeniden açar. Şablondan açılan sunucularda sunucu adı, IP, root parolası, SSH sunucu anahtarları ve makine kimliği yenilenir; kullanıcılarınız, dosyalarınız ve ayarlarınız kalır. Yalnız Linux. Yaklaşık bir dakikada hazır.

API: POST /cloud/servers/{id}/templates (operationId createTemplate, scope cloud_write)

```
gyro template create <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --description string       description
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --mode string              mode (snapshot|shutdown)
      --name string              name
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --wait                     sunucu meşgul olmayana kadar bekle
      --wait-timeout duration    --wait için en uzun süre (default 15m0s)
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

* [gyro template](gyro_template.md)	 - Şablonlarınızı listele

