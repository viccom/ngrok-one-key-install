// Code generated by go-bindata.
// sources:
// assets/server/tls/snakeoil.crt
// assets/server/tls/snakeoil.key
// DO NOT EDIT!

// +build release

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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
	name string
	size int64
	mode os.FileMode
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

var _assetsServerTlsSnakeoilCrt = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x93\xb9\x16\xaa\x4a\x02\x45\x73\xbe\xa2\x73\x56\x2f\x06\x41\x25\xac\x01\xb0\xd4\x02\x0a\x8a\x31\x03\x14\x07\x14\x45\x84\x12\xbe\xbe\xd7\xbd\x41\x07\xef\x9d\x70\x47\x67\x07\xfb\xbf\x7f\x06\x6d\x97\x78\xff\x41\x76\xc8\x89\x43\x10\xe0\xf6\x5f\x2a\x51\x42\xd0\xe7\x8e\x10\x28\x72\x84\x18\x5e\xf6\x59\xf1\x9e\xe5\x25\xc7\x1c\x78\xf0\xd2\xf6\xd7\xf6\xe6\x5a\x42\x85\x80\xc5\x0e\xc0\xa0\xa0\x61\x2d\x1c\x96\xe3\x84\x31\x8c\x81\x39\x48\x79\x4a\xc6\x73\x97\x4c\x95\xfb\x18\x8f\x4f\x6f\xaa\x38\x38\x3b\x42\xfd\x79\xdc\x16\x14\xb7\xc2\xc7\xd4\xa0\x7c\x5f\x3a\x42\x9d\x7d\x0e\x74\xca\xeb\xff\x33\x0a\xdb\x9f\xe4\x2c\x20\x81\x17\x2f\x81\x80\x52\xfc\xfc\x5d\xf3\x9b\xb9\x3e\xa5\xd6\x50\xa6\xc6\x98\xeb\xd6\x97\x12\x02\xc9\xfd\x9f\x6f\x6c\x07\x00\x1f\x01\xb6\x05\x7f\x14\x20\xba\x1c\x10\x60\x36\x58\xc1\xeb\x14\x7f\x05\xec\xdb\xb6\x42\xda\x32\x61\xaa\x24\xdd\x3b\x39\xec\xf0\x4a\x09\xd4\x6b\xd8\xf6\x99\x3f\xaa\xc2\xf5\xae\x32\x8e\x9d\xb4\x94\x03\x25\x94\xb2\xfd\x56\x5d\x7d\x8d\xbd\xa6\xff\x7a\xd2\x9f\x76\x22\xd2\x9b\xfe\x35\x6e\x0c\x43\x83\xf5\xe3\xf6\xf4\x81\x41\x6f\x94\x68\x55\x1f\xf5\xa6\xb2\x76\x22\x7b\x4e\x45\xb4\x28\x85\x0b\x0f\x81\x93\xec\xa4\xc8\xb4\xbc\x74\xef\x67\xd5\x12\x14\xfb\x4e\x0c\x76\x1f\xa0\xfb\x1c\x4c\xbd\x50\x4e\x9f\x87\xf5\xd2\x7d\xc2\x36\xf5\xf5\xbb\xb1\x32\x21\x66\x93\xb0\xf3\x7a\x5f\x33\xcb\x66\x3b\x23\xf6\x36\x53\x28\xc9\xa0\xb4\x85\x2d\x98\xd5\x24\xfe\xec\x69\xf5\xe9\x11\x04\x65\x55\xcc\x01\xba\x73\xf2\x73\x17\x04\x0b\xc3\x3f\x72\xc5\xf0\xce\x71\xd7\x1c\xcb\xed\x43\x2f\x55\x99\x44\xc3\xe2\x96\xf7\x8b\x3b\x4b\xe2\x3d\x78\x98\xd8\x3f\x58\x5b\x2d\x2d\x62\x33\xbb\xd2\x5b\xa6\xf3\xe6\x1d\x24\xe4\xf0\x94\xed\x5d\x44\x47\x60\x71\xe2\x76\xa9\x5c\xab\x6b\xf8\xa2\xe6\x63\x0b\x03\xbf\xaa\xdb\x66\x73\xea\x1e\x92\xf6\x8a\xfb\xc8\x81\x2b\x5d\xa9\xcf\x96\x55\x9f\xb2\x9f\xf1\x1e\x95\x02\x59\x51\x33\x9e\xe6\xcb\xf3\x20\x08\x06\x0c\x40\x0a\x54\x17\x45\xbd\x1b\x91\x6a\x85\x99\x0d\x21\x8b\x01\x30\x08\x94\x00\x43\xf6\xf9\x5d\xd4\x3d\x4c\xca\x18\x5c\xd5\xb0\xdd\xb5\x0c\x17\xeb\xa4\x92\xad\x2c\x7f\x6e\xe8\x71\x5b\x54\x7b\x00\xc0\x3a\x0c\x39\x9f\x7e\x9e\xf8\x6e\xb2\x40\x2e\x11\x51\x21\xdb\x3e\x5b\xe9\x3d\xce\xa9\x5e\xba\x07\x95\x37\xf7\xb5\xf5\x5d\xcc\x47\x23\x9b\x60\xb9\x0e\x40\xd7\xe4\xe4\x89\x3f\xb2\x62\x58\xaa\x6f\xb7\xb6\x16\x97\xb6\xdf\xeb\x75\x17\xf9\x4e\x86\x6c\x80\x07\x79\xda\x8f\xd2\x9c\xd8\xb3\xfe\x44\xa1\xe6\xcc\x82\x6b\x6d\xb8\x31\xf5\xf1\x15\xee\x3b\x2b\x92\x51\xb8\x45\xef\xf5\x7c\x7a\x0d\xd4\x57\x0c\xc0\x27\x9e\x53\xe7\xb5\xcc\x59\x17\x95\xc9\x65\xc3\x5f\xd3\x70\xe4\xd2\xb7\x59\x1f\x0f\xd9\x63\x6a\xbc\xe1\x36\x7c\x3e\xe6\x77\xcc\x02\xff\x1d\x17\x65\x5c\x7e\x94\xda\xd2\x88\xb1\xaa\x5e\x64\x38\xb3\x77\xed\xe1\xd5\xa1\x73\x77\xcd\xe1\xf7\x3c\xbe\xc1\x2c\xf2\x6f\x13\x5a\x92\x0f\x7d\xfd\xe6\xbc\x4e\xa5\xea\xac\x6a\x37\x8d\xe4\xb6\x4a\x23\x76\xf9\xb2\xae\xba\xe5\xc3\x9d\xb7\x0e\xa6\xfa\xf5\x8c\x8c\x24\x2d\x9e\x58\x4f\x7c\xfb\x2c\xb6\xe9\xbd\x61\xb7\xaf\x7e\x5a\xe0\x20\x31\x7f\x50\x3d\xee\x77\xd8\x0c\xc9\x63\x97\x7f\x94\xe2\x90\xa3\x42\x7b\x48\x7f\x63\xb5\x3d\xfc\xef\x80\xff\x17\x00\x00\xff\xff\x6f\x86\x84\x76\xdd\x03\x00\x00"

