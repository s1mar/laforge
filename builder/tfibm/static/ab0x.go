// Code generated by fileb0x at "2018-09-27 20:41:05.971639445 -0700 PDT m=+0.009429257" from config file "assets.toml" DO NOT EDIT.
// modification hash(74f406fa37c4290844ed4bb8afc25436.e5b9c5ef4c0b7aef8593382d0449dfd6)

package static

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileCommandTfTmpl is "command.tf.tmpl"
var FileCommandTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x93\x31\x8f\xd4\x30\x10\x85\x7b\xff\x8a\x91\x45\x01\x12\x44\x14\x88\xee\x0a\xb8\x02\xe8\x10\x14\x14\xe8\x64\xe5\xe2\xd9\xdb\x11\xf1\x4c\xe4\x99\xec\xde\x29\xca\x7f\x47\x93\x90\xb0\xc5\x22\x41\x81\x9b\x64\xf4\x66\x9e\x9f\x3f\xd9\xd3\x04\x19\x0f\xc4\x08\xb1\x93\x52\x5a\xce\x11\xe6\x39\x54\x54\x19\x6b\x87\x10\x79\xec\xfb\xb4\x95\x11\xe2\x50\xe5\x44\x4a\xc2\x69\x9a\xa0\xf9\x80\x06\x71\x53\x13\xb7\x05\x7d\x3c\xa9\xe1\xb0\xcb\x5e\x24\x1e\xcb\x3d\x56\x17\x23\x4c\x01\x20\xe3\x80\x9c\x35\x09\xc3\x0d\x7c\x0f\x00\x00\x91\xee\x4b\xea\xa4\x0c\xa3\x61\x3a\x95\x44\xac\xd6\x72\x87\xcd\x9f\x37\x8a\x01\xe0\x2e\x04\x80\x3d\x15\x56\x6f\x2b\x62\xf8\x0a\x1f\xb1\x5b\x37\x03\x98\x26\xa0\x03\x34\x1f\x45\xad\xf9\xa4\xdf\x88\xb3\x9c\xd5\x0f\x0a\xcb\xea\x84\x19\x3b\x23\xe1\x5f\xfd\xbe\x8e\xa2\xb6\xfc\xdc\x40\x7c\x36\xfd\x7b\xb8\x86\x86\xd3\x9b\xd4\xe6\x5c\x51\x75\x89\xba\x2e\x7b\x1a\x70\xf3\x3d\x13\xd7\xf2\x5b\x1a\x15\xeb\x26\xbd\xcb\x85\x98\xd4\x6a\x6b\x52\x2f\xa6\xa9\xa0\x8c\xb6\xb4\xbc\x7d\x7d\x31\x3b\xb4\xaa\x67\xa9\xd9\x05\x0f\x75\x2b\x65\x40\x23\x3f\x54\xf3\x45\xc4\x3e\x6f\xfa\xbc\x67\x99\x37\x36\xd8\x2b\xfe\x2d\x8d\xff\x08\x64\xb5\x56\x3d\x5e\x21\xb2\x6a\x55\xc4\xae\xb0\x80\x2b\x38\x2a\x9d\x5a\xc3\xf4\x03\x9f\xd6\xbc\x07\xea\xf1\xb9\x93\x21\xce\xf8\x08\xcd\xfb\x91\xfa\xdc\xdc\x0a\x1f\xe8\xc1\xc3\xf6\x49\xf5\x98\x2e\xc6\x92\x4f\x2c\xb7\xec\xc5\x15\x62\xec\x20\xc3\x52\x12\xf7\xfe\x80\xb6\x7b\x0c\x3b\x7f\x7f\x4e\xdb\xf7\xab\x55\xe2\x07\x77\x7b\xb9\x74\xdd\x05\x77\x9b\xc3\xee\xf5\x33\x00\x00\xff\xff\x0d\x32\x43\xcf\x8a\x03\x00\x00")

// FileDNSRecordTfTmpl is "dns_record.tf.tmpl"
var FileDNSRecordTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x64\x8f\x31\x4b\x04\x31\x14\x84\xfb\xf7\x2b\x86\xd4\xb2\x1c\x68\xbb\x9d\x60\x77\x85\x82\x85\x22\x21\x5e\x9e\xb0\x70\x97\x2c\xef\x65\xb7\x30\xbc\xff\x2e\xc9\xde\x5e\xe1\x75\xc9\xcc\x97\x99\x49\xad\x88\xfc\x33\x25\x86\x8b\x49\xbd\xf0\x29\x4b\x74\x30\xa3\x59\xf2\x3a\x45\x96\x6e\x38\x54\x02\x96\x39\x86\xc2\xfd\x08\x28\xcb\xca\x82\x11\xae\x56\x0c\x2f\x5c\xb6\x84\x4d\xf6\xc2\x9a\x17\x39\x71\x8b\x1a\xa6\x79\x7d\xf2\x21\x46\x61\x55\x47\x80\x91\x11\xed\xc4\xf6\x2c\x5c\xab\xbd\x72\x71\x70\xbd\x5c\xa7\x9c\xfc\x2d\x7c\xe7\x7d\x0a\x97\x1e\xeb\xb5\xf0\x7c\xb3\xdb\xc5\xa7\xe5\xf2\xcd\xd2\xcc\x6d\xf0\x6f\x4e\xbc\x2f\x7c\x3e\xbe\xbd\xf6\x8a\xe1\xa3\xa9\x66\x43\x9b\xd2\xc2\xee\x89\x63\x53\xcd\x1a\x70\x9d\xcd\x8a\x11\x9f\xfd\xe3\xff\xd8\xf7\x70\x5e\x3a\xfc\x40\xc0\x17\x01\xa5\x9c\x31\xe2\xf1\x70\x20\xa3\x5a\xc1\x29\xc2\xec\x2f\x00\x00\xff\xff\x4f\xd5\x7d\xe3\x68\x01\x00\x00")

