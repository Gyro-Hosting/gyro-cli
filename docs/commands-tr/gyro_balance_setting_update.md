## gyro balance setting update

Cüzdan Ayarlarını Değiştirme

### Synopsis

Cüzdan Ayarlarını Değiştirme

Otomatik ödemeyi ve düşük bakiye uyarısını ayarlar. — faturaları etkiler

API: PATCH /balance/settings (operationId coreUpdateBalanceSettings, scope Balance/UpdateBalanceSettings)

```
gyro balance setting update [flags]
```

### Options

```
      --alert string              Uyarı bloğu. Açık ve eşik alanlarını taşır; açarken pozitif bir eşik gerekir. (JSON)
      --alert-recipients string   Uyarıyı alacak fatura adresi numaraları. Liste yerine geçer; hesabın sahibi burada yer almaz, her zaman alır.
      --auto-pay-from-balance     Yenilemeler cüzdandan tahsil edilsin mi. Bu değiştiğinde açık yenileme faturaları yeni tercihe göre yeniden ele alınır.
      --backup-card-id int        Cüzdan yetmediğinde çekilecek kart. Sıfır göndermek kart zincirini tümüyle kaldırır.
      --data string               istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                   hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                      help for update
      --idempotency-key string    tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --set stringArray           ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro balance setting](gyro_balance_setting.md)	 - Cüzdan Ayarlarını Değiştirme

