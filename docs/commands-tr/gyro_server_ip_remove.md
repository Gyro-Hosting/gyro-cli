## gyro server ip remove

Ek IPv4 adresini bırak

### Synopsis

Ek IPv4 adresini bırak

**Asgari ücret:** her eklenen IP en az 6 saat ücretlendirilir (`pricing.minimum_hours`): daha erken bırakılsa da asgari süre (bırakırken) düşülür; aynı adresi yeniden almak yeni bir asgari süre başlatır. Bir sunucuya 60 dakikada en çok `adds_per_hour` (varsayılan 10) adres eklenebilir. Adres sunucudan kaldırılır ve havuza döner (aynı adresi geri alma garantisi yok); ücret bırakınca durur (asgari süre yine uygulanır). `confirm=true` gerekir.

API: DELETE /cloud/servers/{id}/ips/{ip} (operationId removeServerIp, scope cloud_write)

```
gyro server ip remove <ID> <IP> [flags]
```

### Options

```
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for remove
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
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