// FileInfraTfTmpl is "infra.tf.tmpl"
var FileInfraTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x5a\x6d\x4f\x1b\xbb\x12\xfe\x9e\x5f\x31\xda\x82\x2e\x48\x64\x9b\x17\xa0\xb4\x52\x3e\x40\xa1\xbd\xe8\xf6\x86\x8a\x97\x7b\x2a\x1d\x55\x96\xb3\x9e\x04\x8b\x5d\x7b\xaf\xed\xdd\x42\x51\xfe\xfb\x91\xbd\xbb\xc9\x6e\x5e\x68\x02\x34\xb4\xe7\x04\x24\x14\x76\xc6\x33\xf3\x3c\x1e\x8f\x3d\xeb\xbc\x7a\xfc\x4f\xed\x15\x7c\x3a\xfc\x70\x76\xfe\xf1\x04\x3e\x9e\x74\x4f\xce\x0f\x2f\x4f\x8e\xe1\xf2\xe4\xfc\xdc\x3e\xfc\x2f\xbc\x3f\xeb\x7e\x38\xfd\x78\x75\x7e\x78\x79\x7a\xd6\xad\xbd\x82\x7a\x1d\xfe\x38\x3c\xef\x9e\x76\x3f\x42\xbd\x5e\x7b\x05\x97\xd7\x5c\x43\x9f\x87\x08\x5c\x03\x4d\x8c\x8c\xa8\xe1\x01\x0d\xc3\x3b\x18\xa0\x40\x45\x0d\x32\x1f\x8e\x25\x08\x69\x00\x19\x37\xc0\xcd\xbf\x74\xed\x15\x04\x52\x18\x14\x46\x03\xe3\x0a\x03\x13\xde\xf9\x70\xa5\x11\x3e\xd1\xbe\x54\x03\x04\x2a\x18\x28\x84\x5e\xc2\x43\x06\xa6\x70\xe2\xd7\x9e\x82\xb4\x66\x50\x29\x6b\x3f\x82\xfb\x1a\x40\x8f\x06\x37\x28\x18\x78\x68\x02\x96\xb6\x3d\xf7\x10\x20\x56\xd8\xe7\xb7\xd0\x01\x4f\x1b\x6a\xf0\xf5\xfd\x3d\x6c\xf8\x47\x36\x0e\xff\xf4\x18\x86\xc3\xec\xc1\x25\xd2\xc8\xfd\xe9\x26\x51\x0f\x95\x7d\x3e\xb2\xee\x9b\xbe\x1b\xea\x39\x7b\x28\x58\x2c\xb9\x45\xda\x81\x3f\xdd\x13\x00\xef\xda\x98\x58\xbf\x7b\x6d\x6d\x71\xc1\xf0\x76\xe4\xe2\xbd\x14\x7d\x3e\xc8\x62\x22\x11\xd5\x06\x95\x07\xc3\xa1\xb7\xb3\xdc\x48\x1d\xd2\x14\x4b\x03\xbf\xba\xbf\x89\x46\x25\x68\x84\x16\xdc\xc3\xe3\x0b\x4d\x67\x22\xa3\x85\x6a\xfd\x4d\x2a\xf6\xe3\xb1\x85\x66\x31\x76\x58\x1b\xd6\x6a\x29\x55\x9c\xf6\x42\x04\x2f\x8d\x34\xff\x8e\x19\xdb\xe6\x2e\x76\xc1\x44\x34\xb6\x9a\x0c\xfb\x34\x09\x0d\x74\xf2\xa9\xf0\x74\x44\xc3\xd0\xb3\x1a\x47\x9f\x5a\xa4\xf5\x65\xf7\x4b\xb3\xd1\xc8\xe2\xf1\x22\x64\x3c\x89\x46\xc2\xdd\x2f\xcd\xfd\x92\x34\xa4\x6a\x80\x23\xe1\x41\x55\x78\x5b\x95\x36\xf7\xbf\xb4\x5b\x85\x78\x2a\x5a\xa9\x17\x8b\x35\xe9\x25\xc2\x24\xcd\x7d\x67\xf6\xea\xe8\xaa\x7b\x79\x45\x9a\xfb\x64\x7f\xd7\xab\xc8\x0f\x2a\xf2\x83\xb1\x3c\x40\x61\xa4\x7e\xe3\xc4\xef\x4f\xba\x97\x67\x17\xe4\xcd\xa4\x74\xbf\x2c\x2d\xd9\x66\xd8\xe3\x54\x64\xa6\x8f\x4f\x8e\x4e\x0f\xbb\xe4\x60\x52\xfa\xb6\x2c\x7d\x3b\x96\x7e\x6b\xdd\xb4\x9d\xe8\x8f\xd3\x2e\x69\x35\x1a\xed\xfa\xc5\xe5\x71\xfd\xe2\x73\xab\xbe\x57\x51\x3a\x28\x2b\x1d\x38\xa5\xf3\x56\xfd\xe2\x73\xb3\xa2\xd5\x6c\x95\xd4\x9a\x2d\xab\x56\x95\xef\x97\xe5\xfb\x25\xb9\x23\x3e\x56\x32\xe5\x0c\x15\x78\xbc\x17\x65\xbc\x6b\xd9\x37\x21\xbd\x43\x45\x16\x4a\xdf\xb1\x3a\x8d\xb9\x1b\x52\xe4\x61\x55\x72\x83\x77\x4b\xd8\xb9\xc1\xbb\xcc\xcc\xb0\x56\x53\xa8\x65\xa2\x02\x74\x31\x92\x40\x46\x71\x62\x90\x68\x7d\x9d\x69\x79\x93\x05\x83\xcc\x29\x18\x19\xbc\x90\xf6\x30\xcc\x23\xa9\x0e\x2b\x2c\xd6\xc0\x16\x4f\xb4\xf5\xc3\x43\x91\x76\x9c\xe2\x89\x48\xb9\x92\x22\x42\x61\x32\x75\x30\x48\xa3\xce\x3c\x57\x35\x80\x38\xe9\x85\x3c\x28\x70\x6f\xdc\xdb\x7a\xba\x35\x1f\xbf\xc2\xd0\x45\x30\x1e\x46\xec\x08\xc7\xc2\xf6\x2c\x22\x34\x06\x89\xe2\xe6\x8e\x0c\x94\x4c\x62\x0f\x3c\xbc\x35\xa4\x47\x75\xbe\xd4\x4b\x13\x37\x23\x7a\x62\xe6\x44\x4e\x46\x56\xdc\x9a\xd3\x81\xe2\xb1\xe1\x52\x58\x53\x27\xb7\xc6\x26\x44\x08\x85\x6b\x70\xae\xa1\x2f\x15\x3c\x92\x28\x17\x66\x67\xec\x73\x31\x94\x29\xe3\x4f\x07\x69\x8d\xac\x18\xa3\x73\xf9\x63\x88\x5c\x3c\xc7\x44\x8e\xac\x4c\x83\x3c\x15\x3f\x09\xe4\xd8\xe7\x62\x28\x71\xa0\x50\xeb\xa7\xe3\xcc\xed\xac\x18\x69\xe1\x75\x31\xac\x4f\x4e\xda\xc2\xc8\x8a\x51\x16\x49\x7b\x7f\x0f\x8a\x8a\x01\xc2\x86\x40\x63\xa5\x3b\xee\x13\xbc\xeb\x4c\x1a\x17\x41\x98\x30\x64\x5d\x34\xdf\xa4\xba\xd1\x30\x1c\xda\xc1\xbc\x0f\x22\x1b\x6c\xdd\x7b\x8e\x8e\xe1\x70\x11\xe6\x6c\x88\xf9\xb0\xa2\x8a\x3f\x85\xc3\xaa\xb9\x15\xb3\x39\xe1\xdc\x31\x63\x0f\xc5\xc3\xf2\xa7\xda\xc3\x09\x45\x54\x62\x37\x86\x51\xc1\x27\x8d\x8c\x94\xec\x64\x9f\xa3\xe0\x62\xb4\x26\xd0\x5c\xa3\x22\xc5\x79\xea\xf4\x73\xea\xf6\xff\x58\x2a\x43\xdc\x8c\x92\x88\xdb\x21\xcd\x89\x87\xd4\x1e\xca\xf7\xf7\xf6\xda\x7b\x56\xa0\xa4\x91\x81\x74\x3b\xa7\x09\xdc\x89\x4c\x61\x24\x0d\x12\x1e\x3f\xb8\xaf\x53\x16\x71\x41\x78\x6c\x27\xfb\x75\xbb\xe5\x0e\x06\x55\x34\x9c\x65\x1b\xe4\x34\x50\xbf\x40\xe8\x73\xf6\xe3\x2d\x70\x8a\x96\xe6\x6a\x69\x49\xd8\xef\x41\x4b\xeb\x59\x68\x69\xcc\xa2\xa5\xb5\xb7\x3b\x41\x0a\x0f\xa2\xdf\x83\x95\xf6\x0c\x56\x70\xa5\x4b\xa8\xe1\xbb\xdf\xd7\x8d\x55\xc2\xde\x5d\x29\xec\x59\x4b\xe4\x45\x60\xef\x3d\x07\xec\xa7\x2c\x81\x55\xa3\x4e\x19\x9f\xb9\x4d\x2c\x0d\xba\xd5\x9a\x89\xba\xf5\x93\x41\xa7\x8c\x3f\x0a\x73\xf3\x39\x30\xb7\xdb\x07\x6f\x67\xa1\x2e\x9e\xff\xbc\x85\xbd\x2c\xec\xf1\xb1\xf4\x79\x66\xfb\x91\x05\xed\x87\xd8\x8a\x46\x21\x03\xb7\x0c\x2b\x63\x84\x8f\x25\xa6\xf9\x82\x25\xef\x97\x26\xa6\xf5\x72\x45\xf1\x97\xe6\xa5\xfd\x82\x07\xa6\x5f\x93\x98\x79\xbb\xc9\xca\x9a\x8e\x9f\x4c\xcb\x63\xca\xee\xbc\xdd\x66\x65\x1d\xc7\x0a\x39\x79\x81\x77\x01\x65\xaa\x2b\x4d\xf4\x6f\x90\x88\x23\xea\x96\xe5\xbc\x02\x74\xe9\x8c\xac\xd2\xf4\xeb\xe7\xe6\x33\xd3\x34\xf3\xcd\x4a\x25\x73\x77\x60\xe3\x5a\x6a\xa3\xa7\xf3\xf6\xdf\x52\x9b\xa3\xbb\x3c\x6b\xed\x40\x80\x2c\x6d\xf1\xff\x59\xa6\x8f\x72\x16\xc0\x89\x72\xa3\x9c\xe5\x36\x9d\xc9\xcc\x78\xae\xe4\xd4\x36\x8a\xb9\x23\xee\x0d\xd6\xbb\x0e\xc4\x8a\x0b\xd3\x07\xcf\x6c\xb2\xfa\xa6\xae\x6f\x6a\x6f\xc6\x1b\x24\xe7\xd2\x99\x73\x91\xb9\xb1\x23\xb3\xb3\xef\x47\xd2\x88\x70\xa1\x0d\x15\x01\xe6\x77\x24\x55\xd7\xc5\x9b\xb4\xec\xe7\xba\xb0\xda\x99\xa3\x3b\xd2\x64\x32\xa2\x2e\x13\x2c\x03\xbe\xa0\x36\x95\x68\x18\xc4\x26\xf0\xa5\x1a\x8c\xf5\xa4\x26\x0a\xfb\xa8\x50\x04\x48\x02\xc9\x30\x9b\xc1\x94\x2a\x3f\x8d\xa4\xfe\xd3\xb9\x71\x88\xce\x2e\xac\x83\xaf\x65\x1f\xd4\xd0\x00\x85\x41\xf5\xe0\xdb\x82\xb1\x9a\x57\x09\x51\x64\xd3\x46\x74\x8c\x68\x13\xa7\xd9\x68\x34\x4a\x48\x13\x15\xde\x91\x1e\x0f\x43\x2e\x06\xd0\x01\xa3\x12\x1c\x49\xfb\x21\x4d\xa5\x72\xb7\x2f\x05\x1d\x45\xcc\x9a\x7f\xc7\x52\xd4\xa7\x39\xb9\x17\xfc\x3b\x4e\xc6\x9f\xdf\x24\x11\xce\xca\x97\xd0\xee\x4a\x2e\xcb\xe1\x89\x4b\x2c\x7f\xd1\x2b\x2c\x97\xdb\x3b\x23\x7b\x5f\x6b\xa3\x8f\xf9\xb5\xd1\xd4\xba\x99\x13\xc0\x43\x6d\xe5\xce\x42\xea\xc5\x62\x9d\x1d\x8d\xe2\x29\xb5\xf8\x1e\x17\x4e\x65\xa3\xda\x59\x48\xbd\x74\x90\x59\x6c\xc0\x83\xf1\x5b\xfa\x13\xe6\xea\xc2\x56\x9e\x7b\x6e\xd2\x2f\xdc\x7b\x62\x5d\x7d\xf8\x3f\xaa\x34\x78\x89\x46\x45\x6c\x46\x92\xec\x65\x32\xe1\xcc\xdb\xde\x1e\xaf\xd2\xec\xdb\x00\x24\x42\x43\xad\xd6\xc4\x9d\x60\x31\xfb\xe7\x18\x1e\x6a\x8d\xe6\x83\x54\x76\xe2\x67\x24\xc1\x56\x51\x30\x36\x35\x71\xc5\xa2\x5a\x15\x6c\xd8\xfe\x11\xd5\xb8\x5d\x5c\x1c\x66\xce\x87\x05\x38\x99\x98\x38\x31\x30\x4e\xe4\x52\x41\xf1\x8b\x79\xe3\x71\xb9\x38\xa4\x34\x4c\x70\x5c\x80\x67\x54\x18\x7f\x56\xcd\xf0\x79\x9c\xee\x12\xca\x98\x3b\x44\xe7\xa6\x97\x0d\x28\x4b\xeb\xe7\x8f\x67\x3a\x8e\xac\xbe\x0f\x0c\x6c\x85\x28\xf2\x50\x3e\x2b\x99\x72\xcd\xa5\x40\xa5\xb7\xa1\x51\x9e\xcd\x71\xcd\xd7\x06\x63\x22\xdc\xec\xec\xc0\x46\xac\x64\x3a\x2a\xfe\x15\x03\xe5\xd1\x59\x8a\xe9\xa4\xe7\xbe\x96\x73\x9b\x9f\x9c\xde\x87\x52\xe0\x94\x5a\xbe\xeb\x94\xdc\x54\x23\x29\xac\x11\x67\x64\x6c\xd3\xbf\xb0\xfb\x14\xc3\x18\x05\xd3\x44\x0a\x6f\x9c\x3b\xf3\x68\x73\x09\x55\xe1\x6d\x7b\x2a\x1c\x0c\x35\x3e\xd1\xbd\x48\xc2\x90\x14\x6e\xfc\xb8\xe0\x88\x6c\x6a\x62\x51\x6e\xb2\xc9\x28\x60\xeb\x18\x03\x55\xe1\x60\x7b\x46\x64\xd9\x1e\x3f\x8b\x3c\xeb\xc3\xff\x0f\x17\x0c\x3c\x26\xec\xce\x14\xe4\x5f\x98\x59\x10\x46\x25\x1a\x6f\x3a\xb7\x16\x65\x43\x68\xa2\x51\xa5\xa8\x46\xf0\x3d\xf0\x9a\xad\x37\xee\x0d\x51\x73\x89\x80\x4a\x4c\x94\x88\x75\xcc\x95\x49\x5a\xc8\xe0\xa1\x31\x34\xb8\xce\x53\x37\x3b\xbd\x4c\x0f\x33\x18\xc5\x21\x35\x58\x25\xb0\x9c\xc2\x8f\x9b\x8e\xfc\x80\x58\x7c\xd7\x61\xb5\xf3\xf1\x82\x2c\x56\x70\x3f\x9d\xc6\x40\x46\x11\x15\xec\x1f\x45\xe1\x08\xf3\xd3\xe9\xcb\xf6\xec\x7f\x14\x7b\x05\xe4\x25\xc8\x9b\x7c\x52\xfd\xbf\xfc\xdf\xc4\x3e\xf1\xf7\x6b\x90\x26\x5d\xcd\x6a\x90\xf2\xb6\xd4\x1e\x3d\xe6\x7d\x21\x61\xdd\x40\xfd\xad\x1b\xa8\xdf\xab\x25\x2a\xe5\xeb\xba\x33\x5a\x77\x46\xeb\xce\x68\xdd\x19\xad\x3b\xa3\x75\x67\xb4\xee\x8c\xd6\x9d\xd1\xba\x33\x7a\xce\xce\x68\xea\x56\xea\xaf\x00\x00\x00\xff\xff\x7d\xfa\xea\x15\x50\x38\x00\x00")

