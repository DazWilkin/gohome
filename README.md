# Google Home (hidden) API

See:
https://rithvikvibhu.github.io/GHLocalApi/

Provides a partial (!) implementation in Golang.

**NB** Not all field are implemented

Example:
```golang
client := NewClient(
  &http.Client{},
  "you-device-ip:8008")
info, err := client.DeviceInfo()
if err != nil {
  log.Fatal(err)
}

log.Printf("%v", info)
```