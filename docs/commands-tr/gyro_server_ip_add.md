## gyro server ip add

Ek IPv4 adresi ekle

### Synopsis

Ek IPv4 adresi ekle

**Asgari ücret:** her eklenen IP en az 6 saat ücretlendirilir (`pricing.minimum_hours`): daha erken bırakılsa da asgari süre (bırakırken) düşülür; aynı adresi yeniden almak yeni bir asgari süre başlatır. Bir sunucuya 60 dakikada en çok `adds_per_hour` (varsayılan 10) adres eklenebilir. Ek IPv4 adresleri adres başına ön ödemeli bakiyeden saatlik ücretlendirilir (varsayılan adres başına aylık 2 €, güncel fiyat `pricing` alanında); eklendiği saatten çıkarıldığı saate kadar. Sunucu başına en çok `max` (varsayılan 5). Adres sunucuya otomatik tanımlanır (birincil arayüze /32, yeniden başlatma ve yeniden kurulumda korunur); ters DNS: PUT /cloud/servers/{id}/rdns. Yalnız ön ödemeli tr-ist-1 sunucuları; askıdayken eklenemez.

API: POST /cloud/servers/{id}/ips (operationId addServerIps, scope cloud_write)

```
gyro server ip add <ID> [flags]
```

### Options

```
      --count int                count
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for add
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
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

* [gyro server ip](gyro_server_ip.md)	 - Sunucunun genel IPv4 adresini (ya da --ipv6) yaz