// FileOriginalInfraTfTmpl is "original_infra.tf.tmpl"
var FileOriginalInfraTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x5d\x7b\x73\xdb\x38\x92\xff\x5f\x9f\x02\x27\x3b\x95\xdd\x9c\x45\x39\xce\x26\xb5\x9b\xad\x5c\x95\xd7\x72\x1c\xdf\x78\x6c\x95\xe4\x24\x55\x37\x93\x61\xd1\x24\x2c\x62\x4c\x12\x0c\x00\x4a\x71\x3c\xfa\xee\x57\x00\x48\x8a\xe0\x9b\x92\x6c\x67\x3c\xc8\xd6\x66\x22\x11\x8f\x06\xfa\x81\xee\x5f\xa3\xc5\x9d\xf5\xff\xf4\x76\xc0\xd9\xe1\xfb\x8b\xc9\xc9\x31\x38\x39\x3e\x3f\x9e\x1c\x5e\x1e\x8f\xc0\xe5\xf1\x64\xc2\xbf\xfc\x19\x1c\x5d\x9c\xbf\x3f\x3d\xf9\x38\x39\xbc\x3c\xbd\x38\xef\xed\x80\xc1\x00\x7c\x3e\x9c\x9c\x9f\x9e\x9f\x80\xc1\xa0\xb7\x03\x2e\x5d\x44\xc1\x35\xf2\x20\x40\x14\x58\x11\xc3\xbe\xc5\x90\x6d\x79\xde\x2d\x98\xc1\x00\x12\x8b\x41\xc7\x00\x23\x0c\x02\xcc\x00\x74\x10\x03\x88\x3d\xa7\xbd\x1d\x60\xe3\x80\xc1\x80\x51\xe0\x20\x02\x6d\xe6\xdd\x1a\xe0\x23\x85\xe0\xcc\xba\xc6\x64\x06\x81\x15\x38\x80\x40\x70\x15\x21\xcf\x01\x2c\x99\xc4\xe8\x6d\xb2\xd2\xde\x0e\x38\xfc\x3c\x05\x47\x38\xb8\x46\xb3\x88\x58\x0c\xe1\xa0\x17\x12\x3c\x47\x0e\x24\xa0\x6f\x2d\x68\x1f\xdc\xf5\x00\xb0\x6c\x1b\x52\x6a\xde\xc0\x5b\xf0\x0e\xf4\xef\xee\xc0\xae\x71\x84\xfd\x10\x32\xc4\x7b\x18\x87\x9f\xa7\xc6\xe1\xf8\xf4\x27\x78\x0b\x96\xcb\x7e\x0f\x00\x0a\x6d\x02\x59\x53\xfb\xa9\x68\x15\x77\x21\x70\x86\x70\x50\xd9\x7c\x22\x1f\xcb\xb6\x21\xc1\x62\x83\xdf\x81\x7e\xbf\xb7\xec\xf5\x78\x0f\xe4\x80\xb7\xef\x40\x48\x50\xc0\xae\x41\xff\x19\x7d\xe6\xf4\xc1\xae\x71\x1c\xcc\x11\xc1\x81\x0f\x03\x66\x8c\x09\xbc\x46\xdf\xc0\xae\x31\xc6\xce\xe9\x08\x2c\x97\xbd\x1e\x83\x84\xf0\xed\xf5\xc5\x2a\xaf\x2c\xfb\x06\x06\x0e\xe8\xd3\x57\x72\xd9\x00\x5c\x45\xf6\x0d\x64\x95\x54\x4d\x99\xc5\xe0\x94\x61\x02\x8d\xff\xc8\x96\x92\xc0\x6e\xcb\x01\xa0\x6a\xa3\x46\xd8\xb7\x10\x6f\x36\x14\x8f\xb2\xab\x39\xb7\x7c\x98\x3e\x48\x56\x34\x4c\x17\x64\xb0\x6b\xca\x89\x93\xe3\x77\xe5\xdf\x5a\x1c\xec\xba\x68\x18\xd8\xe4\x36\xe4\x9b\xcb\x48\x04\x7b\x00\x2c\x39\x2f\x77\x00\x27\x63\x6c\x21\xd2\x23\x90\xe2\x88\xd8\x50\x08\x22\x27\xc5\x0c\x2d\x44\xfa\x72\x74\xe4\x80\xe5\xd2\xa4\xd4\xe5\x0f\xe4\xf7\x9c\x65\xbc\x55\xc0\xb7\xe6\x5d\x55\x33\x2e\x3f\xd1\x95\x87\xec\xcc\xe2\x14\x52\xa7\xd3\x0f\x63\xd1\x80\x13\xf2\x07\x70\x99\xef\x09\x9a\x05\x71\xf3\xd0\x7e\x9b\x8e\x9b\xa3\x70\x1e\xda\x0a\x71\xe2\x33\x27\xca\x46\x0e\x31\xaf\x3c\x6c\xdf\x54\xee\xcd\xd1\xe9\x68\x12\xef\x0c\x0c\xac\x2b\x0f\x9a\x4e\x40\x4d\x17\x53\xc6\x57\x43\x93\x4d\xea\x01\xc0\xac\x19\x8d\xa5\xf3\xbc\xb0\x50\x3e\xa5\x78\x94\x11\x95\x74\xce\x12\xf1\x91\x8d\x2f\xa1\xe5\x2b\xe3\xf4\x57\xdc\x40\xb3\x45\xe5\x82\x51\xc0\x20\x09\x20\x33\x67\x16\x83\x0b\xeb\x56\x59\x3d\x9a\x2d\xe4\xea\xe7\xa1\x6d\x22\x87\x8f\xbf\x7b\x17\x6f\x93\xa1\x50\x6c\x20\x27\xd9\x5e\x82\x23\x06\x4d\xc6\x37\xa0\x72\xd6\x4c\x1b\x65\x42\x07\x5e\x5b\x91\x97\x21\xa6\xc3\xe4\xdc\x00\xf1\x61\xe3\x8d\x55\x19\xb6\x6f\x88\xff\x0d\xf7\xe5\x6e\xc5\xe3\x2b\xc3\xe6\x77\xc2\x50\xf6\x41\xce\xc1\xb7\xb4\x89\x7f\x68\xb6\x90\xeb\xdb\x32\x17\xe7\x0e\x32\x69\x74\x15\x40\x56\xb9\xad\xf2\xb1\x2a\xc0\x69\x2f\x65\x33\xf9\x9f\xa6\x0d\x2d\x15\xfa\x2c\xe9\xff\x1b\xf9\xe1\x07\x4c\x19\xcd\x8a\xbe\x6f\x85\x66\xac\x9c\x28\x34\x71\x60\x7a\x56\x14\xd8\x6e\x6b\xe9\x5f\xd1\xbb\xfd\xed\x13\x6c\x69\x23\x94\xa6\x45\x29\xb6\x91\x38\x4c\x95\xed\x24\x4c\x3e\x92\x9b\x29\x09\x4d\xf6\x73\xb5\xa1\xf2\x7b\xa3\x6c\x55\xc9\xd6\x66\x27\xcb\x4a\x61\xe6\x7b\xa3\x52\x2f\x32\xda\xe6\xb8\x76\x58\x67\xcd\x4c\xde\xc0\xc4\x21\x5f\x09\x55\x75\xcd\xb5\x43\xb9\x0c\x47\x1c\x51\x8a\xcd\x2d\x18\xb7\xd1\x87\xa3\xf1\x85\x1c\xc5\x18\x9d\x4f\x33\xdb\x9e\xe9\x6d\x52\x48\xe6\x90\x70\x53\xf7\x8b\x60\xc8\xdd\x1d\x20\x56\x30\x83\x60\xd7\xdc\x03\xbb\xc2\x0c\x8a\x16\xfc\x9c\xaf\x9f\xe3\x3c\x6d\x4b\xf9\xba\xe4\x06\x0b\xd2\x32\xa3\x2c\x97\xfd\xbd\x64\x1e\x7e\xea\x8b\x86\x5f\x7a\x00\x04\x2c\x6c\xa2\x85\x85\xad\x49\xb9\x1c\x4f\xcb\x29\x49\xc7\xa8\x20\xa4\x49\xd8\x05\x0b\xb6\x2b\xe6\x82\xdd\x42\x44\x5b\x4b\x45\xa5\xb0\xaf\xc6\x2a\xd8\x8e\x96\x06\x44\x99\x26\x67\xc3\x15\x1a\x0c\x75\xd6\x8c\x80\xd3\x99\x79\x65\xd1\x6a\xad\xa5\xd0\x8e\x08\x62\xb7\xe6\x8c\xe0\x28\x54\xe8\xe7\xfd\x24\xe5\x45\x77\x42\x3c\xe3\x14\x42\x6a\x13\x24\x88\xe0\x0d\x12\x4f\x7d\x00\xd4\xa6\x60\x7a\xd2\xef\x74\x14\xc1\x19\x81\x34\xe1\xfc\x35\xc1\xbe\x19\x62\xc2\x79\xbb\x2f\xbe\x61\x58\xfd\x1c\x12\xcc\xb0\x8d\x3d\x3e\xf4\xe0\x65\x3f\x77\x82\x09\x21\xce\x9c\x61\x5f\x92\x83\x08\x05\x5b\x9f\x26\x16\x70\x55\x5b\x0e\x1d\x1f\x05\xa7\xe3\x12\x5d\x91\x0f\x32\x9a\x11\xeb\x46\xd2\x23\xd5\x0c\x55\x37\xa4\x9a\x36\xac\xe1\x9f\x8d\x6b\x40\xb6\x1f\xb6\xde\xac\x5a\x55\x4c\xe4\x61\x8b\xaa\x48\x67\xdc\xee\xaf\x23\xb8\x73\x07\x55\xc9\x2d\x7f\xd4\x4e\x6c\xe7\x0e\xea\x2c\xb5\x55\xbc\x58\x69\xfc\xc1\x81\xc2\x14\x50\x7c\x94\xf2\x27\x6b\x27\x98\x5d\xc2\xa7\xe4\x69\x22\x73\x19\xa6\xed\xb5\x93\x91\xd5\x0c\xaf\x5e\xfd\xf3\x5f\x95\x94\xa5\x0f\x1f\x92\xb6\x2e\x3a\xa8\xca\x42\x56\x0f\x13\x8f\x42\x69\x60\xa8\x72\x2b\xd8\x97\xa3\xaa\xc9\xc9\xda\xaa\xac\xf3\x6f\x66\x30\x80\x14\xad\xe2\x1d\x35\x94\x1f\x3c\xa3\xf9\x58\xfe\x44\x76\xe0\x0e\xa4\xf1\x21\xe9\x14\x6b\x4a\x4e\x55\x20\x4a\xf4\xa3\x30\xcb\x72\x29\x9f\xde\x09\x4e\x50\x66\x05\x36\xcc\xfa\xf4\xf2\x1b\xa3\xa2\x6f\x7a\x4e\xc1\x10\x06\x0e\x35\x85\x36\xfd\xd2\x6f\xd3\x55\xd8\x95\x79\x68\x27\x82\x22\xbc\xdb\x02\xe5\xc9\x28\xd5\xe4\x73\xd2\x7b\x00\x58\x3e\x2a\xe5\x40\x76\x97\x4e\x20\x3b\xfc\xf9\xb4\x29\x2a\x4f\xa6\x34\xd9\x6d\x08\x1b\xc7\x3c\x8d\x5b\x4f\xd1\xf7\x84\xd3\xd9\xf8\x5b\x6e\x63\x12\xb7\x1b\xa5\xe1\xb8\x91\x74\x90\xb8\x51\xea\x13\x77\x71\x87\xe1\x15\x15\xae\x80\x8f\xbe\x43\x27\x13\x2b\xf0\x3e\x76\x44\x19\xf6\xe5\xe9\x73\x94\xfc\xbb\x36\x12\xd9\xaf\x59\xf0\x99\x45\xd9\x85\xcd\x04\xe4\x91\xcc\x20\x6d\xe9\xd8\x62\x2e\x9f\x63\x2a\x3e\x4d\x60\x20\x00\xb4\x84\x65\xd1\x55\x14\xb0\xc8\x8c\x1c\x6a\x50\xb7\xaf\x32\x41\x9d\x2e\x05\x73\x02\xe4\x89\xff\x97\xb1\x9d\x4f\x1d\x51\x48\x4c\xc7\x62\x96\xdc\xaa\x6b\xe4\xc1\xbf\xf5\x73\xf4\x2c\x97\xfd\xbf\x4b\xdb\x1c\x12\x34\xb7\x18\x34\x51\x98\x30\x35\xdd\x98\xa5\x6c\xc1\xed\xbc\x6a\x28\x4c\xe4\xac\x8c\x49\x47\x53\xd2\xd4\x7c\xee\xa0\x4c\x6b\xd5\x5f\x10\x06\x4f\x38\x0b\x55\x6c\xf8\x64\x79\xc8\xb9\x3c\x1a\x8f\x31\x61\x59\xbf\xba\x61\xd2\x38\xb2\xa4\x33\x93\xd9\xa1\xc9\xbf\x17\x53\xc5\x9a\x5c\xf4\xc1\xd7\x24\xec\xe3\x68\x7d\xc2\x22\xa7\x0d\x61\x75\xc1\x41\x99\x95\xe8\x6e\xae\xcf\x21\x5b\x60\x72\xa3\x58\xec\x41\x6a\xf9\x2b\x8c\xb9\x6a\xc1\x82\xc8\xf3\xcc\xe4\x63\x1f\xf4\xed\x18\x64\x86\x66\x8d\x31\x03\x80\x11\x34\x9b\xf1\x70\x49\x2e\x6d\x7d\xb3\xbc\x94\x72\x8f\xe7\x88\x22\x1c\x70\x6d\x24\xd0\xc7\x0c\x0e\xe0\x37\x68\x27\xf8\xae\x8d\x83\x00\xda\x42\x0d\xef\x62\x66\xf1\xa1\xd4\x08\x05\xa2\xd0\xa8\x39\x3f\x8c\x14\xb1\x88\x37\x0f\x00\x61\x3d\xe3\x31\x28\x75\x93\xaf\xb9\xce\x26\x5f\x4b\x93\x90\x76\x40\x3e\xc4\x11\x13\x4f\xde\xec\xfb\xc9\xd7\x89\xda\xc6\x48\x65\x56\xcd\x0b\x90\xa5\x6c\xfa\x13\xbc\xcd\xaa\x3e\x1f\x44\x6c\x04\xdf\x49\x0f\x05\x30\xeb\x1d\xd0\xc8\xc1\x80\x42\x07\x0c\x10\x78\x4e\x87\xbf\x19\x2f\x28\x75\x07\x43\xf1\xd7\xbf\x9f\x83\x21\xc1\x98\x0d\x0d\x4a\xdd\xa1\x15\x31\x17\x13\x6e\x5e\x39\x29\x34\xef\x2f\x70\x7a\x38\x65\x1f\x43\x0f\x5b\x0e\xad\xd7\x11\xd9\xe8\x3d\xf2\x60\xac\x21\x6a\xef\x23\x1c\x05\x42\xcb\x3c\x18\xa8\xa3\x26\x8d\xd1\x35\x98\xb1\x62\x97\xfd\xa2\xd2\x7a\xd8\xb6\x3c\x3e\xd3\x1e\xd8\x95\xcc\xe7\x1f\x04\x79\x85\x81\x63\xef\x6a\x25\x2d\xbc\x45\x3f\x95\x8a\x52\x51\xd9\x96\xb8\xd4\x88\x4c\x5e\x6c\x38\x4f\x94\x6e\xe5\x82\xb3\x35\xe1\xc9\x08\x90\xf0\x35\xa5\x72\x27\xd4\xf0\xb1\xd2\x4d\x4e\x8d\x07\xff\xe3\x40\xca\x50\x60\xb1\x4c\x3e\x20\xc3\x81\x55\xcb\x65\xc1\xb8\xad\xfe\xad\x1c\xb0\xaa\x5c\x54\x4a\x97\x3c\x7f\x0b\xc2\x92\x1d\x63\x3f\x1e\x3b\x6f\xde\x69\xe2\x79\xb6\x1c\x3d\x1e\x21\xe7\x00\x8c\x64\x9b\x64\xb8\xb6\x47\x7d\xe5\x9c\xe5\x0e\x40\x3a\x3b\xba\x06\x01\x54\x68\xe8\x4f\x8f\x26\xa7\xe3\x4b\x73\x7c\x38\x99\x9e\x9e\x9f\x98\xc7\x93\xc9\xc5\xa4\x9f\x8d\xb1\x6b\xc5\xbc\x52\xd0\xb7\x27\xea\xb5\xc2\x5e\x2f\xee\xb5\x02\xbf\x45\x91\x4f\x04\xb3\x52\xea\x55\x2f\x2b\xd3\x36\x27\xf8\x43\xe6\x87\x22\x47\x47\x95\xc3\x38\xa7\x56\x4d\xc7\x94\x66\x4c\x96\x31\x99\x0f\x85\x43\x2d\x3e\xda\x6c\xd7\xc7\x0e\xf8\xef\x6f\xa0\xb8\xff\x7b\x6a\xcb\xc6\x06\xc4\x07\x83\xeb\x86\x71\xbe\xf4\xf2\x62\x53\xf0\x23\x4b\xec\xdb\xb2\xd7\xc3\x11\x0b\x23\x06\xfa\x29\x3f\xa8\x51\xeb\x1d\xcd\x2d\x2f\x82\xeb\xb1\x59\xc0\x4a\x49\xa6\x1f\x02\xe6\x42\x30\x3a\x9f\x82\x71\x9c\xed\xcf\xa4\xfd\x9d\x20\x4e\xfb\x47\xa1\x63\xa5\x39\xb1\x18\xad\x5e\x53\xc4\xe2\x50\x7f\xf8\x02\x5c\x5e\x8c\x2e\xde\x82\x4f\xa3\xd3\xd4\xc9\x3c\x22\x50\xea\xcb\x8b\xe1\x52\x02\x02\xb1\x55\x5e\xfc\xee\xee\x81\x5d\x93\x9b\xd5\xf3\xca\xa0\xed\x33\x0a\x1c\xbc\xa0\x86\x34\xee\xf1\x00\xbb\x65\x48\xc2\x42\xb6\x9c\x3b\x68\xff\x99\x23\x41\x85\xc5\xef\x6e\x0a\x1b\x08\x3f\xe4\xab\x13\xa8\xe0\x83\x21\xda\xa5\xc3\x95\xfa\xcc\xd5\x90\xc3\x3a\x50\xc3\x1a\x10\xc3\xa6\xd0\x42\x7b\x48\x21\xde\x6d\xbe\xfd\x87\x3f\x9f\xb6\x46\x0d\x8a\xec\x7a\x52\x98\xc1\xc1\xbe\x14\xa5\x76\xa8\xc0\xef\x91\x1f\x9a\xb1\x30\x0a\x4c\xe0\x9b\xef\x75\x04\x05\x9e\x2c\x18\x50\x17\xda\x3e\x46\x48\x9b\xd7\xbd\xb2\x6c\x60\x3c\x70\xc1\x73\x9c\x40\x8a\xbd\x39\x74\xe2\x99\x55\xb7\x31\x1e\x20\x59\xd3\x9e\xfc\xa7\x18\x24\x1e\x30\xed\x2f\x04\x4d\x49\x92\x14\x0c\x80\xdc\x48\x91\x57\x8c\x3b\x4f\xa3\x2b\x27\xb9\x3d\x64\xe6\x77\x6f\xaf\xf1\x94\xfa\xd2\x25\x98\xdf\x30\x88\xaf\x0a\xde\x5b\x39\xff\x25\x96\xa5\x65\x08\xd0\x3a\x02\x68\x98\x62\xab\x71\x40\x41\xbb\x7f\x30\x4f\x7f\x1d\x47\x72\x81\x02\xe2\x57\xbb\x92\x22\xdb\x87\x28\x23\x16\xc3\xa4\x8b\x4f\x69\x51\xba\xc0\xc4\x29\x55\xfd\x31\x76\xc6\xc9\xf3\xd5\xfd\xb8\x4a\x3f\x72\x03\x0f\xff\xe8\xed\xd0\x93\x19\xb4\x47\xf4\xf3\x9f\x28\x5b\xaa\xdc\xfb\x10\x2f\x20\xa1\x2e\xf4\x3c\x30\x38\xc7\xe3\xf8\x3a\xe8\xe0\xf8\x1b\xb4\x23\xbe\x6f\x63\xec\x21\xfb\x16\xfc\xe7\x96\x13\x03\xaa\x38\xb4\x9e\x07\xdf\x5b\xe9\x63\xa5\x59\x69\xc7\xe9\xce\x88\x52\x77\x24\xa9\xc0\xe3\x56\x1c\xae\x03\x95\x36\xe1\x6e\x86\xb7\xa5\x9c\xed\x13\xdf\x41\x04\x0c\x29\x18\x7e\x05\xbf\x66\x34\xeb\xd7\x3e\xf8\xe3\x0f\x30\x87\xe4\xe5\xff\x04\x91\x97\x61\xdc\x97\x16\x30\x52\xc1\xf9\x15\xd7\xa4\x5e\xbf\x32\x09\xb4\x31\x71\x0a\x37\xde\xa0\x2d\xae\x23\x95\x1e\x6c\xdf\x71\x90\x5c\xb9\x2a\xbf\xe2\xfa\xfa\xd5\xff\xe1\x00\xa6\xeb\x16\x9d\x57\x36\x45\x84\x17\xf2\x49\xc2\x25\xce\x07\xf1\x99\x79\x09\x5f\xde\xec\xcb\x0b\xd1\x9c\xba\x82\xd3\xd5\x42\x16\xf6\x92\xf3\xbb\x22\xba\x6c\x1b\x55\x36\x44\x93\x19\x95\xe8\x18\xb2\xfd\x64\x79\xa8\x45\xbc\x76\x63\x79\x48\x07\x6b\x1b\x04\x6b\x7c\x9f\x75\xa4\xa6\x44\x6a\xaf\xba\x47\x6a\x9b\x24\x6f\x75\x9c\xa6\xe3\xb4\x27\x1f\xa7\x6d\x35\xc9\xba\x95\xe4\x2a\xb4\x0f\x06\xfc\xf3\xc3\xa4\x57\xcb\xb3\xab\x76\x44\x3c\x30\xa0\x67\x60\x80\x25\x7a\xcc\x09\xe2\xba\x6f\x50\x17\x3c\x77\x19\x0b\xdf\x0e\x87\x2f\xdf\xfc\xcb\x38\x78\xfd\x0f\x23\xfe\xef\xd0\xb3\x18\xa4\x4c\x34\x1d\xf0\xb6\xcf\x53\x11\x90\xd9\xda\x2b\x8b\xba\x85\xd1\x56\x6d\x32\x58\x75\xf1\x79\x36\x5d\xdb\x25\xe8\x15\xa7\x75\x2e\xe2\xbd\xdf\x78\xb7\x26\x6c\x2d\xf5\x99\x37\xf3\x98\x2b\x73\x1e\xd5\xf9\x8e\x4a\x07\x79\x4b\x79\x8e\xc4\x84\xb4\x0d\x4b\xdb\x25\x9d\x62\xf7\xbb\x39\x34\xf9\x6b\x6d\x72\x49\x28\xd2\x2e\x7f\x54\x9b\x3b\x6a\xcc\x1b\x7d\x49\xad\x87\x0e\x55\x1e\x30\x54\xc9\xc4\x2a\x21\x26\x8c\x1b\xc0\xd8\xdf\x39\xe1\xee\x4e\xa5\x77\x20\x4b\x02\x2f\x8f\xc6\x25\x21\x44\xcd\x6d\xec\xd5\xd5\x2e\x21\x04\xd9\xa9\x8c\x71\x72\x99\x37\x3e\xf9\xe3\x4b\x5f\x55\xb7\xb7\xd7\x19\xaa\xdd\x6d\x6f\x39\x32\xa8\x1d\x76\x98\x19\x76\x6b\x57\xc3\xdf\x95\xcc\xf9\x9e\x60\x7f\x2c\xe7\xc9\x5d\x81\x2e\xb6\xbd\xc4\xd9\x96\xd9\xcb\xd1\xb5\x6b\x79\x94\x7b\xd3\x6b\xb2\x6f\x9b\xf7\xac\x8b\xe1\xfa\x1a\x2a\xf0\x71\xa4\x55\x40\xab\xc0\x9f\x52\x05\x2a\xce\x81\xce\xc1\x61\x6f\x07\x64\x8a\x69\x93\x48\x2e\x26\xa2\x4b\x81\x6d\x49\xf7\xad\x54\xda\x26\x63\x6e\xaf\xb4\xb6\x8c\xd2\xad\xb2\x66\x07\x24\xe5\xb1\x9d\x36\xb5\x4d\xdd\x6d\xc9\x78\xeb\xd6\xe2\x96\x0c\xb5\xed\xa2\x5c\x3a\xeb\x26\x55\xd5\x86\xb7\x52\xba\x82\xb6\xec\x6d\x61\x38\x4b\x7a\x6e\xd1\x36\x76\xa9\x82\xf2\xae\xb3\x3f\x25\x91\xe6\x7e\x53\xea\x3e\x8d\x4e\x3f\x21\x8a\xae\xbc\x38\xce\x7c\xdc\xda\xab\x04\x4c\xcb\x98\xcf\x5c\x96\xe9\x51\xf5\x71\x4d\xe4\x2b\xc5\xf2\xf3\x98\xac\x62\x91\xf6\x65\x9b\x5c\x39\x4d\x25\x2e\xdf\x0d\x34\x2b\x81\xee\x41\x05\xa0\xd5\x6d\xe0\xb6\xf8\x7e\x47\x8c\x4f\x49\x01\xb4\x4e\x02\x74\x9c\xa4\x90\x27\x70\x37\xac\x09\x73\x1f\xad\xfe\xab\xc6\x04\x57\xa5\x0a\x8a\x40\x7a\x2c\xa2\x8f\x02\xa3\x57\x2d\xa0\xb6\x16\xca\xfd\x11\xea\x9e\xdc\x1f\xa4\xc6\xa9\xa3\xf0\x8b\xfe\x69\x95\xe8\x7d\x24\x29\x4a\x2d\x71\x75\xaa\x82\x60\xcc\xa4\xc3\x66\x3a\x70\x8e\xec\xe4\x8a\xad\x03\x3d\xc8\x20\x77\xce\x18\x24\xfe\x0a\xd0\x8b\x0f\x35\x00\xe6\xd8\x8b\x7c\x68\x52\xae\x6f\x32\x18\x48\xf4\x78\x84\xe8\x4d\xac\x86\xea\xf1\x07\xbf\xc6\x8d\x2e\xa6\xa0\xbf\x38\xb8\x79\xf9\xa6\x5f\x0d\xe2\xaa\xd9\xaf\xb5\xae\x28\xa6\x9e\xb4\x3c\x2d\x0a\xd8\x6e\x87\x14\x58\xb1\x16\xa4\x74\x3d\x36\x26\xf0\xa9\xad\x89\x7e\xf5\x9e\xd8\x92\x0e\x9e\xd0\x7a\xe2\xda\xbd\xb6\x0b\x5a\x23\x81\xfc\xa0\xcb\xb1\x61\xc0\x30\x6d\xbd\x1c\xd9\xfc\x07\x59\x4e\x49\xd6\x57\x4d\x82\x1a\xcd\x75\xa7\x65\x07\xaf\x1c\x17\x06\xf6\xed\xea\xe4\x1b\x25\xdf\x21\x98\x3d\xf8\xea\xa6\x5b\x9d\x54\xab\x01\x8d\xe4\xf0\x28\x3e\x11\x05\x5f\x15\xbf\x07\x24\x9d\xeb\x3a\x8b\x2e\x0e\xd5\xc5\xb4\x85\x6a\x85\x0b\x23\xa4\x2f\xb7\xc1\xbb\xb6\xa9\xe7\xee\xfe\xaa\x08\xc1\xd4\xfc\x74\x63\x86\x7a\x3d\x87\xfe\xfe\x32\x63\xdd\x23\x97\xc6\xe4\x59\xee\x46\x5f\x8b\xfb\x7c\xd5\x79\xb4\xcc\x5d\xbe\x33\xeb\x3d\x0f\xed\x2f\xa1\x1f\x26\x57\xf8\xfe\xeb\xe5\xc1\xab\x7f\xbc\x5e\x65\xcb\x6a\xd2\x65\xc9\x22\x0a\xc9\xad\xe2\xb6\x8a\x72\xd1\xc2\xae\xfa\xbe\x15\x08\x3a\xa8\x07\x61\x08\x5e\xee\x57\xf3\x65\xa3\x7c\xf0\x53\x64\x48\xac\x80\xb1\xa7\xa9\x18\x80\xca\x0c\xf1\xd1\xdb\x61\x6c\x03\xb4\xfc\x6f\x41\xfe\x3b\xdf\x47\x8e\x37\x3f\x93\x2e\x76\xa0\x27\xaf\xbc\xca\x47\x15\x37\x5e\x15\xbd\x2a\x2b\xbb\x77\x4b\x4b\xec\x3b\x16\xd9\xb7\x2a\xb3\xdf\xa4\xd0\xfe\x3e\x2a\x13\xb6\x24\x6b\x8f\x75\x4b\x3e\x0b\xc1\x4c\x30\x66\xe9\x45\xea\x6e\x55\x0b\xe5\xd5\xf8\xed\xeb\xf1\x5b\xdc\x7f\xcf\x39\x89\xaa\x2c\xb9\x25\x65\xf2\x4d\x65\xf8\xd5\x37\x92\xdc\xd2\xa2\xfb\x2d\x5d\x3f\x6a\x72\x4a\x37\xa8\xbb\x69\x94\xef\x1a\x09\x7f\x20\x19\x6f\x94\xf2\xd6\x72\xde\x20\xe9\xeb\xcb\x7a\x4e\xda\xbb\x55\xe9\x74\xaf\xd3\xc9\x4d\xd7\xae\x56\x47\x33\x72\x1d\x46\x56\xd4\xf5\xdc\x4b\x65\x4f\xb6\xb6\x27\xfb\xb3\x0e\xaa\x7d\x2b\x54\xf8\xb4\xa8\xf1\xb9\xc7\x7a\xae\xbf\xe0\x61\x56\x55\xeb\xd5\xb9\x26\x68\xfd\x62\xae\xbf\xad\x8e\xb0\x8b\x39\x24\x04\x39\x30\xa1\xfc\xef\xf7\x5e\xdc\xb5\x35\x96\x3f\x6c\xfd\x57\x2b\x56\x37\x16\x7e\x05\x90\x49\xd2\x14\x9a\x56\x08\x77\x9e\x1d\xaa\x92\x57\x57\x81\x2d\x73\xd7\x9e\x9a\x00\x64\x0d\xa1\x68\x08\x45\x43\x28\x1a\x42\xd1\x10\x8a\x86\x50\x34\x84\xa2\x21\x14\x0d\xa1\x3c\x1e\x84\xb2\x72\xd7\x24\x41\xfd\xb3\x8b\xa3\xc3\x33\x73\x7a\x76\x7c\x3c\x36\x0f\xf6\xa5\xfb\xa5\xf4\x68\x71\x4e\xae\xb4\x23\x77\x5e\x1e\xec\xe7\x23\x56\x85\x12\xe8\x51\x58\x3b\x59\x19\xc6\x53\x0b\x0e\x3c\x18\x3c\xd0\x02\x20\xe8\x00\x11\x34\x82\x04\x9b\xc0\x04\x05\xa0\xa0\x2b\xe6\xb3\x0e\xea\x53\x98\xb4\x2d\xf2\xa3\xd9\xbb\x39\x7b\x2b\x91\xa0\x7b\xc2\x82\x54\x34\xa8\xa8\xe6\x0a\x02\xa4\x41\x22\x0d\x12\x69\x90\x48\x83\x44\xea\x8d\x3c\x8d\x11\x69\x8c\x48\x63\x44\x1a\x23\xd2\x18\x91\xc6\x88\x34\x46\xa4\x31\x22\x8d\x11\xe9\x6b\x36\xfa\x9a\x8d\xbe\x66\xa3\xaf\xd9\xe8\x6b\x36\x1a\x41\xd1\x08\x8a\x46\x50\x6a\x0b\x00\x35\x7c\xa2\xe1\x13\x0d\x9f\x68\xf8\x44\xc3\x27\x1a\x3e\xd1\xf0\x89\x86\x4f\x34\x7c\xa2\xe1\x13\x0d\x9f\x68\xf8\x44\xc3\x27\x1a\x3e\xd1\xf0\x89\x86\x4f\x34\x7c\x52\xff\x7b\x43\x1a\xcb\x78\xe0\x58\xae\xf2\x75\x05\x31\x5b\x1e\xf0\x85\x05\x35\xc1\x9d\x78\x3d\x08\x85\x0e\x18\x20\xf0\x9c\x0e\x7f\x33\x5e\x50\xea\x0e\x86\xe2\xaf\x7f\x3f\x07\x43\x82\x31\x1b\x1a\x94\xba\x43\x2b\x62\x2e\x26\xe8\x3b\x74\x38\x41\x54\x47\x6d\x8f\x76\xd0\xfd\x19\xde\xb0\xfe\x38\xd1\x9b\x0e\xde\x9e\x5c\xf0\x56\x10\xf6\x26\x71\x6f\x76\xf4\xb7\x27\xf2\xea\x45\xf0\x0d\x63\xb8\x8a\xd7\xef\x3c\x99\xe8\xed\x87\xe7\x64\xcb\x28\xae\xd5\x4b\x7e\x5a\xbc\xea\xa7\xe5\x0b\x7f\xd6\x0b\xf3\xba\xfa\x8a\xca\x8f\x39\x3e\xf8\x4f\x39\x3e\x29\xe7\xb4\xd3\x7b\xeb\xb6\xa2\x6d\x2d\x5e\x6d\x97\xd5\xae\xfb\x7d\xad\xdd\x53\xf7\x03\xb5\x67\xf3\xa7\xf5\x6c\x7e\xb0\xe3\xf9\x89\x1f\x31\xbb\xd7\x91\xe7\x7d\x28\x79\x41\xf3\x33\x3a\xc8\xbd\x70\x39\x7e\x3b\x73\xc5\xdb\xdc\x94\x71\x1a\xde\xe8\xb6\xa1\x3d\xcd\xbe\xf0\x0a\xdc\xdd\x0d\x5f\x80\xe3\xc0\x11\x3f\x02\x0f\x4e\x19\x94\xb8\xcd\x8b\xe1\x72\x59\xd6\x2a\xf9\x7d\x5e\xa5\xe1\xff\x07\x00\x00\xff\xff\x47\x73\x89\xbc\x51\xaa\x00\x00")

