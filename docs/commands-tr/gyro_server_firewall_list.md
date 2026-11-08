## gyro server firewall list

Gelen trafik güvenlik duvarı kuralları

### Synopsis

Gelen trafik güvenlik duvarı kuralları

v1'de salt okunur. Boş liste Gyro güvenlik duvarı uygulanmadığı (tüm gelen trafiğin sunucuya ulaştığı) anlamına gelir.

API: GET /cloud/servers/{id}/firewall (operationId getFirewall, scope cloud_read)

```
gyro server firewall list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

* [gyro server firewall](gyro_server_firewall.md)	 - Gelen trafik güvenlik duvarı kuralları

