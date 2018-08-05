// Code generated by fileb0x at "2018-08-05 13:32:48.082915523 -0700 PDT m=+0.008840493" from config file "assets.toml" DO NOT EDIT.
// modification hash(d88bf21972afcccf0761832b3b9bc9b6.e5b9c5ef4c0b7aef8593382d0449dfd6)

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

// FileAmiLaforgeTmpl is "ami.laforge.tmpl"
var FileAmiLaforgeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xa4\x53\x4d\x6f\xdb\x3a\x10\xbc\xeb\x57\xcc\x33\x7c\x78\xcf\xc8\xb3\xef\x05\x72\x30\x12\x27\x30\x90\x2f\xa4\x09\xd0\xeb\x4a\x5c\x59\xdb\x48\x4b\x87\xa4\x9c\xa6\x86\xfe\x7b\x41\xd2\x72\x1c\xa0\x40\x0f\xd5\x45\x20\x39\xfb\x31\xb3\xb3\x8b\x59\x51\x00\xab\x6f\xcb\xdb\x87\x9b\x15\x6e\x96\x57\xf7\x8f\xd7\x2b\x2c\x6f\xd7\xb8\xb8\xbf\xbb\x5a\x5f\x3f\x3f\x2e\x9f\xd6\xf7\x77\x45\x31\x5b\x14\xc5\x62\x06\xc3\x55\x4b\x8e\x41\xe8\x55\x5e\x7b\xc6\xfa\x12\xa2\x08\x0d\xc3\x96\xdf\xb9\x0a\x10\x95\x20\xd4\xca\x4f\x0a\x62\x15\xb3\x45\x41\x9d\x60\xb2\xdf\x63\x3a\x5f\x5f\x62\x18\x26\xd8\xc7\x9a\x8b\x19\x08\x9d\x75\x8c\xa6\xef\x48\xe1\x98\x0c\x95\x2d\x43\xa9\x63\x48\x8d\x77\xdb\x63\xeb\xb8\x66\x77\x06\x52\x03\x82\x61\x5f\x39\xd9\xa6\xbc\xff\xda\xf4\xa7\xf6\xbf\x58\x02\x39\xea\xfc\x50\xe7\x2e\x1e\x86\x61\x52\xe0\x53\xcc\xf8\x7c\x79\x72\x17\x51\xb9\x9b\xc8\x61\xeb\xec\x4e\x0c\x3b\x50\x08\x4e\xca\x3e\x30\xc4\xe3\xad\xa1\x80\x96\x6a\xeb\x36\x8c\xde\xb3\x47\xb0\x30\x1c\xd8\x75\xa2\x8c\x5e\x63\x44\x02\x55\x56\x8d\xc4\xbc\xbe\x40\xfc\x42\x23\x3e\xa9\x59\x91\xa2\x4c\xc1\x26\xea\x45\xd8\xc8\x8e\x15\xac\x3b\x71\x56\x3b\xd6\x30\xc7\x57\xe6\xd4\x84\xb1\x95\x47\x6d\x1d\x24\x78\x64\x9a\x7e\x9e\x59\x1e\xfb\x1b\xa9\x3c\x8c\x17\x1f\x3c\x7a\xcf\x2e\xa9\xf1\x26\x6d\x1b\x8b\xc6\x9c\x25\x79\x3e\x79\x69\xa4\x6a\x8e\x8c\x94\xd9\x24\x4a\x2f\x6a\xdf\xe2\xbf\xb2\x5a\xcb\xa6\x77\x39\xb4\xb1\x3e\xe4\xea\xc7\xf8\xb1\xfa\xf3\x78\xf1\x51\x3d\xc7\xa2\xb1\xad\xf1\x20\xec\xc8\x49\x9a\xaa\xe7\x00\x5b\x8f\x74\x10\xa2\x5a\xd1\x48\xac\x41\x1c\xb7\xef\x30\xbc\x65\x35\xac\x21\x2b\x67\xb3\xa9\x76\xd4\xf6\x1c\x03\xe3\x61\x32\xd2\x9f\x9c\xcc\xc7\x70\x2d\xca\x06\x54\xda\x1d\xe7\x3e\x0f\x3d\x9c\x63\x9f\x72\xed\xf7\x70\xa4\x1b\xc6\xf4\x85\xdf\xcf\x30\xdd\x51\x8b\x2f\xe7\x98\xce\x2f\x32\xee\xff\x61\x18\x71\x11\x81\x61\x18\x09\x46\x64\xb6\x51\x7a\x65\x35\x07\xf0\x30\x7a\x86\x36\x1e\xfc\x63\xeb\xd8\x7b\x6c\x58\xd9\x51\x0b\xd1\xda\xba\x2e\xdb\x3f\x2b\x3d\x4e\x22\x29\x11\x2c\x4a\xce\x24\x5f\x7b\x76\xc2\x26\x92\x3d\x2c\x51\xdd\x87\xa8\x7b\xa2\x91\x92\xff\x91\xc4\x53\x44\xfd\x05\x85\x8e\x44\x03\x89\xb2\x8b\x56\xaf\x6c\xb7\x6d\x39\xc4\x81\x8c\x0b\x76\x86\xb2\x0f\x50\x1b\x18\x81\xa9\x83\x75\x1b\xd2\xc3\x7a\xff\x93\x3b\x3d\xc9\x91\x8d\x71\x7b\xbc\xf8\xd8\xf9\xd8\xc1\x27\xf7\x9c\x80\x4e\x16\x16\xe0\x8e\xa4\xfd\x1d\x6a\x95\x1e\x32\x6c\x28\x86\x5f\x01\x00\x00\xff\xff\x75\xa7\xab\xcf\xbc\x04\x00\x00")