// FileRemoteFileTfTmpl is "remote_file.tf.tmpl"
var FileRemoteFileTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x93\x3f\x8f\xd4\x30\x10\xc5\x7b\x7f\x8a\x91\x75\x05\x34\x16\x05\xa2\xbb\x02\x0e\xf1\xa7\x43\x47\x41\x81\x90\x95\x8b\x67\xd9\x11\xf1\x4c\xe4\x71\x76\x39\x59\xf9\xee\xc8\xce\x26\xac\x44\x90\xa0\x38\x57\xb1\xde\xbc\xf1\x9b\x9f\x32\xa5\x40\xc0\x03\x31\x82\x4d\x18\x25\xa3\x3f\xd0\x80\x16\xe6\xd9\x24\x54\x99\x52\x8f\x60\x79\x1a\x06\xbf\x5e\x2d\xd8\x31\xc9\x89\x94\x84\x7d\x29\xe0\xde\x63\xae\xe6\x45\xf5\xdc\xc5\x66\xf7\x9a\x71\xdc\xe4\x7a\xf1\x3c\xc5\x07\x4c\x55\xb4\x50\x0c\x40\xc0\x11\x39\xa8\x17\x86\x5b\xf8\x6a\x00\x00\x2c\x3d\x44\xdf\x4b\x1c\xa7\x8c\xfe\x14\x3d\xb1\xe6\x8e\x7b\x74\x7f\x7f\xc8\x1a\x80\x6f\xc6\x00\x6c\xa9\x30\x81\x5d\xa6\x28\xad\x69\x29\x40\x07\x70\x1f\x44\xb3\xfb\xa8\x5f\x88\x83\x9c\xb5\x4e\x08\xed\xf4\xc2\x8c\x7d\x26\xe1\x4b\x7d\x3d\x47\xd1\xdc\x3e\x6e\xc1\xde\x94\xff\x4f\xe5\x68\x3c\xbd\xf4\x5d\x08\x09\x55\x5b\xc6\xe5\xe4\xc7\x11\xd7\xbe\x67\xe2\x14\x7f\x4b\x93\x62\x5a\xa5\xd7\x21\x12\x93\xe6\xd4\x65\x49\x57\x6e\x8a\x28\x53\x6e\x25\xaf\x5e\x5c\x79\xc7\x4e\xf5\x2c\x29\x54\xa1\x86\xba\x93\x38\x62\xa6\x3a\x94\xbb\x17\xc9\x9f\x56\x7d\xde\xb2\xcc\x2b\x1b\x1c\x14\xff\x95\xc6\x13\x02\x59\x5a\xab\x1e\x77\x88\x2c\x5a\x12\xc9\x3b\x2c\x60\x07\x47\xa2\x53\x97\xd1\xff\xc0\xc7\x25\x6f\xfd\x1d\x9e\x55\x32\xc4\x01\x7f\xc2\x8d\x7b\x33\xd1\x10\xdc\x9d\xf0\x81\xbe\xd7\xb4\x83\x57\x3d\xfa\x2b\xdf\xb6\x07\xf6\xf9\x0e\x32\xae\x24\x4d\xbb\x5e\x96\x64\xcd\x51\x09\xdc\xb7\x4d\x7a\x47\x03\xba\xcf\x8b\xba\x62\x0f\xa8\x99\xb8\x6b\x78\xff\x2c\x7e\x7b\xa5\x2e\x8e\xd9\xcc\x66\x7b\xef\x57\x00\x00\x00\xff\xff\x62\xca\x3e\x64\xac\x03\x00\x00")

