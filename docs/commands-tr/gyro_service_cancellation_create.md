## gyro service cancellation create

İptal Talebi Açma

### Synopsis

İptal Talebi Açma

Hizmetin iptali için bir talep açar. — aynı anda tek talep

API: POST /services/{id}/cancellation (operationId coreCreateCancellation, scope Services/CreateCancellation)

```
gyro service cancellation create <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --note string              Operatöre not.
      --reason string            Gerekçe. Beş değerden biri: gerek kalmadı, pahalı, başka yere geçiliyor, özellik eksik ya da diğer.
      --reason-detail string     Serbest açıklama. Gerekçe diğer ise zorunludur.
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --urgency string           Ne zaman iptal edilsin: hemen ya da dönem sonunda.
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

* [gyro service cancellation](gyro_service_cancellation.md)	 - İptal Talebi Açma