// FileCommandLaforgeTmpl is "command.laforge.tmpl"
var FileCommandLaforgeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x6c\x54\x5d\x6b\xe3\x46\x14\x7d\xf7\xaf\x38\x6b\x02\xeb\x04\x37\xe9\x73\xc1\x0f\x26\xf6\x2e\x86\x8d\x1d\x76\xb7\x50\x28\x25\x5c\x4b\x57\xf2\xdd\x8c\x66\xd4\xf9\x88\xd7\x35\xfa\xef\x65\x46\x23\xc5\xa6\x7d\x30\xd6\xcc\xfd\x3a\x73\xcf\xb9\xf7\xe1\x6e\x32\x01\xd6\x7f\x2c\x9f\x9e\xbf\xac\xf1\x65\xf9\x69\xf7\xf5\xf3\x1a\x8f\xbb\xa7\xa7\xe5\x76\x85\xc7\xdd\xf6\xd3\xe6\xf3\xef\x5f\x97\xdf\x37\xbb\xed\x64\x72\xf7\x30\x99\x3c\xdc\xa1\xe4\x42\x91\x65\x10\x82\x96\xbf\x03\x63\xb3\x82\x68\xf8\x03\xc3\xec\x7f\x70\xe1\x21\x5a\xbc\x90\x92\x7f\xc8\x8b\xd1\xb8\x7b\x98\x14\xa6\x69\x48\x97\x98\x9e\xcf\xb8\xb9\xdf\xac\xd0\x75\x53\x9c\x63\xed\x87\x3b\x10\x1a\x63\x19\x87\xd0\x90\x86\x65\x2a\x69\xaf\x18\x9a\x1a\x86\x54\x38\x99\x80\xd6\x72\xc5\x76\x8e\x98\x82\x50\xb2\x2b\xac\xb4\x29\xf7\xcc\xa4\x7f\x52\xb7\xb1\x0c\xfa\xa8\x45\xae\xb3\x8d\x87\xae\x9b\x4e\x70\x15\x33\x98\x57\x17\x77\xd1\xab\x47\x13\xdf\xd1\x5a\x53\x5b\x6a\x20\x2e\x1d\xe3\x6f\x4f\x8e\x31\x3c\xc3\x1b\xd8\xa0\x31\xd3\x06\x64\xeb\xd0\xb0\xf6\x2e\x03\x18\x42\x87\x22\xcf\xf9\xfc\x5e\xc0\x85\xb6\x55\xa7\xf7\x40\x1c\xd8\x72\x1f\x4c\xb6\x76\x58\xe0\xcf\x09\x00\x9c\xcf\xb0\xa4\x6b\xc6\xcd\xcb\x1c\x37\x64\x6b\xfc\xb6\xc0\xcd\xfd\x32\xfa\xfc\xd2\x75\xc9\x27\xd5\x88\xa6\xae\x9b\xce\x87\x28\xd6\x65\x76\xf8\x2b\x97\x94\x0a\x52\x6b\x63\xf9\x85\xad\x35\xd6\xc5\x87\x69\xe3\xe1\x6d\xe0\x39\x14\x55\xc6\xd6\x8c\xa3\x28\x85\x03\x29\x1f\x1f\xf1\x26\x4e\x8c\x16\x5d\x23\xb4\x46\x83\x34\x52\x28\x2c\xbb\xd6\x68\x97\x01\x5f\x67\x5d\xa0\xa7\x37\x5d\xae\xfb\xbb\xae\xcb\x18\x0a\x63\x54\x69\x8e\x1a\xae\xe5\x42\x2a\x61\x07\x42\x11\x9c\x37\x0d\x8e\x24\x1e\x5e\x1a\x86\x3b\x98\xa0\xca\x44\xfa\x51\xdc\x21\x76\x3a\x19\xa9\xf2\x6c\x13\x11\x03\x07\xfc\x53\xbc\xc3\xac\xe4\x8a\x82\xf2\x58\xe0\xd7\xcc\xc0\x58\x28\xc3\x79\x1c\xce\x23\x94\x52\x5c\x94\x58\x09\x27\x4d\xa2\x42\x29\x73\x8c\x54\x8b\xbb\xa4\x78\xcf\x68\xc9\x39\x2e\x61\xde\xd8\x82\xc9\x89\x3a\x45\xa9\xd3\xe8\x55\x1c\x48\xf4\x1c\x5c\x55\x5c\x78\x79\x63\x75\x42\x43\xaf\xb1\x6b\x11\x33\xb6\xbb\xdd\x73\x0f\x6a\x2c\x99\x41\xad\x86\xf3\x08\xea\x47\x70\x1e\x4a\x5e\x23\x0f\xfe\x00\xe3\x0f\xf1\xc1\xa7\x96\xdd\x3c\xb5\xa3\xa0\xa1\x75\xa7\x04\x40\x57\x4a\x0a\x0f\xe7\x2d\x79\xae\x4f\xbd\x8a\x66\x71\x7c\x52\x9b\x4a\x53\xb8\x0f\xb9\x23\x46\xbf\x8c\xfe\xe7\xa4\x92\xd2\x8c\x02\xdd\xe9\xc7\x6c\xbb\x5f\x99\x3c\x2d\x00\xb5\x6d\x94\x51\x46\x7b\xe1\xb3\xec\x0d\x49\x5e\xdd\xc5\xc8\x88\xc1\x5e\x99\xe2\x15\x8a\xbd\x4b\x80\x4b\xae\x44\xf3\x40\xb1\x09\xbe\x0d\x1e\xb3\xd8\x35\xd1\xf1\x53\xf4\xd0\xc6\x8f\x0e\x05\x39\xbe\x45\x65\x6c\xcf\x02\xff\xe4\x22\x0c\xcb\x03\x31\x79\x8f\x3b\x8e\x8f\x2f\x25\x72\xdb\x92\x4f\xf2\x20\x54\xa2\x38\x7e\x15\x14\x93\x7a\x83\x6f\xdf\x57\x9b\x6d\x4e\x36\x0e\xf3\xc7\xff\x64\xc5\x98\x2b\xaf\xa5\xdd\xfd\xb7\x74\x91\x67\x75\x28\x67\x82\x9f\xc7\x7f\xb6\x16\x71\xf1\xed\x15\xe9\x57\xcc\xa6\xd3\x5b\xec\x4f\xc8\x02\xec\x97\x53\x1a\x21\x65\xea\x08\x27\x6d\x8d\x20\xaa\x44\x29\x36\x59\xfb\x5c\x28\x83\x8d\x0a\xb9\x1a\xb2\x77\x40\xd1\xe3\x1a\x51\xbc\x19\x78\xc9\x30\xae\x1d\xe2\x4d\xef\x30\x30\xd2\x90\x68\x4f\xa2\xd9\xa2\xd7\x74\xab\xd8\x47\x75\x0e\xeb\x72\x8e\x7d\xf0\x71\x07\x30\x3c\x53\x03\x63\x6b\xd2\x79\x61\x7f\xe8\xd1\x5c\xe4\xe8\x8b\x3d\x8d\x17\xef\x1b\x3c\x62\xba\xda\xb9\x17\x4e\x17\xeb\x17\xe0\x86\x44\xfd\x9f\xd7\x3a\x19\x06\xf8\xdd\xbf\x01\x00\x00\xff\xff\x24\xf8\x0e\x67\x92\x06\x00\x00")

// FileCompetitionLaforgeTmpl is "competition.laforge.tmpl"
var FileCompetitionLaforgeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x9c\x93\x5d\x8b\xa3\x30\x14\x86\xef\xf3\x2b\x0e\xd2\x2b\x99\xad\xf7\x0b\xbd\x70\xaa\x1d\x64\x1d\x2d\xea\x5c\x2c\xcb\x52\x5c\x3d\x75\xa4\x35\x91\x24\x33\x43\x91\xfc\xf7\x25\x9a\xba\xad\xec\x87\xee\x65\xce\xfb\x9c\xf7\x21\x11\x1d\x9b\x10\x00\xcf\xdf\xb9\x2f\x61\x06\xa1\xbb\x8b\x93\x27\x1f\x1e\xdd\xd4\x87\x6d\xfc\xbc\xf7\xb3\x20\x0b\xe2\x08\xb6\x71\xb4\x0b\x9e\x5e\x12\x57\x9f\x08\xb1\x1d\x42\x88\x63\x43\xf0\xbc\x8f\x93\x0c\xdc\x30\x84\xf5\x39\x3f\x32\x5e\x21\xec\x82\xd0\x4f\x21\x88\xa0\x60\xf4\x58\x57\x0e\x78\x41\xe2\x6f\xb3\x38\xf9\x0a\xb6\x43\x6a\x5a\x9c\xdf\x4a\x84\x8e\x00\xb4\xb9\x7c\x85\x0d\x58\x6b\xc7\xa0\xf6\xb5\xc4\x22\xea\xdf\xfd\xa2\xe0\x75\x2b\xc5\x2c\xc1\x95\x5d\x66\x28\x58\xd3\xe4\xb4\x14\x33\xef\x60\xe0\x65\x8e\x57\x26\x66\xde\x61\x20\x97\xb5\x53\x94\x1f\x8c\x9f\xe6\x09\x46\x78\x99\xa3\x2e\x91\xca\x5a\xd6\x38\xcf\x72\x83\x2f\xf3\x1c\xeb\xf3\x4c\xc5\x40\xde\xb7\x93\x82\x35\x2d\x6a\x31\xa3\x60\x75\x1d\xac\xd6\x81\x07\x4a\x59\xfd\x2e\x67\x4c\x1e\xda\x5c\x88\x0f\xc6\x4b\x5d\xd2\x03\x09\x63\x72\x7f\x1d\x2a\x65\xe9\x5f\xa5\xa4\xc2\xa4\x5e\x94\xde\x56\x00\xc8\x4b\x8b\xe3\xae\x4e\x33\x3d\x30\x7b\xc6\x51\xb2\x26\xaf\xe9\x1d\xa5\x2d\xde\x30\x1e\xd9\x92\x8a\x83\x40\xfe\x8e\x5c\xc0\x06\xbe\xf5\x33\x80\xae\x03\x9e\xd3\x0a\x61\x75\x78\x80\x95\xe0\xef\xf0\x79\x63\x4a\xbc\x28\x4d\x0d\xff\x49\x29\xc3\xf7\x0e\x8d\x29\x65\x3d\xfc\xea\x40\x5a\x8e\xd0\xf7\xc1\x47\x65\xbb\xcc\x17\x65\xfb\xff\xf4\x01\x28\xed\xe4\xd8\x30\x89\xd7\x87\xee\x0f\x7f\x79\x4d\x03\x4c\x1e\x14\x2b\xfd\x31\x27\x4c\x32\x4c\x47\xea\x84\x97\x29\xf2\x05\x2f\x7d\xae\x63\x81\x05\x47\x39\x25\xd2\x61\x3a\x96\x08\x99\x4b\x3c\xfc\x78\x2b\x4e\xbf\x61\x75\xf6\x38\x44\x63\xab\x64\x3c\xaf\xfe\xbc\xd1\xa7\x77\x3b\x8a\xa8\x9f\x01\x00\x00\xff\xff\x47\xb8\x82\x1a\x8c\x05\x00\x00")