// FileScriptTfTmpl is "script.tf.tmpl"
var FileScriptTfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x55\x4d\x6b\xdc\x48\x10\xbd\xeb\x57\x14\x8d\x17\x6c\x76\x47\xb3\x86\x65\x0f\x06\x07\x6c\x13\x92\x5c\x82\x89\x0f\x39\xc4\xa6\x91\xd5\xa5\x99\x22\x52\x57\xa7\xab\x35\x1f\xc8\xfa\xef\xa1\xa5\x91\x46\x89\xc7\xbe\x84\xc4\x17\xeb\x30\xa8\xe7\xf5\x7b\x5d\x55\xaf\xab\xd4\x34\x60\xb0\x20\x8b\xa0\x24\xf7\xe4\x82\x82\xb6\x4d\x3c\x0a\xd7\x3e\x47\x50\xb6\x2e\x4b\x3d\x2c\x15\x28\xe7\x79\x45\x42\x6c\x75\xd3\x40\xfa\x0e\x03\xa8\x01\xd5\x36\xab\x30\xd2\xb5\x04\x74\x23\x1c\x17\xda\xd6\xd5\x3d\xfa\x08\x2a\x68\x12\x00\x83\x0e\xad\x11\xcd\x16\xce\xe1\x4b\x02\x00\xa0\x46\xc6\x1e\xec\x08\x09\xc0\x5d\x92\x00\x8c\x47\xa3\x07\x55\x50\x89\xbd\x14\x40\xd3\x00\x15\x90\xbe\x67\x09\xe9\x07\xf9\x4c\xd6\xf0\x5a\x62\x1a\x00\x3b\xf8\xc8\x48\x88\x0c\x38\x3b\x07\xe7\xc9\x86\x02\xd4\xd5\xd9\xbc\xcc\x0a\xf6\x0b\x9c\xf7\x99\xeb\xbf\x24\x75\x72\xaa\xe0\xf8\x71\xe0\x27\x7b\xb9\x9c\xad\xc5\x3c\x10\xdb\xdd\xf1\xf1\x59\xb2\x84\xee\xe5\x1c\xd4\x51\x43\xf7\x95\xce\xb9\x72\x75\x40\xbd\xaa\x34\x59\x09\x99\xcd\x31\x7d\xba\x64\x29\xb9\xd5\x7f\x3a\x33\xc6\xa3\x48\x97\x72\xff\x84\xad\xc3\x41\x77\x4d\xd6\x57\x7b\xa8\x16\xf4\x03\x74\x61\x2a\xb2\x24\xc1\x67\x81\xfd\x84\x4d\x15\x72\x1d\xba\x2d\xff\xff\x3b\xe1\xba\x4c\x64\xcd\xde\x44\x20\x06\x75\xc5\x95\xc3\x40\x31\xa9\xf4\x13\x73\xb8\x1e\xf0\x76\x8c\xa5\x4d\x76\x2f\xbb\x9b\x31\x1c\x9d\x89\x60\x90\x79\x54\xe9\x0c\x88\x3f\x31\xab\x78\x0f\xe2\x9f\x37\x5d\x6d\xd3\xcb\x4c\x70\xa2\x66\x50\x02\xd9\xac\xab\x62\x1f\xc3\x68\xd1\xb0\xa9\x69\x00\xcb\x8e\xf4\xbc\x8f\xf3\x50\xb9\x03\x4e\xca\xf2\xd7\x8c\xfc\x8d\x5e\xf6\xd2\x22\xcb\x03\x66\xf6\x98\x67\x0e\x07\x6c\x84\x03\x4e\x7a\x5a\x65\x01\xf5\x57\xdc\xf6\xf1\xc6\xf2\x1c\xc7\x82\x92\x35\xb8\x81\xf4\xb2\xa6\xd2\xa4\x57\x6c\x0b\x5a\xc4\x60\x4b\x2d\xb2\xd4\x13\x9a\xee\x5b\xa9\x6d\xd5\xc9\xc4\xec\x3f\xef\xb5\x35\xbd\x35\xed\xa3\x5e\xf7\x58\x71\xc0\x19\x6e\x30\x7f\x6d\xf9\x97\x68\x79\xb2\x65\xfc\x42\x0c\x93\xba\x9b\xd6\x8e\xd7\xe8\x65\x89\x65\x09\xb3\x8f\x7c\xed\xb9\x2b\xf6\xec\xed\x06\xf3\x3a\x8a\x5e\x73\x49\xf9\x16\x2e\xb7\xf1\x60\xf8\xc9\xf4\x7f\xf6\x3a\xbe\x32\xe4\x61\x2e\x30\xff\x06\xb7\x13\x87\x6e\x15\x3c\x3c\xc0\x0a\xfd\xe9\x1b\x5b\x97\x23\xe3\xee\x75\x36\xbc\xf8\x6c\x78\xe6\x56\xe4\xcb\x8a\x0d\xfc\xbd\x79\xc6\xf0\x1f\x91\xc9\xd8\xb8\xf0\x8b\x9b\xe0\xc9\x2e\x1e\xdd\x10\x98\x15\x4f\x09\xde\x1d\x98\x1f\x6d\x32\xae\xbf\x07\x00\x00\xff\xff\x5e\x94\x8f\xa3\xdf\x08\x00\x00")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileCommandTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "command.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileDNSRecordTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "dns_record.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileInfraTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "infra.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileOriginalInfraTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "original_infra.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileRemoteFileTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "remote_file.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileScriptTfTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "script.tf.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}