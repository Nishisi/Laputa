package laputa

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

const LAPUTA = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x55\x4b\x6e\xe3\x30\x0c\xdd\xe7\x14\xdc\xa9\x03\x90\xea\x7e\x66\x78\x15\x79\x2d\x40\xab\x59\x0b\x84\xd3\x6c\x3a\x40\xd0\x02\xdd\xcf\x01\x0a\xf4\x08\x09\x7a\x97\xf8\x00\xb9\xc2\x80\x94\xfc\x89\xad\x20\x44\x22\x4b\xa2\xc5\xc7\x47\x52\x34\x34\xe5\x7a\x7a\xd5\xc7\x70\x38\xb7\xf5\x2b\x41\x80\xcb\xf1\xc3\xde\xc6\x3a\xdb\xb5\x94\x00\x4f\xf3\xa1\xd1\x78\x87\xec\x4c\x82\x8c\x12\x74\xc9\x88\x88\xd5\x91\x85\x4f\x2a\x6e\xd7\xd8\xbb\x59\x19\x30\xa9\x61\x01\x20\x56\xd3\x81\x6c\x64\x26\x00\x33\x3e\x39\xf2\xf2\x6f\xc2\x38\xf7\x4b\xdf\x37\x9a\x71\xf5\xdb\xac\x16\xb3\xa0\x8f\xc0\x06\x12\x1c\x93\x38\xc6\x15\xd5\xd1\xd9\xbf\xd5\xf2\x46\xb5\x88\x89\x0b\x0e\xba\x0e\x4d\xa2\x88\x24\x66\xce\x22\xd1\xc6\x7c\x3d\x7d\x9b\xc6\xfb\xe1\xf0\xb6\x61\xbc\x30\xb3\xbb\xc1\x79\xba\x1c\x3f\x1a\x2f\x45\x81\x0c\x00\x6a\xb5\x1b\x0e\x5f\x43\xbf\x4f\x44\x89\x68\x38\x7c\xf9\x5f\x44\x14\x01\x22\x0d\xfd\xfe\x2e\xd0\x48\xea\xae\xba\xc8\x9f\x92\xdb\x04\xc9\x96\x19\x7e\x66\xb8\xbc\x7e\xea\x3f\x23\x31\x76\x1d\x12\x76\x19\x20\x4b\x0c\x71\xce\x4c\x9b\x55\x03\xec\x7a\xee\x7f\xd8\x24\x85\x90\x35\x3e\x44\xcc\x3e\x8a\xa4\xa9\xa6\x44\x99\x42\x04\x48\xe0\x73\xa9\xb0\x24\x0a\xf6\x08\x0e\x00\x17\x75\xb0\x20\xfe\x1e\x45\xa2\x83\xa4\x34\x2e\xaf\x9f\x09\xc0\x5f\xcf\x47\x30\x66\x65\x28\x7c\x73\xd7\x65\xf0\x19\xbc\xd2\x05\x9d\x41\x22\x5c\x53\x71\x15\x4b\x85\x9d\x7a\xab\x3f\x07\x19\x91\xa3\xe5\x7f\x14\xad\x83\x10\xb2\x80\xb8\x48\x9a\x25\x66\x8d\x59\x06\x17\x6f\x52\xb1\x30\xae\x90\x80\xcc\x48\x7a\xcf\x42\xf5\x2c\x7b\x75\x73\x39\x58\x35\x14\x37\xcb\xa4\xce\xb3\x48\x88\xe2\xe2\xd0\xef\x77\xdb\xf0\x23\xc7\x20\xe2\xed\x45\x25\xa9\x7f\xc8\xe8\xbd\x47\xe5\x8e\xdb\xd1\x9b\x4e\x57\xbe\x2c\x6c\x95\xbb\x3c\x6d\x60\xae\x7e\x67\x89\xdb\x02\x4b\x62\x11\x2d\x6f\x58\xba\x89\x59\x33\x66\xb7\x02\x6a\xa6\xa7\x3c\x0e\x87\x37\x11\x59\xea\x05\x86\x97\x73\x60\xf2\x56\x00\x5a\x32\x79\xb7\x4a\xb5\xcf\xde\x27\x80\x68\x6d\x44\x84\xa8\x56\xca\xf5\x74\x9a\xc6\x34\xd5\xc2\x6a\x32\x4d\xb5\xb4\x39\xc8\x9a\x42\xd6\x4c\x68\x13\xb1\x1e\xf7\x6d\x2e\x11\xa2\xaf\x96\xbd\x0e\x5a\x2b\xa3\xc1\xfa\x2c\xd7\x9f\x64\x0d\x56\xaa\xf5\xa6\xd3\xaa\xc5\xe0\x1c\x0f\xfd\x5e\x01\x3a\x08\x8e\x20\x6e\x5a\xed\x42\x82\xab\xc2\x11\xc8\x05\xd9\x44\xfd\xa6\xad\x37\xbe\x0d\x12\x82\x73\xb6\x5f\x94\x71\xac\x3d\xbd\xa0\xc7\xba\xf1\x3c\x79\x3e\xe5\xaa\xd5\x3d\xe6\x66\xb9\x66\x05\x2b\x97\xec\xd1\xef\x4b\x45\x4f\x50\x49\x91\xfc\x14\xb6\xf7\xbb\x40\xba\x39\xb3\x9a\xd1\x1a\x98\xd7\xd3\xa9\x33\xac\xb8\x20\xf8\x5c\xcf\x21\xb3\xbb\x4f\x67\x1b\x3f\x68\x60\x2e\x37\x9d\xe3\xfa\x05\xfc\x2e\x3f\x34\x08\xb9\x39\xd8\x6a\x4c\x0d\xa8\x46\xf7\xdf\xca\x83\x1e\x3e\x5b\x9c\xef\xd8\x23\xd1\x33\xff\x03\x00\x00\xff\xff\xa3\x82\x16\x2f\x61\x08\x00\x00"

func (laputa *laputa) Art() (string, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(laputa.art))
	if err != nil {
		return "", fmt.Errorf("Faild to read laputa bytes: %s", err.Error())
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gzErr := gz.Close()

	if err != nil {
		return "", fmt.Errorf("Faild to read laputa bytes: %s", err.Error())
	}

	if gzErr != nil {
		return "", fmt.Errorf("Faild to read laputa bytes: %s", gzErr.Error())
	}

	return buf.String(), nil
}