// FileDNSRecordLaforgeTmpl is "dns_record.laforge.tmpl"
var FileDNSRecordLaforgeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x4c\x51\x5d\x6f\x9b\x3c\x14\xbe\xf7\xaf\x78\x5e\xf4\x5e\xb4\x51\x1b\xb6\x1f\xd0\x49\x28\xa4\x55\xa4\x16\xaa\x2c\x9b\x7a\x57\x19\xfb\x50\xdc\x1a\x9b\xd9\x26\x8c\x46\xfc\xf7\x89\x8f\x64\xbb\xe3\x9c\xf3\xf0\x7c\x39\x5e\x31\x06\x6c\x5f\x92\xa7\xe7\xc7\x2d\x1e\x93\xfb\x7c\xff\xb0\x45\x9a\x7d\xc7\x7e\xbb\xc9\xf7\x29\x36\x79\x76\xbf\x7b\xf8\xb1\x4f\x0e\xbb\x3c\x63\x6c\x15\x33\x16\xaf\x20\x49\x68\xee\x08\x1c\xad\x51\xbf\x5a\xc2\x2e\x85\x32\x08\x15\xc1\x16\xef\x24\x02\x94\x51\x41\x71\xad\x3e\x79\x50\xd6\x60\x15\x33\x69\xfc\xab\x23\x61\x9d\x44\x74\x3a\xe1\xff\xf5\x2e\xc5\x30\x44\x38\xb1\xd1\x42\xbc\x62\x00\x10\xfa\x86\x70\x0b\x49\x5e\x38\x55\x90\x47\x12\x6f\xb2\xe4\x69\x1b\x1f\x5e\x0e\x31\x05\xb1\x9e\x50\x86\xd7\x23\x6a\x94\x93\xc6\x23\xfa\xa0\x3e\x42\x69\xdd\xb4\xf9\xb4\x86\x70\x35\xfd\x85\xae\xeb\xd6\x05\x77\x6b\x61\x6b\xdc\x7e\x43\x69\xed\x97\xaf\xe7\xc5\xcd\x4c\xd3\xd9\x56\x4b\xdc\x21\xea\xba\x2e\xba\x9e\xe8\x8f\x5c\xb7\x67\xfe\xf9\xdb\x96\x17\x31\x32\xc1\xf5\xb8\x9a\xa8\xf0\xa6\x8e\x34\xc7\xe6\x85\x3d\x12\xe8\x37\xaf\x1b\x4d\x23\xcd\xd8\xd4\x92\xe7\x6e\x09\x7c\x18\x87\x61\x88\xd8\x92\xe0\xbc\xcf\xc6\x61\xde\xcf\x72\xe7\xc3\xcf\x69\x1a\x2f\x53\x43\x90\xca\xf3\x42\x93\x84\x57\x75\xa3\x7b\x70\xad\x6d\xe7\x11\x2a\xe5\x21\x6c\x5d\x73\x23\x11\x2c\x0a\x42\xc3\xbd\x27\x09\x7b\x24\x07\xe2\x5e\xe9\x7e\x7c\x1f\x7e\x41\x89\x8a\x2b\x73\x03\x2a\x4b\x12\x41\x1d\x49\xf7\xa8\xf9\x87\x32\x6f\x50\x01\x1c\x59\x9e\x3f\x8f\x09\xf0\x57\xf2\x0e\x93\xa5\xf4\x3c\x0f\xc3\x62\xea\xbd\xf5\x01\x5a\x7d\x10\x3a\x15\x2a\xd8\x50\x91\x9b\x72\xfb\x1b\xf4\xb6\x85\xe0\x06\xbe\x21\xa1\xca\x7e\x32\x60\x4a\xad\x44\x80\x0f\x8e\x07\x7a\xeb\x51\x91\x23\x5c\x39\xe2\x72\xee\xd8\x0a\xff\xdf\xf5\x2c\x6e\xcd\xeb\x05\x7f\x9a\x9e\x46\xda\x4b\x39\xb9\xd9\x2c\xb7\x75\x6a\x97\xfa\x00\xde\x34\x64\x2e\x6e\xff\xc1\x24\xf3\x61\x18\x18\x30\xb0\xe1\x4f\x00\x00\x00\xff\xff\x04\x0c\x2f\x7f\xf9\x02\x00\x00")

