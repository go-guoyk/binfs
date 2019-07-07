// generated by binfs

package binfsecho_test

import (
	"go.guoyk.net/binfs"
	"time"
)

var (
	binfs5e25b9a86aa1de2a6e9b6a5b49263d76c735b704 = binfs.Chunk{
		Path: []string{"testdata", "dir1", "file2.txt"},
		Date: time.Unix(1562488467, 0),
		Data: []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x32, 0xa},
	}

	binfs95e9c658d800db135ecfd6335958a4bcfdb95ec4 = binfs.Chunk{
		Path: []string{"testdata", "dir3", "index.txt"},
		Date: time.Unix(1562488477, 0),
		Data: []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x33, 0xa},
	}

	binfsce1be0ff4065a6e9415095c95f25f47a633cef2b = binfs.Chunk{
		Path: []string{"testdata", "file1.txt"},
		Date: time.Unix(1562488456, 0),
		Data: []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x31, 0xa},
	}
)

func init() {

	binfs.Load(&binfs5e25b9a86aa1de2a6e9b6a5b49263d76c735b704)

	binfs.Load(&binfs95e9c658d800db135ecfd6335958a4bcfdb95ec4)

	binfs.Load(&binfsce1be0ff4065a6e9415095c95f25f47a633cef2b)

}
