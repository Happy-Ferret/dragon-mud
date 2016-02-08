// Code generated by go-bindata.
// sources:
// assets/assets_suite_test.go
// assets/existence_test.go
// assets/raw/Gamefile.toml
// DO NOT EDIT!

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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

var _assetsAssets_suite_testGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8d\xb1\x0e\xc2\x30\x0c\x44\xe7\xfa\x2b\x2c\x4f\x0d\x42\xed\x37\xb0\x20\x66\xda\x1d\x85\x60\x8c\xd5\x36\xa9\x1a\x67\x42\xfc\x3b\x2d\x30\xb2\x9d\xee\xe9\xdd\xcd\x3e\x0c\x5e\x18\x7d\xce\x6c\xf9\x62\x9c\x0d\x40\xa7\x39\x2d\x86\x35\x54\x0d\x92\xa8\x3d\xca\xb5\x09\x69\x6a\x53\xcc\xda\x8a\xc6\x41\x12\xfd\x67\x69\x62\xf1\x04\x50\xd1\xb6\xa4\x51\x08\x1c\xc0\xbd\xc4\x80\xfd\x5a\x1c\x3e\x2f\xb5\xe1\xee\x87\x9b\xde\xe1\x13\xaa\x33\x8b\x66\xe3\xe5\xe8\x75\x3c\xf9\x78\x1b\x79\xa9\xb7\xec\x56\x54\x62\x37\x73\x58\xa5\x3d\xd2\xd7\xc7\xae\xa8\x31\x39\x78\xc1\x3b\x00\x00\xff\xff\xcc\xb6\xb6\x27\xbf\x00\x00\x00")

func assetsAssets_suite_testGoBytes() ([]byte, error) {
	return bindataRead(
		_assetsAssets_suite_testGo,
		"assets/assets_suite_test.go",
	)
}

func assetsAssets_suite_testGo() (*asset, error) {
	bytes, err := assetsAssets_suite_testGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/assets_suite_test.go", size: 191, mode: os.FileMode(420), modTime: time.Unix(1453873521, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsExistence_testGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8f\x41\x6a\x2c\x21\x10\x86\xd7\x7a\x8a\xc2\x95\xc2\xbc\xee\xfd\x83\x59\x24\x24\x8b\x6c\xb2\xc9\x01\x1a\xdb\xae\x76\xa4\xd5\x1a\xb4\x0c\x81\x90\x03\xe5\x32\x39\x53\xec\xcc\x22\x4c\x18\xd0\xda\xfc\x5f\x7d\xc5\x7f\xb6\x6e\xb3\x1e\xc1\xd6\x8a\x5c\x27\xc6\xca\x52\x86\x74\xa6\xc2\xa0\xa5\x18\x40\xf9\xc0\xa7\x36\x0f\x8e\xd2\x38\xcf\xcd\x6d\xe3\x52\xac\xa7\xfc\x2f\xb5\x65\xbc\x6c\x29\xf9\x17\xa4\x5c\xc3\xe8\x43\xde\x3c\xa9\xdb\x19\x25\xf4\x56\x49\x23\xe5\xab\x2d\x30\xc1\x11\x1e\xb0\xba\x12\x66\xd4\xea\xf1\x2d\x54\xc6\xec\x50\x1d\x60\x6d\xd9\x69\x03\xef\x52\xfc\xe6\x0b\xae\xb6\x45\x9e\x1c\xe5\x35\xf8\x81\x29\xc5\x2b\x52\x3c\xb1\x56\xf5\x44\x2d\x2e\x80\xbb\xab\xb7\x03\x9b\x2f\x1d\x61\x0d\xf1\x5a\x2c\xc4\x74\x00\x2c\x05\xfe\x1f\xe1\x6e\x47\x6e\x5f\x30\x3b\xf9\xf5\xa9\x3b\x69\x86\x97\x1f\xbb\xbe\xc7\xe7\x10\xb5\xd9\xa3\x8f\x3e\xfa\xef\xef\x3b\x00\x00\xff\xff\xa5\x9d\xe8\x71\x53\x01\x00\x00")

func assetsExistence_testGoBytes() ([]byte, error) {
	return bindataRead(
		_assetsExistence_testGo,
		"assets/existence_test.go",
	)
}