// FileEnvironmentLaforgeTmpl is "environment.laforge.tmpl"
var FileEnvironmentLaforgeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xa4\x54\xc1\x6e\xe3\x36\x14\xbc\xeb\x2b\x66\x03\x1f\xb6\x86\x37\xbe\x17\x70\x01\xd7\x76\xb2\x42\x13\xa9\x70\x9c\x2d\x8a\xa2\x08\x18\xe9\x45\x62\x4d\x91\x2e\x49\xc9\x50\x0d\xfd\x7b\x41\x52\x92\xed\xa6\xc0\x1e\x72\xb3\xf8\xe6\x71\x86\xe3\x79\x6f\x3e\x8d\x22\x60\xbd\xb9\x5b\x3e\x3f\xec\xf0\xb0\xbc\x4b\xb7\xf7\x1b\xac\xd2\xc7\x5f\x37\xbb\x78\x17\xa7\x09\x36\xc9\xb7\x78\x9b\x26\x8f\x9b\x64\x87\x55\x9a\xdc\xc5\xf7\xcf\xdb\xa5\xab\x44\xd1\x74\x1e\x45\xf3\x29\x48\x36\x5c\x2b\x59\x91\xb4\xa8\x25\xff\xbb\x26\xc4\x6b\x4c\xe7\xd1\x65\xe1\xe6\x74\xc2\xe4\x36\x5e\xa3\xeb\x6e\x70\x72\xa4\xf3\x69\x04\xe0\xaa\x3b\x27\xcb\xb8\x30\x11\x5c\x37\x20\x59\x45\x58\xf4\xad\x89\xfb\xe8\xba\x9b\x08\xc8\xc9\x64\x9a\x1f\x2c\x57\x72\x2c\xaf\x2f\xce\x1c\xea\x4c\x60\x4b\x6e\x60\xc8\x5a\x2e\x0b\xc7\x40\xba\xe2\x92\x0c\x8e\x25\xb3\xb0\xed\x81\xa0\xde\x70\xd0\xaa\xe1\x86\x2b\xe9\x40\x47\x2e\x04\x54\x96\xd5\x1a\x9f\xd9\xd1\xcc\xd0\xb0\x42\x33\x69\x67\x90\xcc\xf2\x86\x7e\x18\x04\xfa\xee\x41\xc1\xce\x7d\x5c\x53\xbf\x32\x43\x39\x94\x44\xab\x6a\xed\xd1\x33\xf7\x13\x82\xef\x49\xb4\x81\xa7\x64\x0d\xf9\xd2\x17\x73\xa0\x8c\xbf\xf1\x0c\x99\x92\x6f\xbc\xa8\x35\xf3\xaf\x69\x98\xa8\xc9\xa0\x24\x4d\x03\x6f\x00\x60\x81\x93\xa7\x39\x9d\xa0\x99\x2c\x08\x93\x3d\xb5\x33\x4c\x1a\x26\xf0\xe3\x02\x93\xdb\x55\xc0\x7d\xe9\xba\x01\xe7\x10\xe8\xba\x41\xb4\x43\x06\x4f\x7d\x95\x64\xde\x83\xbb\xa8\x7f\x86\xa3\x9c\x22\x4e\x56\x0f\xcf\xeb\x38\xb9\x47\xb2\xd9\xfd\x96\x6e\x7f\x79\xc2\x32\x59\xe3\x6b\xfa\xb4\x7b\x42\x9c\xe0\xf7\xf4\x79\x7b\x95\x94\xa9\x33\xe0\x27\x18\x5e\x1d\x44\x8b\x9c\x32\xc1\x34\xc1\x96\x64\x08\x5c\x66\xa2\xce\x29\x7f\x91\x64\x8f\x4a\xef\xf1\x2a\x54\xb6\x37\x38\x72\x5b\xf2\xde\xab\xcb\x54\x9c\x82\xfa\x6b\x57\xac\x1a\xee\x41\x7f\x8d\x01\x93\x39\x4a\x65\xac\xb9\x75\xb2\xcf\xae\x48\xb2\x33\x4c\x7c\x25\xf8\xf2\x55\x19\xfb\x73\x9b\xf4\xf4\xe1\xc5\xef\x54\x79\x83\x24\xd9\x3e\xb1\x4e\xc2\x88\x09\x77\x2d\xf0\x87\x3f\xbe\xe4\x7a\x71\x4c\x9e\x25\x60\x06\xef\x11\xee\x2b\xfb\x11\x98\x9d\x1b\xcf\xa6\x03\x7f\x06\xeb\xaf\xce\xc3\x9f\x7e\x91\x2a\x97\xa0\x8c\x49\x58\x56\xbc\x37\x8b\x19\x1c\x49\x88\x31\xa1\xac\x30\xdf\xcf\xc9\x37\xa6\xcd\x87\x52\xe2\x2a\xce\xfd\x9c\xde\xb8\x24\x30\xd9\xa2\x61\x9a\xb3\x57\x41\xc6\xeb\xad\x78\x51\xda\x41\x54\xe3\xe8\xbe\x2b\x6a\xe7\xa4\x7f\x54\xd4\x5f\xb5\xb1\x7e\xdc\x7c\xba\xa0\x6c\x49\x61\x10\xcd\x6c\xf4\x31\xcc\x5d\x0b\xe6\x23\x26\x78\x66\x61\xac\x66\x96\x8a\xd6\x4f\x1d\x3e\x6b\x62\xb9\x0b\x2f\x72\x95\x99\x4f\xe3\xf8\x2b\xf9\x32\x76\x84\xc7\xe4\x6a\x5c\x07\xa9\x5c\xf5\xb5\xdb\xb5\x1a\x95\xb2\xc3\xc1\x29\x5d\xe0\xbf\x98\x65\x28\xbc\x7f\x41\xc5\xb8\xb4\x8c\x4b\xd2\xe0\x06\x99\xaa\x0e\x82\xac\xdb\x1e\xca\x2f\x3b\x26\x66\x78\xad\x2d\xa4\xb2\x04\x4b\xac\x82\xd2\x05\x93\xfc\x1f\x3f\x26\x9f\x06\xa9\x17\xb7\x04\x79\x8f\xe3\xc1\x79\x29\x3b\xba\xab\x9d\x7b\x01\xba\x58\xbf\x00\x55\x8c\x8b\xff\x43\x6d\x7c\x21\xc0\xba\xa8\xfb\x37\x00\x00\xff\xff\x9e\x66\x95\x08\x5e\x06\x00\x00")

