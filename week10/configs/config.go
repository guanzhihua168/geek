package configs

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	App *viper.Viper // 配置
)

//读取文件数据
func bindataRead(data []byte, name string) ([]byte, error) {

	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

//读取文件信息
type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configsConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\xd1\x72\xda\x30\x10\x7c\xf7\x57\xdc\x4c\x9f\x63\x04\x34\x40\xf5\xd4\x84\x90\x09\x9d\xd0\x32\xb5\x33\x79\xec\x9c\xed\xc3\x36\x23\x5b\x42\x3a\x83\xc9\xd7\x77\x64\xd3\x38\x4c\xfa\xa6\xd9\x5d\xed\x6a\x4f\x17\x91\x3d\x92\x95\x01\xc0\xef\xa6\xde\xe8\x8c\x24\x64\x94\x34\x79\x00\xf0\xc4\x6c\xb6\xda\xb2\x84\x85\x10\xc2\x2b\x08\xb3\xb8\xac\x48\x37\x2c\x61\xe6\x91\x57\x5b\x32\x7d\x84\xee\x8c\xf1\x5e\x0f\xb4\xc3\x46\xf1\x16\x73\x8a\xca\x37\x92\x30\xf6\xea\x0d\xb6\x1f\x11\x31\x08\x97\xba\x66\x6a\xf9\xda\xfc\x59\xe7\x11\x1e\x69\x8b\x5c\x48\x70\xac\x2d\xe6\x34\x52\x3a\x77\x3d\xf7\x58\x2a\xfa\x89\x15\x49\x40\x63\x06\x68\xd5\xb2\x84\x50\x69\xdf\xe0\xc5\x28\x8d\xd9\x67\x93\xa6\xc3\xdd\xa0\xe8\x86\xf0\x62\x95\x84\x82\xd9\xc8\xd1\x68\x3c\x99\x87\x22\x14\xe1\x58\xfa\xee\x23\xc7\xc8\x65\xfa\xae\x5f\x57\x98\xd3\x06\xdb\xbe\xc9\x2d\xc0\x17\xd8\xdc\x5f\xb3\x77\x4a\xe9\xd3\xaa\x65\xe7\xc7\x01\x70\x03\xe1\xde\xe4\xc3\x91\xde\xcf\xa6\xce\x83\x55\x85\xa5\xf2\xc2\x27\xed\x58\x82\xab\xd8\x84\x87\x43\x98\xea\x2a\x00\xe8\xbf\xe0\xeb\xec\xd6\x07\x38\xb2\x7d\xe7\xe9\x78\x3a\x5b\xcc\xbf\x2d\x26\xdf\x07\x21\x3a\x77\xd2\x36\x93\x50\x1c\xb5\xcd\x4e\xb4\x57\x6f\x26\x49\x8a\x5d\x11\x00\xac\x5d\x14\x3d\x4b\x60\xdb\x50\x00\xf0\x68\x75\xf5\x5f\x8f\x58\xff\x7b\xef\x27\xf2\xc7\x6b\xec\xb9\x88\x52\x4b\x2c\x81\xb2\xec\x9c\xee\xcf\x9d\xb5\x6b\xc8\x4a\x48\x94\xce\x6f\x1c\xd9\x63\x99\xfa\x8c\x55\x6b\x4a\x4b\x12\xe6\x13\x21\x82\x07\x64\x4c\xd0\x51\xb7\x1c\xf7\xf1\xd9\x90\x84\xea\xec\x0e\xea\xaa\x94\xd5\x9a\xaf\x8a\x78\xe0\x02\xf6\xa3\x19\xfe\x65\x3a\x15\xb3\xce\xac\xbf\xea\xc3\xff\x0c\xe1\x31\x26\x8a\xb6\x96\x76\x65\x7b\xe1\x02\x80\x65\x81\xd6\xf9\xb7\x37\xbc\x5b\x74\x39\xd6\x75\xdb\x2b\x21\xee\xe7\xb2\xc1\x76\x9d\x29\x5a\xea\xba\x76\xc3\xd2\xfe\x32\x54\x5f\xa0\xa9\xf8\x1b\x00\x00\xff\xff\x9a\x37\x5d\xef\x31\x03\x00\x00")

//读取Yaml的压缩过的文件
func configsConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsConfigYaml,
		"configs/config.yaml",
	)
}

//解析yaml文件
func configsConfigYaml() (*asset, error) {
	bytes, err := configsConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/config.yaml", size: 817, mode: os.FileMode(420), modTime: time.Unix(1576903324, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"configs/config.yaml": configsConfigYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"configs": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{configsConfigYaml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
