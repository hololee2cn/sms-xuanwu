package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIni(t *testing.T) {
	var (
		inicontext = `
appname = test
#comment
httpport =     8080
mysqlport =3306

PI = 3.1415926
ok= true
`
		keyValue = map[string]interface{}{
			"appname":   "test",
			"httpport":  8080,
			"mysqlport": int64(3306),
			"PI":        float64(3.1415926),
			"ok":        true,
		}
	)

	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))

	confDir := filepath.Join(appPath, "conf")
	err = os.MkdirAll(confDir, os.ModeDir)
	if err != nil {
		t.Fatal(err)
	}

	confPath := filepath.Join(confDir, "app.conf")

	t.Log(appPath)

	f, err := os.Create(confPath)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(confDir)

	_, err = f.WriteString(inicontext)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}

	Init()

	for k, v := range keyValue {
		switch v.(type) {
		case int:
			ret, err := Int(k)
			if err != nil || ret != v {
				t.Errorf("err: %v, k: %v, ret: %v", err, k, ret)
			}
		case int64:
			ret, err := Int64(k)
			if err != nil || ret != v {
				t.Errorf("err: %v, k: %v, ret: %v", err, k, ret)
			}
		case string:
			ret, err := String(k)
			if err != nil || ret != v {
				t.Errorf("err: %v, k: %v, ret: %v", err, k, ret)
			}
		case bool:
			ret, err := Bool(k)
			if err != nil || ret != v {
				t.Errorf("err: %v, k: %v, ret: %v", err, k, ret)
			}
		case float64:
			ret, err := Float64(k)
			if err != nil || ret != v {
				t.Errorf("err: %v, k: %v, ret: %v", err, k, ret)
			}
		}
	}
}