// FileHostLaforgeTmpl is "host.laforge.tmpl"
var FileHostLaforgeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xa4\x56\x5d\x6b\x1b\x47\x14\x7d\xd7\xaf\x38\x31\x06\x3b\x66\x6d\xd3\x3e\xf4\x21\xe0\x82\x1b\x39\x89\x20\xb1\xdc\xd8\x29\x85\x52\xc4\x68\xe7\xae\x34\xd1\xec\xcc\x66\x3e\x24\x2b\x46\xff\xbd\xdc\x99\xd9\x95\x45\x03\x69\xe8\x93\xb4\x77\xce\x9e\x7b\xee\xe7\xce\xe5\xd9\x68\x04\xdc\xfc\x79\xfd\xe1\xee\xfd\x0d\xde\x5f\xbf\x99\x7e\x7c\x7b\x83\x77\xd3\xfb\x07\xbc\x9e\xde\xbe\x99\xbc\xfd\xf4\xf1\xfa\x61\x32\xbd\x1d\x8d\xce\x2e\x47\xa3\xcb\x33\x48\xaa\xb5\x70\x04\x81\x68\xd4\x97\x48\x98\x8c\xa1\x0c\xc2\x92\x60\xe7\x9f\xa9\x0e\x50\x46\x05\x25\xb4\xfa\x2a\x82\xb2\x06\x67\x97\xa3\xa5\xf5\x01\x47\x4f\x4f\x38\xbe\x98\x8c\xb1\xdb\x1d\xe1\x89\xbd\x5e\x9e\x81\x4f\x8c\x68\x09\xa7\xc6\xf2\x9b\xb5\x8e\x52\x99\x05\xde\xfc\x3e\xbe\x7d\xc9\xaf\x62\x0f\xb9\x2a\x14\xef\x7a\xc3\x6e\x77\x54\x68\x1a\x45\x5a\x22\x58\x48\xf2\xb5\x53\x73\x42\x58\x2a\x9f\x5e\xcd\x24\xd9\xde\x25\x41\x3d\xcf\xf8\x99\x6d\x4f\xe5\xe8\x4b\x54\x8e\x64\x85\x05\x19\x72\xaa\x86\x8f\x5d\x67\x5d\x20\x09\xdb\x91\x13\x81\xf5\xf9\xad\x0f\xd4\x7a\x9c\xc6\x79\x34\x21\x56\xa8\xc9\x04\xeb\x2b\xac\x84\x56\x15\x36\xca\xfc\xbc\xfa\xe9\x97\x0a\x14\xea\x8b\x12\x87\xf5\x83\xe7\xe9\xfd\x33\x87\x36\x29\x10\x9a\xdf\xd2\x1a\x76\x4d\xce\x29\x49\x29\xa3\x92\x1a\x11\x75\x80\x6a\xc5\x82\x9f\x02\xb9\x56\x19\x92\x98\x6f\x31\xbd\xc7\xf9\xaf\xf0\x4b\x1b\xb5\xc4\x9c\xb8\x08\x8d\x68\x95\xde\xc2\x36\xec\x2c\x39\x15\xad\x1a\xbc\x5e\x7f\x98\x3c\x73\xcb\xf4\x5a\xf8\x00\x5b\x07\x0a\xfc\x0e\x5b\x38\x63\x27\x1e\x93\x3b\x08\x29\x1d\xf9\x42\xc3\xc0\x59\x06\x5e\x21\x91\xbd\x17\x3e\x4c\x93\x61\xb7\x2b\x8c\xbd\x58\xaf\xbe\x12\x94\x87\x6f\x85\xd6\x15\x6c\x58\x92\x2b\x51\x7a\xd4\xc2\xb0\xd8\x96\xa4\x8a\x6d\x05\x2d\xdc\x82\x2a\x3c\xa6\xdf\xec\x4b\x19\x1f\x84\xa9\x69\x96\x78\x7a\xf1\x93\x62\xbd\x67\xe3\x3e\x0a\x4f\x21\x27\x4a\xf9\x55\x71\x6c\xf0\xf6\xb7\x52\x74\x36\x3e\x8d\x00\xa0\x70\xe5\xba\x2b\xbf\xba\x28\x3c\x23\xa0\x97\x7f\x90\x78\x67\x6d\x40\x27\xbc\xdf\x58\x27\xa1\x1a\x6c\x6d\xc4\x46\x98\xc0\x4d\x96\xcb\x59\xe0\xb3\x01\x35\x54\xb7\x9c\xdc\xf5\x07\x83\xdc\x24\x85\xe9\x1b\xab\xb5\xdd\x70\x23\x69\xe5\x83\x07\x0f\x54\x6e\x7e\x92\x90\xd4\x91\x91\x64\x6a\x45\x9e\xeb\xb2\x16\x4e\xd9\xe8\x11\xb6\x1d\x79\x9c\x33\xc1\x16\x8e\x1a\x72\x64\x6a\xf2\x89\x70\x32\x2e\x33\x90\xea\x68\x3d\x65\x74\x72\xb8\x16\x5a\xc9\xfc\xfc\x2a\x19\x80\x73\x48\xe3\x67\x8e\x6a\xeb\xe4\x60\x6a\x94\xa6\xe1\xa1\xb6\x6d\x2b\xcc\xfe\x30\x8f\xca\x10\x00\x3d\x52\x1d\xd3\xe4\x58\x27\xc9\xe5\xd6\x9d\x13\x84\x4f\x99\x5a\x8a\x35\xc1\x77\x54\xab\x46\x91\xc4\x92\x1c\x33\xa7\xbc\x75\xce\xae\x95\x57\xd6\xcc\x7c\xa0\x8e\x67\xe2\xaf\xc4\xfa\xf4\x04\x27\xcc\x82\x70\x3c\xab\x70\xfc\x88\x57\x57\x38\xbe\xb8\xeb\xc1\xf7\x09\x7b\x9e\x0a\x86\x9c\xe7\x47\xce\x6b\xd5\xbf\x4b\x46\x96\xe3\xbf\x4b\x3d\x59\x07\x77\x9b\xed\xc8\xa0\x51\x8e\x36\x42\x6b\xb8\xa8\x39\x65\x36\x57\x99\xfc\xd0\xf9\x9d\x95\x27\x1e\xca\x34\x4e\xf8\xe0\x62\x1d\xa2\xa3\x24\x3c\xcb\xa6\xc7\xce\x7a\x92\xb3\x50\x77\x33\xde\x06\xdf\x11\x7e\x93\xe1\x0f\xaf\xef\xee\x12\xf8\xbf\x2b\xef\x1d\x45\xf9\x23\x8e\x3e\x8d\x7f\xd8\xd1\xe5\x19\x3e\x47\x1f\xa0\xd5\x8a\xb0\x51\x61\x59\x26\x35\x35\x4a\x35\xa4\x2f\x57\x71\x0b\x81\xda\x9a\x46\xab\x3a\xc0\x07\x27\x02\x2d\xb6\x39\x3d\xa7\x8e\x84\xcc\x43\x68\x6b\xff\xa2\xdf\x76\x66\x36\xe0\xf3\x08\x4a\xbb\x9f\x10\xf3\xba\x9c\x5d\x8c\x6d\x1a\x0f\x06\x88\x8e\xfb\xbe\x9f\xd2\x67\x98\xeb\x7c\x70\x30\xad\xec\x4f\x59\xcc\xb5\xad\x57\xd0\x14\x72\xdf\x49\x6a\x94\x21\xd4\xd1\x07\xdb\xc2\xc6\xd0\xc5\x80\x53\x61\x24\x94\xe1\xbf\xca\xf4\x8d\x7d\xc2\xab\xc8\xd3\x4b\x34\xd6\xe5\x8f\xc5\xbe\xa5\xf3\x1a\xb2\x45\x37\x6f\x99\x20\x15\x7f\x37\x3a\x11\x96\xdc\x3b\x22\xcd\x0a\xff\xab\x05\x93\x06\x8b\xfb\x87\xf1\xe4\xb6\x90\x11\x37\xf9\xc2\x89\xf6\xe4\x5f\xac\x18\xb8\xca\x5a\x9b\x5e\xdc\x27\x43\xd9\x11\xbd\x3b\x1b\x43\xc5\xbf\xe4\x5c\x5a\x0e\x73\x2d\xcc\x0a\xa7\x47\x47\x2f\x79\xf5\x97\x4d\x5b\x81\x23\x4b\xa3\xa7\xed\xa2\x6f\xea\x79\x54\x5a\x42\x2a\x97\x4e\x33\x17\x64\x74\xbc\x6d\x86\xe1\xe3\x87\xbd\x20\x46\x1c\x2a\x62\x4b\x5f\x97\x22\xe3\x10\xc0\x96\x0c\xe8\x2b\xb2\x16\xce\x73\x25\xbe\x51\x08\xee\x04\xb5\x88\x2e\x5f\x08\x3a\xe1\x44\xcb\x9f\x32\x8f\xcd\x52\xd5\x4b\xb4\x62\xcb\xbb\xc3\x10\xf1\xfa\xe3\x1c\x96\xd5\x51\x1f\x0a\x4e\x91\x91\xf3\x17\x59\x7a\x72\x78\x55\xaa\xb4\x1f\x8e\x15\x6d\x2b\x1c\xaf\x85\xce\x03\xf2\x07\xa3\xfa\xa9\xe0\x00\x56\xb4\xc5\x6e\xd7\x87\xc3\xb8\x3e\xd0\x83\x19\x19\x1a\x4d\x2c\xb8\x8a\x5d\xfa\x18\xa6\x2b\x81\xd0\xbc\x26\xac\x6b\x73\x38\x39\x86\x61\x01\xce\x73\x63\xcc\xcb\x2a\xfd\x12\xc9\xf1\x0a\xb4\xa6\xbf\x25\x35\x31\xad\x96\x14\x42\x22\xff\x6e\x08\x0f\x8c\xfa\x1f\x21\xb4\x42\x99\x20\x94\x21\xc7\x9f\xe5\xda\xb6\x9d\xa6\x40\x7c\x51\x18\xae\x1e\xf3\x18\x60\x6c\x20\x04\x12\x2d\xac\x5b\x08\x53\xee\x6f\x2f\xb2\xd2\x67\x1c\xb9\x0d\x3e\x0c\x86\xfd\x9d\x8e\x15\x1c\x5c\xd5\x9e\x81\x6e\xfb\x1b\x1b\x83\xa8\x15\x4a\x7f\x0b\x75\x93\x0e\xfa\xc6\xda\xfd\x13\x00\x00\xff\xff\x8c\x75\x13\x8d\x9e\x0a\x00\x00")

