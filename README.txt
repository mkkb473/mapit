go run cmd/main.go -bytes-size=1000000
    -sampling-bound: should be greater than 131072 (8*(4**7))
    -bytes-size: should be greater than 4

KV is struct with []byte as Key and Iterator as Value;
    Key is parsed from randomly generated int32 with BigEndian
    Iterator is very simply constructed, probably very slow.