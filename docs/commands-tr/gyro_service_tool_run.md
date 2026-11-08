## gyro service tool run

Bir araç eylemi çalıştır (posta kutusu oluştur, snapshot'tan dön, ...)

### Synopsis

Bir araç eylemi çalıştır (posta kutusu oluştur, snapshot'tan dön, ...)

Gövde alanları aracın kendi alanlarıdır (araç verisine ve rehberdeki modül tablosuna bakın). Keşif yanıtında `confirm` işaretli eylemler `"confirm": true` ister.

API: POST /cloud/services/{id}/tools/{tool}/{action} (operationId runServiceToolAction, scope services_write)

```
gyro service tool run <ID> <TOOL> <ACTION> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for run
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

* [gyro service tool](gyro_service_tool.md)	 - Hizmetin araçlarını, eylemlerini ve güvenlik politikasını keşfet