// FileIdentityLaforgeTmpl is "identity.laforge.tmpl"
var FileIdentityLaforgeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x6c\x92\xcf\x6e\xdb\x3c\x10\xc4\xef\x7a\x8a\xf9\x0c\x1f\xf2\x19\xa9\x75\x2f\xe0\x83\x11\xdb\x81\x80\x34\x0e\x82\xb4\x68\x4f\xc1\x86\x5c\x59\x9b\xd0\xa4\x4a\x52\x36\x54\x41\xef\x5e\xe8\x8f\x95\x36\xe8\x51\x3b\x3f\xed\xce\x0c\x98\x2e\x92\x04\xd8\x7e\x5f\x7f\x79\xb8\xdb\xe2\x6e\xbd\xdb\x3f\xde\x6e\x91\x6d\xb6\xf7\x4f\xd9\xd3\x0f\xdc\xec\xef\x77\xd9\xed\xd7\xc7\xf5\x53\xb6\xbf\x4f\x92\x45\x9a\x24\xe9\x02\x9a\x95\x21\xcf\x20\x54\x56\x7e\x56\x8c\x6c\x03\xb1\x88\x05\xc3\xbd\xbc\xb2\x8a\x10\x2b\x51\xc8\xc8\x2f\x8a\xe2\x2c\x16\x69\x22\x9a\x6d\x94\x58\x63\xd6\x34\x98\x2f\xb3\x0d\xda\x76\x86\xa6\xbb\x9e\x2e\x10\xeb\x52\x14\x19\x94\xec\x83\xb3\x04\xcd\x91\xc4\x84\xee\x47\x20\x17\x1f\xa2\xa5\x23\x63\x35\xfe\xbd\x9b\x26\x6d\x3b\x4b\x00\x43\x1f\x80\x3b\xfa\x4b\xe7\x23\x89\x99\xc4\x6d\xff\x35\x28\x25\x85\x70\x76\x5e\x4f\xe2\xc3\x65\xd0\xe9\x09\xa0\x39\x28\x2f\x65\x1f\xe3\xc2\x6c\xfe\x98\x8d\x58\xba\x40\xed\x2a\x28\xb2\x08\x25\x2b\xc9\x6b\x10\x54\x15\xa2\x3b\xa2\x0a\xec\x41\x27\x8a\xe4\x21\x79\xcf\x9d\x25\x14\xb8\xd2\x9c\x53\x65\x62\xb7\x77\xf6\xff\x90\x75\xc0\x9e\x73\x31\xef\x61\xd6\xfd\x6c\xd7\x8d\xde\xaf\x9d\xc8\x07\x18\x8e\xfd\x3a\xcd\xb9\x58\xbe\xdc\x53\xce\xe6\x72\xa8\xfc\xd0\x7d\x49\x9e\x8e\x1c\xd9\x07\x9c\x0b\x51\x45\xef\xf1\x85\xe1\x39\x7a\x39\xb1\xc6\x4b\x8d\x21\x4e\x00\x47\xb5\x1c\x7c\xf4\xeb\x57\x68\x12\x00\x68\x1a\x78\xb2\x07\xc6\xfc\x8d\xeb\x6b\xcc\x4f\x64\xf0\x79\x85\xf9\xf2\x5b\x47\x7d\x6a\xdb\x0b\xd5\xe9\x68\xdb\x8b\xf3\x8e\x1b\x6a\xee\x55\xb6\x7a\x84\xdb\x31\xc4\x6b\x15\x22\x8c\xbc\x31\xce\x12\x0b\xb8\x58\xb0\xef\xde\x02\x87\xeb\x7f\xd5\xe9\x6c\x6e\x44\x45\x84\xe8\x29\xf2\xa1\x46\xc1\x9e\x71\xe5\x99\x74\xff\xf8\xb4\x53\xe1\xbf\xb1\x48\x67\x9f\x27\x7e\x48\xa1\xdd\xd4\xe8\xde\xde\x8c\xda\x72\xe3\x26\x8b\x54\x96\x9d\xc5\x15\x3e\x32\xeb\x41\x18\xad\xb7\xbf\x03\x00\x00\xff\xff\x84\xf6\xe0\xe5\x34\x03\x00\x00")

// FileNetworkLaforgeTmpl is "network.laforge.tmpl"
var FileNetworkLaforgeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xa4\x92\xcf\x4f\xdb\x4e\x10\xc5\xef\xfe\x2b\xde\x17\x21\xc1\x37\xa2\xe4\x5e\x89\x43\x44\x02\x8a\x4a\x93\x2a\xa2\xb4\x37\x34\xb6\xc7\xf1\xc0\x7a\xd7\xec\x8f\x04\x37\xf2\xff\x5e\xad\xbd\x89\xe8\xa9\x87\x9e\xa2\x78\x9e\x3e\x33\xef\xed\x9b\x4e\xb2\x0c\x58\xfc\x9c\x7d\xfd\xf6\xb0\xc0\xc3\xec\x6e\xbd\xb9\x5f\x60\xb5\x78\xfc\xb1\xde\x7c\xc1\xed\x7a\x75\xb7\xbc\xff\xbe\x99\x3d\x2e\xd7\xab\x2c\x9b\x4c\xb3\x6c\x3a\x41\xc9\x85\x22\xcb\x20\x04\x2d\x6f\x81\xb1\x9c\x43\x34\x7c\xcd\x30\xf9\x0b\x17\x1e\xa2\xc5\x0b\x29\xf9\x45\x5e\x8c\xc6\x64\x9a\x69\xf6\x7b\x63\x5f\x71\x76\x38\xe0\xfc\x7a\x39\x47\xdf\x9f\xe1\x10\x77\x4f\x27\xd0\xd4\x30\xf6\xa2\x14\x72\x1e\x30\x49\x7d\xe1\xe0\x42\x5e\x9a\x86\x64\x80\x60\x54\xde\x24\xca\x2a\xfe\xe9\xfb\xb3\x44\x29\xa4\xb4\x10\xf7\x11\x00\x4b\x7a\xcb\xa8\x8c\x85\xaf\xc5\x9d\x3e\x0f\xac\x41\x7f\x64\xdd\x2e\xe7\x9b\x0f\x2c\x52\xca\xec\x5d\xfc\x81\xb7\x54\x55\x52\xa0\xb2\xa6\x19\xd8\x9e\xa9\xb9\x70\x78\x09\x4d\x8b\xdc\xbc\xb3\x83\x37\x7f\xe2\x2f\x9d\x34\x41\x91\x67\x07\x65\x0a\x52\xe2\xbb\xff\xc7\x9d\xbb\x52\x9e\x77\xe2\x24\x57\xd1\xc6\xb0\xf9\x69\xbe\x7c\x4a\x5f\xfa\x3e\xad\xdf\x91\x75\x50\xec\xd1\x99\x80\x92\x2b\xd1\x8c\x22\x38\x6f\x1a\x14\x46\x57\xb2\x0d\x76\x4c\xb6\x25\x4b\x0d\x7b\xb6\x0e\xfb\x5a\x8a\x1a\x0d\x75\x31\x44\xcd\x5c\x72\x39\x18\x77\x2d\x17\x12\x0d\xb4\xd6\xc4\xd5\x46\x8b\xde\x22\x0f\xa2\x4a\xb6\xee\x3a\xdd\x15\x17\xde\xe0\x90\x01\x88\x67\x8d\xb9\x9d\xbf\x72\x77\x85\xf3\x1d\x29\x7c\xbe\x89\xa7\x46\xd5\xa7\xbe\x3f\xaa\xe2\x1c\x7d\x7f\x0c\x31\xea\x62\x84\x69\xca\xba\x4c\xe2\xa3\x2d\x4f\x5b\x07\x7e\x6f\x2d\x3b\x87\x2d\x6b\xb6\xa4\x20\xba\x32\xb6\x19\xed\x8c\x1e\x8e\x4d\xa0\x98\x89\x37\xc8\x79\x40\xe2\x2d\xb0\x15\x2e\x61\xf4\xb1\x6e\x55\xf0\xc1\xf2\x68\x61\x80\xff\xd5\xc2\x63\x54\xfd\x83\x85\x97\xe0\x3c\x94\xbc\xc6\xbe\xfa\x1a\xc6\xd7\x6c\xe1\xbb\x96\xdd\xd5\xf0\x58\x05\xe9\x94\x78\x07\x1a\x1e\x4b\x49\xe1\xe1\xbc\x25\xcf\xdb\x0e\x35\x5b\xc6\xa5\x65\x2a\x07\x03\xa5\x29\xdc\x7f\xa9\x1b\x46\x3f\x9f\xf4\xa3\x8b\xd2\x9c\xfa\xb9\xd6\xb7\x69\x76\x3d\x37\xa7\x13\xa9\x6d\xe3\x89\xa9\x49\x1f\x34\xb3\x71\x90\x4e\xef\x7f\x07\x00\x00\xff\xff\x76\xcd\x4f\x47\xe6\x03\x00\x00")

