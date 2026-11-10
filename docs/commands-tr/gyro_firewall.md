## gyro firewall

Ağ güvenlik duvarı politikası (kurallar, varsayılan kural, durum)

```
gyro firewall [flags]
```

### Options

```
  -h, --help   help for firewall
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

* [gyro](gyro.md)	 - Gyro Hosting komut satırı aracı
* [gyro firewall confirm](gyro_firewall_confirm.md)	 - Güvenlik duvarı denemesini kalıcı yap
* [gyro firewall delete](gyro_firewall_delete.md)	 - Güvenlik duvarı kurallarınızı kaldır
* [gyro firewall get](gyro_firewall_get.md)	 - Ağ güvenlik duvarı politikası (kurallar, varsayılan kural, durum)
* [gyro firewall set](gyro_firewall_set.md)	 - Ağ güvenlik duvarı politikasını değiştir

