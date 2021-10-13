package smartconf

import (
    "testing"
    "fmt"
)

func TestSmartConf(t *testing.T) {

    zonecfg := NewDefaultZoneConfig()
    if obj, err := zonecfg.JSON(); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(string(obj))
    }
}