// FileRemoteFileLaforgeTmpl is "remote_file.laforge.tmpl"
var FileRemoteFileLaforgeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x5c\x52\xcb\x6e\xdb\x30\x10\xbc\xeb\x2b\xa6\x46\x0f\x89\xe1\xc6\x87\x9e\x73\x30\x62\x39\x30\xe0\x58\x86\xeb\x02\xbd\x19\x34\xb9\x8a\xd8\x50\xa4\xca\xa5\x92\x0a\x86\xfe\xbd\x20\xe5\x57\x73\xe4\xce\xec\xcc\x70\xb0\xd3\x71\x96\x01\xf9\xaf\xd9\xcb\x66\x95\x63\x35\x5b\x14\xdb\xe7\x1c\xdb\xfc\xa5\xd8\xe5\x58\x2c\x57\x39\x9e\x8a\xf5\x62\xf9\xfc\x73\x3b\xdb\x2d\x8b\x75\x96\x8d\xa7\x59\xe6\xa9\x76\x81\xf6\xa5\x36\x84\xd1\xf1\x88\xaf\x0f\xcb\x39\xfa\x7e\x84\x63\x06\x4c\xc7\x50\xc4\xd2\xeb\x03\x31\x42\x45\x90\xad\xf7\x64\x03\x8c\x93\xc2\xe8\xd0\xc1\x95\x08\x95\x66\xa4\xfd\xbb\x34\x9e\xa0\xf5\x66\x02\xfe\x8e\x6f\x50\x54\x8a\xd6\x04\x3c\x62\x34\x3a\xc1\xf7\xf7\x18\x4f\x33\x80\x5d\xeb\x25\xed\x43\xd7\x50\x84\x93\xf5\x8f\x34\xdb\xc5\x51\xdf\x8f\xb2\x21\x41\xf4\x6d\xad\xfe\xd3\x12\x84\x52\x9e\x98\x11\x5c\x4a\x10\x28\x81\xc9\xfb\x20\x98\x14\x9c\xfd\x4f\xf7\xc6\xe8\x93\xc7\x8d\x7e\x52\xd2\xce\xc6\xe5\x28\x37\x34\x02\xee\x38\x50\x7d\x35\xe0\xca\xb5\x46\xe1\x40\x68\x8c\x90\xa4\x06\x71\x45\x1c\xb4\x1d\xf6\xcf\x0e\xf3\x9b\xd9\xd5\xa6\xb5\xfa\x2f\x1a\xf2\x35\x43\x5b\x38\x19\x84\x81\x75\x61\xa0\x25\xad\x01\x3c\xab\x6c\xd2\xeb\xba\xaf\x34\x8b\x83\x21\x05\xd6\x75\x63\x3a\x08\x63\xdc\x07\x0f\xed\x4b\x57\xd7\xc2\xaa\xd8\x4b\xcc\x27\x38\x75\xf1\x4e\x1e\x24\x58\x9b\x2e\x3a\x8a\x0b\x4b\x56\x42\xdb\x09\xa8\x2c\x49\x06\xfd\x4e\xa6\x43\x2d\xde\xb4\x7d\x85\x0e\x10\x58\x17\xc5\xe6\xf4\xb9\xb3\xe5\x23\x86\x8f\x9d\xdf\x7d\x7f\x0a\xf5\xbb\xe5\x00\xa3\xdf\x08\x1f\x3a\x54\x70\xa1\x22\x8f\x58\x3d\x4f\xd0\xb9\x16\x52\x58\x70\x43\x52\x97\x5d\x0a\x60\x4b\xa3\x65\x00\x07\x2f\x02\xbd\x76\xa8\xc8\x13\xee\x3c\x09\x95\x8a\x56\x4e\xf2\x97\xd3\x7d\x38\xbb\xbf\xf0\xe3\x2d\x02\xca\x5d\xca\x29\xec\xd3\x09\x7b\x98\xbb\x54\x52\x24\x88\xa6\x21\x7b\x49\x7b\xc3\x99\x0d\x40\xdf\x67\x40\x9f\xf5\xff\x02\x00\x00\xff\xff\x78\x77\x17\x0a\x29\x03\x00\x00")

