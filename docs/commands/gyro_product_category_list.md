## gyro product category list

List product categories

### Synopsis

List product categories

List product categories. Core WISECP client endpoint (Turkish reference: https://dev.wisecp.com/tr/satistaki-urunler).

API: GET /products/categories (operationId coreGetProductCategories, scope Products/GetProductCategories)

```
gyro product category list [flags]
```

### Options

```
  -h, --help          help for list
      --type string   Yalnız bu türün kategorileri. Bilinmeyen değer boş liste verir, hata değil.
```

### Options inherited from parent commands

```
      --api-url string     API base URL (env GYRO_API_URL)
      --debug              log HTTP requests to stderr (secrets masked)
      --lang string        message language: tr or en (default from locale)
      --no-color           disable colours (also NO_COLOR)
      --no-header          omit the table header
  -o, --output string      output format: table, json, yaml, text
  -p, --profile string     profile to use (env GYRO_PROFILE)
      --query string       JMESPath query applied to the result
  -q, --quiet              print only ids
      --timeout duration   overall timeout, e.g. 30s, 10m
      --wide               show all table columns
  -y, --yes                do not ask for confirmation
```

### SEE ALSO

* [gyro product category](gyro_product_category.md)	 - List product categories

