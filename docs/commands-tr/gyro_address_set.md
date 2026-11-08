## gyro address set

Adresi Güncelleme

### Synopsis

Adresi Güncelleme

Adresin tamamını yeniden yazar. — tam gövde

API: PUT /addresses/{id} (operationId coreUpdateAddress, scope Addresses/UpdateAddress)

```
gyro address set <ID> [flags]
```

### Options

```
      --address string           Açık adres.
      --city string              Şehir. Listeden numara, liste boşsa serbest metin.
      --company string           Şirket bilgileri. Kurumsal kişide ad zorunludur. (JSON)
      --country string           Ülke kodu. Başvuru ucundan alınır.
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --email string             Geçerli bir e-posta.
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --identity string          Kimlik numarası.
      --is-default               Bunu varsayılan profil yap.
      --kind string              Kişi türü. Öntanımlı olarak bireysel.
      --label string             Serbest etiket.
      --name string              İlgili kişinin adı.
      --notifications string     Kanal matrisi. Tam değiştirmedir; hiç gönderilmezse hepsi açılır. (JSON)
      --overwrite-invoices       Açık faturaları bu adrese yeniden yaz.
      --phone string             Kişinin telefonu.
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --state string             İl. Listeden numara, liste boşsa serbest metin.
      --surname string           Soyadı.
      --zipcode string           Posta kodu. En çok yirmi karakter.
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

* [gyro address](gyro_address.md)	 - Fatura adresleri