func assetsExistence_testGo() (*asset, error) {
	bytes, err := assetsExistence_testGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/existence_test.go", size: 339, mode: os.FileMode(420), modTime: time.Unix(1453873521, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsRawGamefileToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x54\x4f\x6f\xdb\x3e\x0c\xbd\xe7\x53\x10\xe9\xc1\x97\x1f\x8c\xfe\xd0\x4b\x2f\x3d\xad\xdd\x2e\xdd\x86\xb5\xc3\x2e\x45\x31\xd0\x16\x63\x6b\x95\x25\x4f\x94\x93\x05\x45\xbf\xfb\x48\x49\x49\xd3\x61\x3b\x24\x30\xff\x3d\x3e\x3d\x52\x3a\x83\xeb\x88\x43\xf0\xd7\x94\xd0\x3a\x86\x0e\x99\x60\x63\x1d\xad\x56\x67\xf0\x01\x27\x02\x4f\x69\x17\xe2\x13\xf4\xc1\x6f\xec\xb0\x44\x4c\x36\xf8\x76\xf5\x20\xfe\x47\x4d\x7a\x37\xa2\x1f\x08\xd2\x48\x30\x87\x98\xf2\xc7\xa0\x85\x71\xf1\x0c\xc1\xaf\xd4\xf8\x9e\x43\x57\x70\x79\x7e\x79\xfe\x8f\x1a\xa6\xb8\xa5\x08\x73\xb4\x5b\x4c\xd5\x5f\x21\x5a\xf8\x3a\x5a\x7e\x1b\xb2\x1a\x70\x7b\x01\x13\x62\x9e\xfa\x84\x9d\x13\xe6\x31\x4c\x19\x6d\xc2\x7e\xb4\x9e\x60\x0c\x9c\xac\x1f\x4e\x3b\xa0\x37\x5a\xad\x1e\x45\x80\x1d\xee\x21\x05\x89\x7a\x23\x68\xcd\xfd\xcd\xdd\xb7\x9b\xbb\x06\x26\x62\xc6\x81\x58\x63\xe8\x9c\xd8\x13\x75\x14\xa5\xed\xe6\x78\xc6\x16\x3e\xcb\x67\x14\x1b\xbd\xfc\x59\xfe\x0f\x6c\x6a\x0a\x36\x4b\x5c\xf0\x90\x0f\x04\x33\x8d\x50\x62\x49\x48\x60\x34\xf9\x28\xed\xaa\x1e\xec\x44\xa4\xff\x55\x59\x17\x06\xad\x95\xc1\x88\x0a\x4c\x49\x11\xe4\x63\xa6\xde\x6e\x6c\x7f\xc0\x92\xac\x41\x28\x6c\x42\x2c\x7a\xc6\xf0\x43\x9a\x01\x2f\xfd\xa8\xcd\x27\xfc\x65\xa7\x65\xaa\x70\x8e\xb6\xe4\xb2\x02\x61\x49\xf3\x22\xca\x63\x1c\x28\xb1\xcc\x53\xc2\x32\x4f\xa8\x29\x57\xb0\x36\xd4\x2d\xc3\x5a\x5d\xb2\x24\xb4\x51\x35\x15\xa2\x56\x64\xf7\xad\xda\xfb\x99\x94\x26\x5b\x4e\xaa\x4d\x13\xb8\xc9\x1d\x1a\x5d\xa3\x46\x67\x47\xc5\xa9\x99\xa2\x86\x48\x35\xe2\x96\x00\x2b\x94\x14\x65\xb0\x86\x93\x11\x56\xb5\x58\x0c\x8a\xb1\x81\xdd\x28\x28\xf9\x64\xc7\xec\x6c\x15\xf4\x8a\xc9\x63\x58\x9c\x81\x8e\x80\xac\x0e\x24\xe3\x21\x44\x72\xb2\xad\xd2\x6a\xc6\x34\x9e\xe8\x95\x17\x1c\x44\x30\xec\x38\xb8\x25\x95\x84\x16\x3e\x2e\x2e\xd9\xd9\xbd\x39\x27\x60\xa4\x0c\xc7\xcb\xac\xf3\x21\xd3\x8a\xf9\xa0\x72\xb5\x35\xe5\x31\xeb\x06\x45\x09\x11\x2e\xf0\xba\xd8\x85\xb0\x78\xca\xc9\xd6\xf9\xba\xd4\x6b\x54\xce\x64\x50\xf6\x56\xae\x5c\x0b\xef\x95\x0e\xb8\x2a\xe2\xb1\xd9\x31\x43\xd6\xdf\x91\x5e\xce\x7e\xa4\xfe\x49\x8b\x05\x6c\x67\x9f\x6c\x11\xd8\xd0\x06\x85\xfd\x31\x5d\x37\xfc\xfe\xcb\xad\x4d\x24\xa3\x3d\x76\x31\x3a\xdb\x30\x4f\xe4\x53\xe6\x8c\x06\xe7\x24\xcb\xa3\x14\x7f\x3a\x49\xbe\x50\xe6\xa6\xf3\x7a\x7f\xc5\xf9\xfc\xac\x4f\xc0\x27\xf9\xbd\xbc\xb4\xa6\xbb\xa8\xdb\x70\x52\x36\xed\xa5\x70\xfd\xa7\x77\x96\x5b\x37\x44\xe2\x9a\xbf\x70\x71\x97\xbc\x19\x99\xe5\x45\x31\xaf\x9e\xbf\x37\xac\xc5\xcc\x6e\x0a\x86\x0e\x17\x3e\x2f\xfa\x01\xff\x55\x9d\x37\xa9\xba\xbc\x96\xf5\x41\x58\xaf\x7e\x07\x00\x00\xff\xff\xd5\xd9\xdb\x0b\xe1\x04\x00\x00")

func assetsRawGamefileTomlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRawGamefileToml,
		"assets/raw/Gamefile.toml",
	)
}

func assetsRawGamefileToml() (*asset, error) {
	bytes, err := assetsRawGamefileTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/raw/Gamefile.toml", size: 1249, mode: os.FileMode(420), modTime: time.Unix(1454902392, 0)}
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
	"assets/assets_suite_test.go": assetsAssets_suite_testGo,
	"assets/existence_test.go": assetsExistence_testGo,
	"assets/raw/Gamefile.toml": assetsRawGamefileToml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"assets_suite_test.go": &bintree{assetsAssets_suite_testGo, map[string]*bintree{}},
		"existence_test.go": &bintree{assetsExistence_testGo, map[string]*bintree{}},
		"raw": &bintree{nil, map[string]*bintree{
			"Gamefile.toml": &bintree{assetsRawGamefileToml, map[string]*bintree{}},
		}},
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

