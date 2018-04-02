// Code generated by go-bindata. DO NOT EDIT.
// sources:
// deploy/data/virtlet-ds.yaml
package tools

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

var _deployDataVirtletDsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xdb\x73\x1a\xbb\x19\x7f\xf7\x5f\xf1\x4d\x3c\xd3\x24\x0f\x6b\xec\x4c\x4f\x93\xc3\x4c\x1f\x1c\xc3\x71\x99\xd8\xc0\x00\x71\xfa\xc6\x08\xed\xc7\xa2\x83\x56\xda\x4a\xda\xb5\xe9\x5f\xdf\x91\xf6\xc2\xde\xc0\x6b\xc7\xf6\x34\xd3\xf2\x04\x2b\x7d\x3f\x7d\xf7\x8b\x58\xcf\xf3\x4e\x48\xc4\xee\x50\x69\x26\x45\x1f\xf0\xc1\xa0\xb0\x5f\x75\x2f\xb9\x58\xa1\x21\x17\x27\x5b\x26\xfc\x3e\x0c\x08\x86\x52\xcc\xd1\x9c\x84\x68\x88\x4f\x0c\xe9\x9f\x00\x08\x12\x62\x1f\x12\xa6\x0c\x47\x93\xfd\xd6\x11\xa1\xd8\x87\x6d\xbc\x42\x4f\xef\xb4\xc1\xf0\x44\x47\x48\xed\x76\x83\x61\xc4\x89\x41\xfb\x1d\xa0\x0c\x64\x3f\x75\x30\xfb\xe1\x64\x85\x5c\xe7\x3b\x00\x54\x2c\x0c\xab\x6f\xcb\xe1\xed\x67\x23\xb5\x19\xa3\xb9\x97\x6a\xdb\x07\xa3\x62\xcc\x9e\xfb\x42\x4f\x25\x67\x74\xd7\x87\x2b\x1e\x6b\x83\xea\x0f\xa6\xb4\xf9\xc1\xcc\xe6\x1f\x29\x49\xb6\xf1\xd4\x41\x4c\x47\x03\x60\xda\x01\x80\x91\xf0\xe1\xe2\x23\xa0\x20\x2b\x8e\x70\x77\xab\xed\x13\x1d\xab\x84\x25\x98\xf3\x01\x54\x0a\x43\x98\x40\x05\x0a\xb5\x21\x6a\x0f\xf7\xc1\x48\x58\x21\xd0\x0d\xd2\x2d\xfa\x1f\x81\x08\x1f\x3e\x7c\xfa\x68\x41\x32\x48\xb3\x41\x88\x35\x82\x5c\x83\xd0\x28\x0c\x2a\x60\x02\x98\x60\x25\xd8\x92\x78\xd3\xd1\xa0\x22\xda\x29\xac\xa4\x34\xda\x28\x12\x41\xa4\x24\x45\x3f\x56\x08\x02\xd1\x77\x9c\x52\x85\xc4\x20\x10\x8b\xb5\x66\x41\x48\x22\x8b\x5e\x32\xcf\xde\x6a\x19\xa0\x46\x95\x30\x8a\x97\x94\xca\x58\x98\x71\xc5\x2c\xc5\x99\x52\xf0\x9d\x35\x07\xdc\x65\x1a\x88\xa4\xaf\x41\x0a\x27\x8d\x90\x3e\x6a\xb8\x67\x66\x63\x3d\x4a\x91\x59\x6a\xb6\xbf\xe7\xda\x72\x66\xcd\xa0\xc8\x7a\x6d\x45\xdd\xed\x8d\x6c\xa9\x2f\x1b\x4f\x01\x14\xfe\x2b\x66\x0a\xfd\x41\xac\x98\x08\xe6\x74\x83\x7e\xcc\x99\x08\x46\x81\x90\xc5\xe3\xe1\x03\xd2\xd8\x58\x67\x2e\x51\xa6\x98\x73\xe4\x48\x8d\x54\x0b\x54\xa1\xae\x2e\x7b\x10\x12\x43\x37\xc3\x87\x48\xa1\x76\xfe\x5f\x5d\xb7\x3b\xb6\xb8\xeb\x57\xc4\xa9\xed\x00\x90\x11\x2a\x62\xa4\xea\xc3\x48\x34\x16\x13\xc2\x63\x6c\xc0\x5a\xe0\x9a\x6e\xad\xdc\x57\xb9\xdd\x0b\x82\x53\x58\x6c\xb0\xe6\x14\x40\x65\xc4\x50\xe7\x00\xef\x35\xac\x39\x3e\x24\x92\xc7\x21\x82\xaf\x58\x52\xf8\xcd\xa9\xf5\x04\x6b\x19\x1f\xd7\x24\xe6\xc6\xd9\xdf\x59\x8d\xc7\x01\x13\xe0\x33\xe5\x1c\x13\x85\x8e\x15\x6a\x30\x1b\xb2\xf7\x60\x47\xc7\x94\xd3\x9d\x3d\xce\xba\x16\xfa\xb0\xda\x01\x67\x2b\x7b\x36\xfc\xa5\x88\x03\x7c\x60\xda\xe4\x6e\x60\xbd\xf5\x24\x97\x32\x0d\xef\x48\x61\x44\x14\x7a\xd6\x1e\x85\x2a\x58\x48\x02\xec\x43\xc8\x14\x11\x86\xe9\x5e\x35\x07\x64\xeb\xd3\x98\xf3\x3c\x84\x47\xeb\xb1\x34\x53\x85\x36\x5a\x8a\x5d\x54\x86\x21\x11\xfe\x5e\xc3\x1e\xf4\xca\xc7\x9d\xe9\x4d\xb1\x94\xea\xe8\xd6\xfa\xb7\x2e\x13\xa4\x4c\x6e\xbf\x68\x6f\xaf\x49\x2f\xd5\x91\xf6\x7c\xa6\x4a\xd6\x0b\x2d\xf1\x94\x98\x4d\x1f\x7a\x99\x36\xbd\x2a\x41\x03\x57\xc5\xa2\x01\xa0\x64\x44\x02\xe2\x1c\x16\xbe\xb2\x54\xcd\x4c\x0a\xc2\x0f\x1c\x55\xc6\xc8\x71\x7d\x49\xb7\xa8\xb4\xa4\xdb\x03\x44\x09\x51\x96\xb0\x97\x6e\x3c\xab\xec\xcc\x41\xb8\x0c\x0e\x50\x5b\x33\x96\x57\x4f\x61\x2d\x55\xea\x2a\x4c\x04\xce\x57\xd2\x23\x38\x5b\xf5\x32\x97\xe8\x39\x9b\xe9\xd4\x1f\x5c\x5e\xa8\x58\x3c\x3f\x34\x21\xca\xe3\x6c\x75\xe4\x60\xaf\xbe\xa5\x10\x1a\x93\x03\x64\xe5\x15\x8d\x34\x56\xcc\xec\x6c\x40\xe1\x83\x29\x87\x5f\xa4\x58\xc2\x38\x06\xe8\x57\xd2\x29\x00\x8a\xa4\xe9\x13\x77\xa3\xd9\xe2\x66\xb8\x58\x0e\x46\xf3\xcb\xaf\x37\xc3\xe5\xb7\xbb\xdb\x12\x96\x0b\xee\x3f\x94\x0c\xab\xf1\x9d\xa6\xdc\x5b\x12\x7d\xc3\xdd\x0c\xd7\xf5\xe0\xaf\x54\x3c\x2f\xdd\x5c\xdb\xe2\x92\x8e\xcf\xb4\x2d\x14\xcb\x6d\x12\xd6\x96\x65\x94\x3a\x4b\x26\xc1\x49\x71\x6c\x2d\x7b\x14\x46\x4e\xad\xd3\x39\xf2\x4e\x61\x24\x80\x12\x8d\x70\x6f\x93\xcf\x9f\x48\x0d\x70\x49\x09\x2f\x02\xde\x21\xd8\xd5\x7b\x22\x8c\xcd\x32\xb6\x92\x31\x03\x42\x1a\x90\xeb\x35\xa3\x8c\x70\xbe\x03\x92\x10\xc6\x5d\xb5\x93\x02\x5f\x20\xb0\x33\x41\xba\xc4\x74\xd9\x35\xf4\x4e\xf7\xd6\xba\x47\x03\x25\xe3\xe8\xa4\x6e\x89\xda\xe3\x2a\xa9\x75\xee\x50\xfa\x31\x47\xdd\x20\x6c\x3e\x57\x48\xfc\x89\xe0\xbb\x9a\x6b\x55\x21\x6d\xdd\x6e\x60\xd5\x1e\x76\x02\xaa\x26\x96\x9f\x49\x36\x5e\x23\x6f\x58\xb9\xeb\x6e\xd1\xde\xac\x1d\xa2\xae\x3b\x1d\x1c\x70\xc6\x26\xb5\xcd\x59\x8f\x50\x7b\x36\x99\xa1\xd1\x25\x97\xb5\xa5\x87\xcb\xc0\x15\x35\x56\x94\xab\x0d\x2a\x84\x15\x52\xe2\x5a\x2d\xb3\x41\x75\xcf\x34\x16\x25\xec\x9e\x71\x6e\xdb\x27\x3f\xa6\x08\xa8\x94\x54\x65\x48\xce\xb6\xb6\x4f\x63\x25\xc7\x3a\x85\xef\x59\xfb\x26\x6d\x55\xf3\xb2\x3e\x8b\x6e\x88\xf2\x31\x81\x35\xe3\x08\xef\x53\x1d\xc8\xa0\x97\x84\xba\x47\xd6\xfe\xe7\xdf\x56\xab\x95\xf7\x05\x7f\xff\xec\x5d\x5c\xe0\x67\xef\xf7\xdf\xfe\x76\xe1\x9d\x7f\xfa\xeb\xa7\x73\x42\xcf\xcf\xcf\xcf\x3f\xf5\x28\x53\x4a\x6a\x2f\x09\x97\xe7\x67\x5c\x06\xef\xfb\x30\xb6\xdd\x26\xdd\xa4\x88\x52\x15\xa5\x78\xd7\xcc\xa6\xa1\xf6\x0e\xa7\xf1\x12\x2b\xcd\xe4\x9f\x29\xf3\x71\xea\xa6\xd1\xfe\x9f\x8e\x8b\x4f\x2d\x1d\xd7\xac\xd3\xd6\xd5\xfc\xb7\xe7\xde\x2e\x49\xf5\xe5\xb2\xcf\x69\x9a\x13\x5d\x2b\x5a\x4e\xb8\x40\x14\x16\xed\xbf\x6d\x3c\x75\x1c\xa1\x0a\x99\xf8\x05\x73\xf5\xaf\x98\x53\xab\x28\xb1\x76\x3c\xe0\x03\x52\xd7\xf9\x2a\x81\x06\x75\xd1\x04\x67\xdd\x6f\x2f\x75\x9d\x9e\xdd\xd6\x38\xa8\x43\x87\xdd\x2e\x77\x76\x48\xcf\x4e\x9b\xad\xa8\x76\xa1\xb5\x53\xef\x52\xfb\x9e\x9f\x46\xcb\x3b\xd0\xd0\xdc\xba\x59\x1b\xdc\xe0\xd4\x3d\xf6\xec\x77\xcf\x28\x22\x34\x77\x8c\x35\xf3\xb2\x93\xe6\x30\x2f\x7b\xed\x57\x36\xfd\x2f\x66\xd9\x16\xa6\x27\x3f\xc6\x37\x93\xcb\xc1\x72\x3a\x9b\x2c\x26\x57\x93\x9b\xb7\x63\x5d\xde\x0b\x2e\x89\xbf\x8c\x94\x34\x92\x4a\xfe\x3c\x01\x6e\x26\xd7\x37\xc3\xbb\xe1\xdb\xf1\xcd\x65\xc0\x31\xc1\x67\xb2\x7b\x75\x79\x33\xba\x9a\x2c\xe7\xdf\xbf\x8e\x87\x8b\x37\xe3\x99\x12\xce\xa8\xf4\x74\xbc\x12\x95\x64\xda\x81\xf1\xd1\xed\xe5\xf5\x70\x39\x1b\x5e\x0f\xff\x39\x5d\x2e\x66\x97\xe3\xf9\xcd\xe5\x62\x34\x19\xbf\x19\xef\x2e\x0b\x2c\x15\x06\xf8\x10\x2d\x4b\x79\xe0\x99\xfe\x9e\x05\xe9\xcd\xe4\xfa\x7a\x34\xbe\x7e\xf3\x40\xe5\x32\x08\x98\xa8\x6f\xe9\xc8\xfc\x7c\x36\x9a\xdc\x2d\xe7\xdf\xa7\xd3\xc9\xec\xed\x9c\x47\x2b\x26\x93\xa5\x8e\xa3\x48\xaa\x27\x3a\x4f\xce\xf8\xec\xf2\xc7\x72\x30\xbc\x1b\x5d\x0d\xe7\x6f\xc6\xb6\x22\xf7\x4b\x1f\x13\x46\x2b\x95\xa5\x03\xd3\xa9\xc7\x97\x5c\x7d\xbe\x1c\x8c\x66\x75\xbe\x8f\x97\xb0\x1c\xeb\xdb\xf7\xaf\xc3\xd9\x78\xb8\x18\xce\x97\xd3\xc9\xc0\x7a\x5d\x43\x01\x7d\x78\x57\x2b\x51\xef\xea\x3d\x71\xa9\x88\xbe\xde\x2d\x60\x12\xea\xa7\x5e\x14\xbc\x75\x97\xf6\xb3\x03\x5c\xb7\xf1\xcb\x3b\xd8\x03\x1f\xef\x9e\x53\x8d\x95\xee\xa2\x2d\x6a\x69\xb6\x5e\x4b\xe5\x2e\x7b\xad\x67\x42\xea\x99\x40\x28\x45\xad\x0b\x7b\xbb\x7f\x2e\x2c\x7e\xb9\x19\x69\x72\x58\x97\xe6\x28\x61\xfb\x6d\x4e\xcb\x5d\xce\x51\x94\xb6\x61\xa1\x4d\x4d\x47\x41\x2a\x93\x40\x63\x38\x38\x4a\x5a\x1e\x95\xea\xc3\xd3\x29\x2c\x26\x83\x89\xed\x28\xc4\x7b\x03\x1b\xa2\x7c\x2a\x7d\xcc\x2e\x7c\x21\x6d\xd3\xdd\x50\x68\xc3\xde\x5d\x71\xec\x09\x37\x4c\xa7\xb7\x1a\xd9\xe0\x04\x57\xb3\x11\x44\x4a\x3e\xec\x80\x09\x6d\x08\x4f\x8b\x8d\x9d\x1b\xcb\x07\x32\x91\x9a\xd2\x79\xc4\xfe\xbf\xa1\xb3\x2e\xa2\x1c\xbb\x67\x3e\x70\x55\xfd\x28\x5e\x5b\x14\xb6\xc5\x60\x27\xa0\x7a\xe0\xb5\x85\xe3\xe3\x40\xa5\x08\xad\xdf\x9d\x1f\x25\xfe\x89\xc1\xa9\xe3\xd8\xd4\x49\x09\xad\x33\xd4\xc1\x09\xaa\x0b\x64\xdd\x30\x95\x2b\xfb\x2e\xfa\x2c\x8a\x4d\x39\xb7\xb5\xe5\xc4\x4e\x60\x47\xad\xfc\x14\xb0\xb6\x59\xf9\xd8\xa4\xdc\x89\xbb\x16\xb5\xd7\xc6\xbc\x4e\x7c\x09\x34\xa2\x8e\xe3\x9e\x55\x6c\x57\x74\x1c\xfd\x43\x25\xcc\x4b\x67\xd1\xd6\x31\xf4\xf8\xb0\x5a\x7f\x79\x40\xad\x08\x3d\x23\xb1\xd9\x48\xc5\xfe\xed\xf6\x9c\x6d\xbf\xe8\x33\x26\x6b\xef\x12\x64\x7f\xc0\xcf\x24\xc7\xaf\x4c\xf8\xb6\x67\x3c\xfc\x52\x81\x92\x1c\xb3\x56\x89\x44\xec\xda\x26\xf5\x23\x27\x9d\x00\x34\xce\x68\x40\xea\x78\xf5\x27\x52\x5b\xfb\xbd\x6c\xf7\xbc\xf2\x8f\x77\xf7\x17\x1b\xac\x06\x9a\xe7\x3d\x4d\x27\xcf\x78\x9f\x42\xd9\xaa\x64\xf7\x7b\x85\x4e\x4a\xff\xf4\xbc\x4b\x9b\x2c\x85\x5a\xc6\x8a\x62\x69\xa5\x78\x03\x20\xb5\x70\x82\x6a\x55\x5e\x75\x17\xd8\x2f\x61\xd5\x16\x99\x8a\xa3\x3d\x85\xc4\x47\x95\xcb\x50\x93\x20\xe3\xbe\xc2\x7b\x8d\xef\x82\x6b\x0f\x02\xa7\x24\x0f\x38\xd3\xe9\x97\x7b\x62\xe8\xe6\x95\x24\xc8\xc3\x25\xd6\xa8\xec\xca\x4f\x0b\xe2\x81\x46\xaa\xd2\xe4\x51\x13\xea\x55\x23\x2b\x2f\x37\x42\xfa\xe8\xad\xb2\x6d\x2f\x18\x66\x0d\x53\x97\xe3\xed\x29\xe0\xd7\x59\x07\x97\xc2\xa6\xbe\xdf\x77\xaf\x96\xbc\x72\xea\x09\xf7\x46\x7e\x05\xfd\x1c\x72\xa4\x5f\x24\x2d\x79\x54\xf9\x8f\x65\x20\x12\xb1\xfd\x9b\x64\x19\xf2\xe1\xac\x14\x6b\x23\xc3\x7c\xc1\x47\xf7\xe2\x4f\x51\x86\xda\x93\xd4\xa1\x83\x33\x1e\xed\x91\x07\xcf\xcb\xf6\xb8\xaa\x16\x92\x28\x62\x22\x68\x3d\x2a\xcb\x2a\xe9\x8f\x57\x8f\xca\x8a\x76\x5f\xde\xd9\x2c\xec\xcb\x3a\x58\x59\x13\x49\x2e\x6a\x0d\xf0\x19\xb5\xed\x3f\x01\x00\x00\xff\xff\x6e\x0c\xc2\x5f\x92\x28\x00\x00")

func deployDataVirtletDsYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployDataVirtletDsYaml,
		"deploy/data/virtlet-ds.yaml",
	)
}

func deployDataVirtletDsYaml() (*asset, error) {
	bytes, err := deployDataVirtletDsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/data/virtlet-ds.yaml", size: 10386, mode: os.FileMode(420), modTime: time.Unix(1522279343, 0)}
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
	"deploy/data/virtlet-ds.yaml": deployDataVirtletDsYaml,
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
	"deploy": &bintree{nil, map[string]*bintree{
		"data": &bintree{nil, map[string]*bintree{
			"virtlet-ds.yaml": &bintree{deployDataVirtletDsYaml, map[string]*bintree{}},
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