// FileScriptLaforgeTmpl is "script.laforge.tmpl"
var FileScriptLaforgeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xa4\x56\x4d\x6f\xdb\x46\x10\xbd\xeb\x57\xbc\x08\x06\x22\x1b\x8c\x5d\xa0\xb7\x02\x3e\x18\x96\x13\x08\x48\xac\x20\x76\x8b\x02\x45\x61\x8c\xc8\x21\xb9\xd1\x72\x97\xd9\x0f\x2b\xac\xc0\xff\x5e\xec\x07\xf5\xd1\x06\xe8\xa1\x27\x92\xb3\xb3\xf3\xde\xcc\xbc\x9d\xe5\xcd\xd5\x6c\x06\x3c\xfc\x7e\xf7\xe9\xf3\xc7\x07\x7c\xbc\x7b\xbf\xfe\xf2\xe1\x01\x4f\xf7\x5f\x56\x9f\x9f\x71\xbf\x7e\x7c\xbf\xfa\xf0\xeb\x97\xbb\xe7\xd5\xfa\x71\x36\xbb\xba\x99\xcd\x6e\xae\x50\x71\x29\xc9\x30\x08\x5e\x89\x6f\x9e\xb1\x5a\x42\x28\xb8\x96\xa1\x37\x5f\xb9\x74\x10\x4a\x38\x41\x52\xfc\x45\x4e\x68\x85\xab\x9b\x99\x2d\x8d\xe8\x1d\xe6\xfb\x3d\x2e\xae\x57\x4b\x8c\xe3\x1c\xfb\x80\x7c\x73\x05\x42\xa7\x0d\xa3\xf5\x1d\x29\x18\xa6\x8a\x36\x92\xa1\xa8\x63\x88\x1a\x83\xf6\xe8\x0d\xd7\x6c\x0a\x90\xaa\x40\xa8\x38\x45\x0b\xa1\x17\x3a\x3e\x49\x5e\x06\x14\xa4\x5d\xb7\x19\xe7\x31\x7c\x8c\xe3\x7c\x86\xb3\x3d\xd3\xf2\xf2\xc4\x16\xbc\x12\x9b\x90\x86\x24\xd5\x78\x6a\x38\x7e\x64\xea\xc2\x62\x67\x84\x73\xac\x42\xb2\x0b\xdb\xb2\x94\x05\x7a\xbd\x63\x13\xdf\x33\x81\xc3\xd6\x09\xe5\xe3\x64\x38\x42\x24\x32\x1b\xb6\x31\x7e\xe9\x8d\x61\xe5\x20\x75\x49\x52\xb8\x01\xba\x86\x6b\x85\x9d\x80\x17\x71\xa1\x80\x37\xb2\x80\xfd\x19\xef\x50\x71\x4d\x5e\xba\x00\x31\xcf\xcb\x97\x19\xde\x6a\x6f\x4a\x7e\x71\x43\x7f\x64\xf0\x14\x6d\xcf\xc1\x74\x9e\x66\x6e\x1f\x55\x95\x61\x6b\xe1\x74\xe4\xe0\xce\xd2\xde\x90\xe5\x0a\x5a\x9d\x45\x3e\x81\xfa\x07\xca\x09\x82\xf5\x7d\x2f\x07\x90\x69\x7c\xc7\xca\x59\xb4\x6c\xf2\x56\x32\x8d\xc5\x2d\xfe\x98\x01\xc0\x7e\x0f\x43\xaa\x61\x5c\xbc\x14\xb8\x20\xd3\xe0\x97\x5b\x5c\x5c\xdf\x05\x9f\x77\xe3\x18\x7d\x22\x44\x58\x1a\xc7\x79\x31\xed\x62\x55\x65\x87\x3f\x33\xa4\xa8\x21\x1a\xa5\x0d\xbf\xb0\x31\xda\xd8\xd0\x34\xa5\x1d\x9c\xf1\x5c\x40\x52\xad\x4d\xc3\xd8\x09\x29\xd1\x92\x74\xe8\x8d\x7e\x15\x56\x68\x25\x54\x03\xdf\x6b\x05\x52\x88\x5b\x61\xd8\xf6\x5a\xd9\x4c\xf8\x3c\xea\x2d\x92\x8c\xa3\xf1\x21\xd9\xc6\x31\x73\x28\xb5\x96\x95\xde\x29\xd8\x9e\x4b\x51\x0b\xb6\x20\x94\xde\x3a\xdd\x61\x47\xc2\xc1\x89\x8e\x61\x5b\xed\x65\x15\xc5\xbd\x13\xb6\x0d\xb5\x8f\x8b\x54\x3b\x36\x49\x16\xba\xeb\x82\xde\xf9\xbb\x70\x16\x8b\x63\xcf\x7f\xca\xad\x3e\x00\x65\x3a\xf7\xd3\xf7\x81\x4a\x25\x6c\x38\x4a\x15\xac\xe8\x62\x2b\xa4\xd4\x3b\x9b\xc4\x35\x85\x77\x1a\x1b\x46\x4f\x36\xb6\xf9\x95\x0d\x98\xac\x90\x43\x10\x39\x1d\xbc\xca\x96\x84\x2a\xc0\x75\xcd\xa5\x13\xaf\x2c\x07\x74\xb4\x0d\x55\x0b\x9c\xf1\xb8\x5e\x7f\x4e\xa4\x0e\x90\x99\xd4\x72\xfa\x3e\x90\xfa\xea\xad\x83\x14\xdb\xd0\x07\xd7\x42\xbb\x36\x24\x3c\xf4\x6c\x8b\x58\x8e\x92\xa6\xd2\x0d\x91\x80\xaa\xa5\x28\x1d\xac\x33\xe4\xb8\x19\x92\x8a\x16\x61\x4c\xc4\x32\x55\xba\xb4\x6f\x72\x45\xb4\x7a\x39\xf8\xef\xa3\x4a\x2a\x7d\xd0\xe7\x5a\xdd\xe7\xb5\xeb\xa5\xce\x53\x01\xa0\xbe\x0f\x32\xca\x6c\x4f\x7c\xee\xd2\x42\x94\xd7\x78\x72\x66\x84\xc6\x46\xea\x72\x0b\xc9\xce\x46\xc2\x15\xd7\x42\xf1\xd4\x62\xed\x5d\xef\x1d\x16\xa1\x6a\x42\x85\x57\xa1\xa6\x32\xbe\xb5\x28\xc9\xf2\x25\x6a\x6d\x52\x17\xf8\x3b\x97\x7e\x9a\x91\x08\xc1\x13\xef\x70\x7c\x5c\x25\x42\x6f\x7b\x72\x51\x1e\x84\x5a\x48\x0e\x6f\x25\x85\xa0\x4e\xe3\xe9\x79\xb9\x7a\xcc\xc1\x38\x88\xb9\x31\xd4\xbd\xfd\x57\x54\x1c\x62\xe5\xf1\xbb\xbe\x7e\x8a\x86\x7c\x56\x27\x38\xed\x5d\x11\x9e\x6c\x0c\xc2\x7c\xdf\x48\x52\x5b\x2c\xe6\xf3\x4b\x6c\x86\x69\xe8\xa4\x21\x1c\x8f\x90\xd4\x4d\xa0\x13\xb0\x37\x5e\xc8\x0a\x95\x30\x71\x35\xc5\x42\xe5\x4d\x50\xc8\xd9\x21\x3b\x12\x0a\x1e\xe7\x8c\x82\x65\xea\x4b\xa6\x71\xee\x10\x2c\xc9\x61\xea\xc8\x2b\x19\x1b\x3a\xf1\x83\x46\x04\x25\x88\xc6\x9b\x74\x05\xf5\x64\xa8\x63\xc7\xc6\x62\xd7\x8a\xb2\x45\x47\x43\x10\xbe\x62\xae\xb8\x8a\x35\xcc\xe7\xb5\x3c\x27\x1c\x33\x63\x63\xaf\x13\xf5\x08\x78\x9b\xbb\x74\x9c\x5c\x5b\x1e\x0a\x5c\xbc\x92\x4c\xb3\xeb\xb7\xe0\x35\xcd\xae\x90\xc0\x96\x07\x8c\xe3\x94\x4e\xf0\x9b\x12\x3d\x9b\x63\x07\xa1\x51\x13\xba\xd8\xc7\xb1\xdc\xb0\x62\x43\x12\x42\xd5\xda\x74\x29\x9d\x94\x43\xec\xc2\x86\x11\xef\xcb\x78\x90\x63\x48\x7c\xf3\x6c\x44\x9a\xda\xf9\x5e\xae\xbd\xf3\xd3\xf0\x8d\xc1\xff\x33\x85\x67\x6a\xfe\x57\x0a\x1d\x09\xe5\x48\x28\x36\x48\xd3\xa6\x97\xec\xc2\xdc\x98\x2e\xec\x02\x1b\xef\xc2\x74\x66\x38\xa6\x0e\xda\x34\xa4\xf2\x1f\xc3\x9b\xc4\xf4\x24\x46\x92\xc1\xa7\x83\xe1\xf8\x0f\x11\x18\x9c\xdd\xfa\x27\x4e\x27\x3f\x00\x00\x77\x24\xe4\x8f\xbc\x1e\xe2\xc2\x24\xac\xf1\xef\x00\x00\x00\xff\xff\xa7\xcc\x0c\x0c\x12\x09\x00\x00")

// FileUserLaforgeTmpl is "user.laforge.tmpl"
var FileUserLaforgeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x3c\x8b\x31\x0a\xc2\x30\x18\x46\xf7\xef\x14\x1f\xc1\xa9\x83\x3d\x81\x43\xb4\x49\x08\x84\x04\x4a\x73\x80\x82\x19\x0a\xd6\x41\xc9\x14\xfe\xbb\x4b\x55\x32\xbe\xf7\x78\xe3\x00\x90\x93\xb1\x3a\x87\x85\x41\xdb\x34\x3b\x43\x17\xd2\x55\x07\xde\x52\xb4\xde\xe5\x59\x2f\x3e\x45\x60\x18\x81\xfa\x2e\x2f\xaa\xd6\x78\x3a\xfb\x89\x22\x8a\x0d\xe4\x73\xdd\x0b\x2f\x7f\x1f\x0f\x10\x51\x20\x6b\xdd\xee\xdd\xe7\xfc\x3b\x40\x96\x7d\xdd\x1e\x3d\x98\x2f\x1d\x45\x3e\x01\x00\x00\xff\xff\x26\xf3\x6c\x27\x8f\x00\x00\x00")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileAmiLaforgeTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "ami.laforge.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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

	rb = bytes.NewReader(FileCommandLaforgeTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "command.laforge.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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

	rb = bytes.NewReader(FileCompetitionLaforgeTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "competition.laforge.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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

	rb = bytes.NewReader(FileDNSRecordLaforgeTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "dns_record.laforge.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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

	rb = bytes.NewReader(FileEnvironmentLaforgeTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "environment.laforge.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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

	rb = bytes.NewReader(FileHostLaforgeTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "host.laforge.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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

	rb = bytes.NewReader(FileIdentityLaforgeTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "identity.laforge.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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

	rb = bytes.NewReader(FileNetworkLaforgeTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "network.laforge.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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

	rb = bytes.NewReader(FileRemoteFileLaforgeTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "remote_file.laforge.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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

	rb = bytes.NewReader(FileScriptLaforgeTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "script.laforge.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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

	rb = bytes.NewReader(FileUserLaforgeTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "user.laforge.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