func assetsServerTlsSnakeoilCrtBytes() ([]byte, error) {
	return bindataRead(
		_assetsServerTlsSnakeoilCrt,
		"assets/server/tls/snakeoil.crt",
	)
}

func assetsServerTlsSnakeoilCrt() (*asset, error) {
	bytes, err := assetsServerTlsSnakeoilCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/server/tls/snakeoil.crt", size: 989, mode: os.FileMode(420), modTime: time.Unix(1471340282, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assetsServerTlsSnakeoilKey = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd5\xb7\x12\xab\x46\x00\x85\xe1\x9e\xa7\xb8\x3d\xe3\x41\x04\x81\x28\x97\x20\x40\x88\x9c\xe9\x60\x89\x22\x8b\xb0\x82\xa7\xf7\xf8\xd6\x3e\xed\x69\xfe\xee\xfb\xe7\xbf\x09\xb2\xa2\x99\x7f\x5c\x0f\xfc\xb1\x5d\x2d\x04\xbe\xfc\x47\x97\x93\xbf\x0f\x66\x68\x9a\x3c\x21\x4d\x00\x40\x17\x81\x23\x03\x5a\x68\x8e\x60\x43\xc2\xd2\x75\xb9\x48\x5e\x87\x64\x10\xe1\x38\x87\xba\x2a\xd1\x84\x7d\x6b\xdc\x6e\x89\xad\xfd\x86\x14\xb3\xc1\xa5\xe0\x19\x65\x18\x6e\x13\x6e\xfc\x7a\xdc\xe8\x8d\x79\x91\xd4\x6f\xd1\x96\x42\x45\x1e\x55\x2d\xd3\xce\x31\x0c\x29\xc0\xbe\x1d\x2c\xc0\x18\xad\xa1\x91\xf9\xe2\x2d\x77\x82\x7d\x7a\xf2\x19\x21\xef\x22\x52\x45\xd0\x31\xfb\x19\xaa\xde\x9d\x37\xa3\x97\x15\xe7\x97\x9d\xbe\x46\xb4\xca\x8b\x2d\x7e\x4e\xfb\x58\x10\x51\x7c\x7b\x7e\xa2\x2c\xcd\xe1\x60\xb3\x71\x7c\x8c\xd0\x79\xd7\x9c\x92\x7d\x41\x87\x97\x1d\x95\x09\x30\x93\x3b\x5c\x1c\x64\x32\x92\x91\xc3\x57\xa1\x75\x9a\x24\x2c\x7a\xdb\xce\xf2\xf4\xb4\xc5\x8f\xaf\xfd\x94\x4b\x14\x52\xc6\x7a\xfb\x04\x63\x96\xc1\x58\xbd\xb3\x47\x4f\x65\x37\x5c\xf3\xd6\x4b\xc9\xb0\x4f\xad\x9c\x68\x5e\x4d\x49\x93\x7f\x02\xe4\x3b\x23\x0d\xee\x71\x63\xb4\x31\xe5\x57\xb3\x1d\x6a\xfa\x80\xcb\xaa\x67\xec\x80\xf7\x35\x65\x8c\x70\x78\x63\x85\xc9\xb8\xf7\x0f\xc1\xb6\x72\xd8\x55\x18\x57\x8c\x3d\x39\x05\x8b\xf7\x14\x68\x8a\x80\x25\xcf\xc3\x22\xfe\x31\xf3\x4e\xa4\x22\xef\x55\x7b\x71\xd6\x83\x8e\x34\x09\x38\x40\x00\x93\x26\x00\x49\x1d\x17\x35\x0d\xdd\xe7\xd6\x19\xe2\x30\x62\xb9\x79\xee\x37\x3a\x28\x2e\x1d\x79\xeb\x32\x64\xab\xec\xc2\xb7\x44\x5c\xc8\xf1\x8d\x45\xa1\xce\x9d\xba\x4e\x53\x5d\x62\x6e\xeb\x5e\xf7\x7b\x5f\x48\x0f\x5b\x42\x7d\x58\x49\x87\xec\x6c\x35\x35\x60\x23\xfc\xe5\x73\xec\x27\xd1\x02\x69\x54\xd1\x4e\xad\x03\x97\xad\x59\xea\x8b\xea\xc1\xe4\x5f\xf8\xba\xa5\x97\xcf\x19\x33\xc7\xdb\x02\xa2\x5b\xb5\x72\xfb\x89\x1c\xc2\xab\x97\x69\xe6\x03\xcc\x1f\x36\x24\xa9\xe1\xe0\xf4\xab\xd0\x03\x92\x75\xcb\x60\x3c\xb8\x76\x62\xcb\xa3\x5c\x74\x7f\xe5\x35\x36\xad\x9d\x5e\x57\x0f\x4f\xcc\x6d\x3c\xc1\x4d\x00\xaa\x54\x96\x4a\xae\x50\xfa\xfb\x4a\xb7\xf9\x17\xc3\xc9\xb7\xf0\x1d\x7a\x3e\x09\xee\x82\x7d\xd6\x2d\x75\xf2\x47\xf1\x85\x89\x09\xf3\x6c\xb6\xb7\x84\x24\xac\xa5\x4c\x7b\xaf\xca\xc0\xe2\x5f\x83\x91\x86\x87\xd0\x76\xfb\xd1\xc8\x4e\x7b\x45\x1b\x87\x85\x5c\x19\x73\x6f\x94\x14\xd3\xe0\xbb\x96\x60\xd2\x77\x77\xf6\x89\xed\x1d\x53\xc5\x79\x21\xae\xab\x76\x50\xc3\xa0\x2f\xbb\xe1\x34\xfd\xf3\xa9\x44\x0f\x2f\x56\x39\xa7\x0e\x05\x28\xe2\x08\x1f\xb0\x75\x16\xe3\x52\x91\xc5\x3a\x91\x01\x9e\xe6\x1d\xb7\x29\x7d\xcc\xdf\x42\xcd\x58\x7f\x54\xe6\xb2\xa7\x4c\x24\x45\x9d\x84\x33\x7b\x2b\x96\x76\x7f\x4f\x96\x92\x14\x17\xe9\xa9\x0c\xc3\x8c\x65\x9d\x61\x1b\x6d\xc6\x1b\xa7\xce\x91\x59\x7c\x24\xc0\x54\x4a\x23\x96\x25\x5b\x07\x8a\xe1\xe6\x71\xa3\x92\xbc\x59\x79\x30\x4c\x89\x0e\xe4\x11\xb5\x79\x7b\xd9\xf0\x05\xdd\x53\x5b\x27\xb3\xef\x91\x26\x6c\xac\xaa\x3a\x0d\x50\x0a\xb2\xa2\x17\xc9\xec\xab\xc2\x73\x2d\xe2\x16\x3d\xbb\xab\x65\x1b\x0e\x89\xe3\x42\xcd\xad\x18\x8a\xe8\x68\x3d\x0e\x7f\x55\xf9\x59\xd3\xc6\xd5\x6e\x7f\x8b\x99\x1c\x87\x18\xfc\xc5\x5e\xc6\xe1\xf3\xba\xe0\x64\x7a\xb2\x7b\xf0\x6b\xc5\xbd\xe0\xe3\xcf\xe3\x47\x7e\xd0\xbb\xb0\x87\x2a\x3a\x11\xc0\x47\x3b\xc1\xb5\x8a\x28\x67\x29\xc8\x6a\x5e\x79\x88\xc4\xd4\x27\xda\x70\x60\xa0\x19\xae\xe3\xf1\xcc\xcd\x43\x8a\xe5\xc6\x32\xb9\x7c\x51\xcb\xc7\xe6\x66\xb7\x12\x29\x1f\x4f\xc1\x03\xaf\x48\x05\xcb\xbf\x0b\xbd\x7d\xd9\x49\xe1\x6c\x43\xd3\x43\xaa\xcd\xaf\x28\xd8\x3e\xb9\x8f\x91\x2d\x08\xa9\x01\x9f\xbb\xb2\x2e\x2c\xe0\xdd\xef\x5f\xa4\x02\x08\x00\x47\xba\x31\x67\x53\xbe\x69\x2f\xef\xe0\x0e\xc5\x3a\x01\xcd\x32\x49\x03\x59\x88\xe0\x07\x7b\xee\xdc\x2f\xa5\xf8\xca\x77\x4c\x7b\x20\x85\xa6\xd7\x87\xb7\x45\x45\x33\x6b\x03\x1e\xc0\xea\xd3\xd3\x1b\xe4\x86\x2a\xf8\x95\x6a\x71\xf1\xe4\xf1\x50\x45\x50\xef\xde\xfb\x1d\x92\xee\xdc\x66\x42\x34\xa2\x4f\x70\xf4\x87\x9a\x32\xd8\x2e\xdb\x6d\x5f\x08\xf6\xd3\xbd\x45\x4f\x6a\x72\xab\x63\x68\x73\x5b\xe8\xf2\xaf\x1d\xcd\x3c\x67\x3e\xaf\xd7\xb3\x1e\x09\xfd\x80\x01\x43\x4c\x5f\x5d\x22\x4b\x2b\x32\xc5\xfc\xa6\x17\xa6\x2c\x16\x2e\x86\x12\x6b\xd9\x2d\xa4\xe4\x4c\xa6\x43\x2b\x36\x7c\x25\x11\x13\x47\x17\x6a\x47\xd2\x51\xcf\xf8\x13\xfe\x26\x1f\x8d\xaf\x1b\x52\x50\xc1\x03\x7a\x6a\xc4\xde\x98\xf4\xb8\xfd\x7c\xe1\x59\x41\x26\xc6\xba\x9a\x95\x64\x69\x5a\xa7\xe5\x1d\x9e\xc6\x0c\x78\x8d\x8c\xd5\x3e\x1f\x38\xb4\xbe\xcc\x40\x6a\xc1\x4f\x55\xdc\x4f\x1c\x3d\xf1\xa0\x93\xa5\x69\xe0\x96\x70\xec\xc1\xea\x72\x57\xf6\x9c\x76\xa6\xc3\xa4\xc3\xd6\xeb\x16\xda\x7c\x52\x5b\x81\x6a\x35\xcb\x18\x58\xe9\x97\x6b\x92\x02\x5f\xf8\x98\x62\x6e\xa5\xbe\xb3\xad\x26\xaf\x45\xd0\x7d\x06\x6d\xe8\xc6\xe0\x31\xb7\x29\xba\x68\x25\x9c\x68\x42\xc5\x82\xf5\x46\x8b\x48\x17\x6a\x10\x7e\xf9\xd5\xa7\x93\xab\xdb\x1e\xc9\xf3\xfe\x85\xef\xeb\xcd\xa2\x2f\x1e\x08\x83\x5e\x90\x24\x61\x07\x4f\xbd\x67\xb5\x19\x5d\x6e\xa5\x5c\xd0\xf6\xfc\x69\x3a\xbf\x33\xa6\xcc\x5b\xcc\xc3\xe7\xcd\x71\xe6\x32\x69\xdf\xdd\xf1\x33\x82\x81\xc5\x49\x48\x31\x10\x16\x71\x64\x9c\xe3\x57\xee\x2e\xea\x7b\xa9\x55\xc6\xc4\x8a\x90\xd2\x9e\x8a\xef\xd1\x18\xc2\x43\x65\x4f\x19\xb3\x36\x17\x1c\x6c\xba\xbf\xec\x54\x41\xfd\x2b\xde\x90\x7c\x83\x8f\x40\x65\x1c\xca\x78\xdd\x89\x81\xf9\xc0\xb5\xa0\xec\x87\xc4\xe3\x04\xe1\x5d\x87\xcc\xf2\x3c\xf6\x97\x14\xd9\x94\xfe\x9f\x9a\x7f\x03\x00\x00\xff\xff\x3f\x96\x0e\xcb\x8b\x06\x00\x00"

func assetsServerTlsSnakeoilKeyBytes() ([]byte, error) {
	return bindataRead(
		_assetsServerTlsSnakeoilKey,
		"assets/server/tls/snakeoil.key",
	)
}

func assetsServerTlsSnakeoilKey() (*asset, error) {
	bytes, err := assetsServerTlsSnakeoilKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/server/tls/snakeoil.key", size: 1675, mode: os.FileMode(420), modTime: time.Unix(1471340282, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"assets/server/tls/snakeoil.crt": assetsServerTlsSnakeoilCrt,
	"assets/server/tls/snakeoil.key": assetsServerTlsSnakeoilKey,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"server": &bintree{nil, map[string]*bintree{
			"tls": &bintree{nil, map[string]*bintree{
				"snakeoil.crt": &bintree{assetsServerTlsSnakeoilCrt, map[string]*bintree{
				}},
				"snakeoil.key": &bintree{assetsServerTlsSnakeoilKey, map[string]*bintree{
				}},
			}},
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

