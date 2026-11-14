## gyro limits

Hesap sınırları ve kullanım; nasıl artırılır (destek talebi)

### Synopsis

Hesap sınırları: hesabınızın tüm sınırları ve şu anki kullanım (ön ödemeli bulutta yeni hesap sınırı, Gyro Bulut İstanbul
sunucuları, şablonlar, sunucu başına ek IP ve disk, API sipariş sınırları, anahtarlar, webhook'lar, güvenlik duvarı
kuralları ve grupları, depolama kovaları ve anahtarları, DNS zone'ları, kapsayıcı paketleri, Gyro Pipeline,
uptime izleyicileri, projeler).

Sınırlar talep üzerine artırılır: destek talebi açın.
  gyro ticket create --department-id 3 --subject "Limit increase: <anahtar>" --message "İstenen sınır: <N>. Gerekçe: ..."
Paket sınırları (kova, zone, kapsayıcı servisi ...) paketi yükselterek ya da ek hizmetle de büyür.
Bir komut sınır yüzünden reddedilirse hata çıktısı sınırı, kullanımı ve hazır talep komutunu gösterir.

```
gyro limits [flags]
```

### Examples

```
  gyro limits
  gyro limits -o json --query "limits[?key=='compute.max_vms']"
  gyro ticket create --department-id 3 --subject "Limit increase: new_account.servers" --message "Requested limit: 10. Reason: ..."
```

### Options

```
      --fresh   yeniden oku (yanıt 60 sn önbellekli)
  -h, --help    help for limits
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

