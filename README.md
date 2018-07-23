# XML Parser

The test scenario reads from `request.xml` and pushes the content to `Parser` so to have it unmarshalled into `Menu` and `Food` (array) structs.

```
$ go test
PASS
ok  	playground/xml_parser	0.006s
```

### Reference

- [Unmarshal](https://golang.org/pkg/encoding/xml/#Unmarshal)
