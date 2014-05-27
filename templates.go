package namedwebsockets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func templates_console_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xcc, 0x58,
		0x7d, 0x73, 0x1a, 0xbf, 0x11, 0xfe, 0xdb, 0xfe, 0x14, 0x1b, 0x4d, 0x26,
		0xbc, 0xd8, 0xdc, 0x39, 0x69, 0x32, 0xd3, 0x62, 0xa0, 0x63, 0x63, 0xa6,
		0x75, 0xc7, 0xc1, 0x9e, 0xda, 0x49, 0xa6, 0x4d, 0xd2, 0x44, 0xdc, 0xc9,
		0x70, 0xf1, 0x71, 0x62, 0x24, 0x61, 0x4c, 0x1c, 0xbe, 0x7b, 0x57, 0xd2,
		0xbd, 0xe8, 0x80, 0x23, 0x6f, 0xed, 0xcc, 0x8f, 0x3f, 0x38, 0x38, 0xed,
		0x6a, 0x9f, 0x7d, 0xf4, 0xec, 0x4a, 0x77, 0x9d, 0x27, 0x67, 0x97, 0xfd,
		0x9b, 0x7f, 0x5d, 0x0d, 0x60, 0xa2, 0xa6, 0x31, 0x5c, 0xbd, 0x39, 0xbd,
		0x38, 0xef, 0x03, 0x69, 0xf9, 0xfe, 0xbb, 0x3f, 0xf5, 0x7d, 0xff, 0xec,
		0xe6, 0x0c, 0xfe, 0x7e, 0xf3, 0xfa, 0x02, 0x5e, 0x7a, 0x47, 0xcf, 0x7d,
		0x7f, 0x30, 0x24, 0xbd, 0xfd, 0x8e, 0xb6, 0xd4, 0x17, 0x46, 0xc3, 0xde,
		0xfe, 0x5e, 0x47, 0x45, 0x2a, 0x66, 0xbd, 0x21, 0x9d, 0xb2, 0xf0, 0x1d,
		0x1b, 0x5d, 0xf3, 0xe0, 0x8e, 0x29, 0xb8, 0x61, 0x52, 0x41, 0x9f, 0x27,
		0x92, 0xc7, 0xac, 0xe3, 0x5b, 0x93, 0x7d, 0x34, 0x96, 0x81, 0x88, 0x66,
		0x0a, 0xdd, 0xf6, 0xfc, 0x66, 0xb3, 0x89, 0x17, 0x68, 0xc2, 0xa9, 0xe0,
		0x34, 0x0c, 0xa8, 0x54, 0x85, 0xfb, 0x01, 0x5c, 0xf0, 0x80, 0xc6, 0xc5,
		0x0d, 0x39, 0x89, 0xa6, 0x10, 0x47, 0x23, 0x41, 0xc5, 0xd2, 0x7a, 0xb5,
		0x7e, 0xf3, 0x63, 0x66, 0xb1, 0x53, 0x9d, 0x5c, 0x9d, 0xc3, 0x1b, 0x49,
		0xc7, 0xac, 0xbd, 0x3e, 0xb5, 0x63, 0xa4, 0x3f, 0xbe, 0x5f, 0x80, 0x05,
		0x9a, 0x84, 0x10, 0xf0, 0x24, 0x61, 0x81, 0x82, 0x45, 0xa4, 0x26, 0xc0,
		0xd5, 0x84, 0x09, 0x98, 0x31, 0x26, 0x24, 0xcc, 0x65, 0x94, 0x8c, 0x01,
		0x6f, 0x80, 0x44, 0x62, 0x40, 0x32, 0x71, 0x1f, 0x05, 0x0c, 0x12, 0xfd,
		0x27, 0x4a, 0xcc, 0x40, 0x30, 0x17, 0x82, 0x25, 0x0a, 0x12, 0xa6, 0x16,
		0x5c, 0xdc, 0x15, 0x51, 0xee, 0xa9, 0x80, 0x85, 0x84, 0x2e, 0x8e, 0x2c,
		0xb6, 0x90, 0x53, 0x27, 0xd3, 0xe5, 0xb5, 0x9d, 0x4f, 0x93, 0x4e, 0x1a,
		0xc7, 0x0e, 0x4a, 0x2e, 0xda, 0x9b, 0x98, 0xfb, 0xbf, 0x82, 0x92, 0x5b,
		0x94, 0xb1, 0x5e, 0x07, 0x08, 0x99, 0x1e, 0xa8, 0x80, 0x58, 0x5e, 0xaa,
		0x9d, 0xf0, 0x3c, 0xcf, 0xc3, 0x49, 0x13, 0x0c, 0xcc, 0xcc, 0xec, 0x82,
		0xa9, 0xb9, 0x48, 0x58, 0x08, 0x9f, 0x17, 0xf2, 0x33, 0xf0, 0xd1, 0x17,
		0x0d, 0xf3, 0xcb, 0x1c, 0xd9, 0x8d, 0xa3, 0x3b, 0x06, 0x14, 0x12, 0x2e,
		0xa6, 0x18, 0xff, 0x1f, 0xf4, 0x9e, 0x5e, 0x1b, 0xe9, 0x40, 0xa1, 0x09,
		0x6b, 0xee, 0x65, 0xd3, 0x37, 0x9b, 0x3e, 0x7e, 0xd7, 0x6f, 0xe7, 0x49,
		0xa0, 0x22, 0x9e, 0xd4, 0xc7, 0x31, 0x1f, 0xd1, 0xb8, 0x01, 0x8f, 0xa8,
		0x3b, 0x94, 0x9b, 0x0f, 0xcd, 0x93, 0x78, 0x41, 0x97, 0xb2, 0x99, 0xaf,
		0x9a, 0xe2, 0xc0, 0xe7, 0x02, 0xf8, 0x22, 0xb1, 0x69, 0x4e, 0xb8, 0x54,
		0xad, 0x11, 0x95, 0x88, 0x67, 0x26, 0xf8, 0x83, 0x16, 0xda, 0x9e, 0xce,
		0x94, 0x25, 0xe1, 0x8c, 0x47, 0x89, 0x7a, 0x23, 0xe2, 0x53, 0x1c, 0xc5,
		0xb4, 0xc9, 0x42, 0xb6, 0x7d, 0x3f, 0x77, 0x6a, 0x3f, 0x3e, 0x3e, 0x5d,
		0xad, 0x7c, 0x72, 0x6c, 0x42, 0x65, 0x08, 0x20, 0x92, 0x6f, 0x69, 0x1c,
		0x85, 0x0e, 0x1b, 0x75, 0x59, 0xfc, 0xd6, 0xc8, 0xd0, 0x7a, 0xcf, 0x72,
		0x00, 0xfe, 0x7f, 0xde, 0x9f, 0xb4, 0xfe, 0x4d, 0x5b, 0x5f, 0x8f, 0x5a,
		0x7f, 0xf9, 0xe0, 0x7d, 0x6a, 0x7d, 0x7c, 0x7c, 0x7e, 0xf8, 0xe2, 0xd5,
		0xab, 0xd5, 0x53, 0xdf, 0x53, 0x58, 0x4c, 0x25, 0x4f, 0x4d, 0xe9, 0xde,
		0x6a, 0x3f, 0xc3, 0xb7, 0x56, 0x7b, 0x5d, 0xc8, 0x11, 0xb8, 0x5e, 0x87,
		0x20, 0xe7, 0x23, 0xcc, 0x4b, 0xf1, 0x80, 0xc7, 0xf2, 0x10, 0xd1, 0xe5,
		0xda, 0xca, 0xa0, 0x44, 0xb7, 0x50, 0x7f, 0xf2, 0x1d, 0xd4, 0x99, 0xed,
		0x9e, 0x9a, 0x08, 0xbe, 0x00, 0x72, 0x9e, 0xdc, 0x6b, 0x73, 0x48, 0xed,
		0x0d, 0x98, 0x36, 0x10, 0x2c, 0x61, 0xc7, 0xcb, 0x00, 0x46, 0xc4, 0x4e,
		0xbe, 0x5a, 0x39, 0x85, 0x68, 0xd6, 0x29, 0x3e, 0x80, 0xba, 0x83, 0x0f,
		0xfe, 0x0a, 0x64, 0x94, 0xfd, 0x21, 0x80, 0xd3, 0x1b, 0xea, 0x49, 0x03,
		0xed, 0x88, 0xbf, 0x16, 0xab, 0x9c, 0x67, 0xca, 0xd5, 0x71, 0x4e, 0xd6,
		0x96, 0x6e, 0xf3, 0x43, 0x84, 0xad, 0x2d, 0x97, 0x86, 0x5f, 0xe6, 0x7d,
		0x17, 0xd7, 0x4a, 0xcc, 0xd9, 0x06, 0x92, 0xb5, 0x16, 0xf7, 0xff, 0x47,
		0x71, 0x4b, 0x63, 0x59, 0x86, 0x81, 0x55, 0x31, 0x78, 0x98, 0x71, 0x24,
		0xdc, 0x56, 0x4b, 0x0e, 0x41, 0x9a, 0x61, 0xa3, 0x07, 0x3b, 0xe2, 0x6d,
		0xf2, 0x96, 0x61, 0xa9, 0x34, 0xc0, 0x94, 0xd2, 0x2a, 0xf4, 0xa6, 0x3c,
		0x9c, 0xc7, 0x0c, 0xbe, 0x7d, 0x83, 0xc7, 0x55, 0xc3, 0x63, 0x18, 0x53,
		0x28, 0xdd, 0x3f, 0x36, 0x9d, 0x1c, 0x6d, 0xbb, 0xe1, 0xcb, 0x64, 0xad,
		0x85, 0xde, 0x60, 0xf2, 0x7b, 0x61, 0xcb, 0x0e, 0x45, 0xc8, 0x55, 0xa3,
		0xae, 0x26, 0x91, 0x11, 0x4d, 0xc7, 0xcf, 0xf6, 0xaa, 0xd2, 0xb6, 0x95,
		0xaf, 0x51, 0xcc, 0xc7, 0xf5, 0xa9, 0x1c, 0xa7, 0x48, 0x42, 0x1e, 0xcc,
		0xa7, 0xd8, 0xcd, 0xbd, 0x31, 0x53, 0x83, 0x98, 0xe9, 0x9f, 0xa7, 0xcb,
		0xf3, 0xb0, 0x5e, 0x43, 0xb3, 0x5a, 0xc3, 0xa3, 0xb3, 0x19, 0x2a, 0xbc,
		0x3f, 0x89, 0xe2, 0xb0, 0x9e, 0x9b, 0x06, 0x82, 0x51, 0xc5, 0x6e, 0xd8,
		0x83, 0x1a, 0xf2, 0x90, 0xd5, 0xf5, 0x5a, 0x9e, 0xe1, 0x8d, 0xba, 0x56,
		0x75, 0x0d, 0x9b, 0x6c, 0x0d, 0xaf, 0x18, 0x41, 0xff, 0xfb, 0x90, 0xd4,
		0x1a, 0x66, 0xe1, 0x56, 0x2e, 0x04, 0xa9, 0xa8, 0x9a, 0x4b, 0x07, 0x45,
		0x86, 0xe9, 0x78, 0x27, 0x24, 0xeb, 0x86, 0xa8, 0x14, 0x86, 0xc6, 0xbd,
		0x41, 0xe9, 0x5d, 0xa8, 0xab, 0x43, 0xa5, 0x11, 0x8a, 0xd4, 0xdd, 0xcc,
		0xb5, 0x64, 0xf5, 0xce, 0x10, 0xca, 0x9c, 0x34, 0x7d, 0x4b, 0xef, 0x24,
		0x7d, 0x3e, 0x37, 0x53, 0x1c, 0x65, 0x37, 0x6f, 0x45, 0x84, 0xf9, 0xc6,
		0xcb, 0x6b, 0x13, 0x0a, 0x47, 0xde, 0x93, 0xfe, 0xe5, 0x70, 0x38, 0xe8,
		0xdf, 0x9c, 0x0f, 0xff, 0x46, 0x0e, 0x81, 0x5c, 0x5e, 0xe1, 0x59, 0x02,
		0xaf, 0xfd, 0x8b, 0xcb, 0xeb, 0xf4, 0x96, 0xfe, 0x39, 0x38, 0x23, 0x1f,
		0x8d, 0x32, 0xf3, 0x1c, 0xd3, 0x5e, 0x5d, 0x4f, 0x33, 0xd4, 0x92, 0x70,
		0x40, 0xc0, 0xb3, 0x67, 0x2e, 0x26, 0x0f, 0x19, 0x0d, 0x4d, 0x50, 0x06,
		0x9d, 0x2e, 0x3c, 0xcf, 0x64, 0xa2, 0x79, 0xa9, 0x9d, 0xc4, 0x66, 0x34,
		0x9b, 0x91, 0x85, 0x5e, 0xcd, 0x12, 0x95, 0x16, 0x53, 0xaa, 0x82, 0xb4,
		0x38, 0x9d, 0x32, 0x42, 0xf8, 0x95, 0x5c, 0xea, 0xe0, 0xc8, 0x24, 0xb6,
		0xc0, 0xb9, 0x6d, 0x72, 0xda, 0xd7, 0xed, 0x5c, 0x3b, 0x7c, 0xf3, 0x8e,
		0x86, 0x13, 0x04, 0x13, 0x86, 0x09, 0x84, 0xc7, 0x59, 0x8e, 0x5b, 0x9a,
		0xb3, 0x9b, 0x76, 0xe5, 0xd9, 0x60, 0x73, 0x97, 0x00, 0x86, 0x45, 0x5f,
		0x39, 0xc5, 0xda, 0xde, 0xbd, 0x65, 0x93, 0xc1, 0xaf, 0x54, 0x67, 0xb5,
		0xf4, 0x1c, 0x61, 0xce, 0x0c, 0x1c, 0x48, 0xad, 0xdc, 0x77, 0xb5, 0x4e,
		0x49, 0x4a, 0xa9, 0xbb, 0x26, 0x3c, 0xe1, 0x28, 0xfe, 0x52, 0x7b, 0x63,
		0xf7, 0x48, 0x42, 0x96, 0x57, 0x25, 0x3f, 0xc5, 0x62, 0x6e, 0x68, 0xb5,
		0xac, 0xaf, 0xf7, 0x55, 0x1a, 0xf8, 0x78, 0xbc, 0x3b, 0x42, 0xae, 0xdd,
		0x8d, 0x00, 0xf9, 0x88, 0xed, 0x93, 0xeb, 0x14, 0xa4, 0xe2, 0x31, 0x43,
		0xd8, 0x43, 0x87, 0x5c, 0x45, 0xb7, 0xcb, 0xf4, 0x50, 0xa5, 0x26, 0x14,
		0x8f, 0x5a, 0x78, 0x78, 0x11, 0xac, 0xd0, 0xda, 0x3a, 0xfd, 0x9e, 0x44,
		0xfc, 0xf5, 0xda, 0x27, 0xfc, 0x4c, 0x58, 0x1c, 0xf3, 0x5a, 0xde, 0x96,
		0x37, 0xd8, 0x9b, 0x32, 0xa9, 0xcf, 0xa6, 0xd5, 0x04, 0x6a, 0xbd, 0x98,
		0x1b, 0x5e, 0x48, 0x15, 0x85, 0x6e, 0x17, 0xdc, 0x79, 0xb3, 0xed, 0x7a,
		0xa3, 0x4e, 0x7f, 0x8d, 0x98, 0x83, 0x03, 0x87, 0x9a, 0xbd, 0xbc, 0xb8,
		0x86, 0xa8, 0x25, 0x3d, 0xb0, 0xa5, 0xba, 0x34, 0x41, 0x27, 0xc1, 0x5d,
		0xc2, 0x17, 0x31, 0x0b, 0x31, 0x0f, 0xd7, 0x4c, 0xeb, 0x88, 0xc6, 0xb1,
		0x65, 0xce, 0x5a, 0x6f, 0x27, 0x69, 0xc1, 0xe2, 0x80, 0xeb, 0x52, 0x4b,
		0x4f, 0x12, 0x56, 0xd5, 0x15, 0x99, 0xe7, 0xc6, 0x59, 0xee, 0xff, 0xbb,
		0x3c, 0x07, 0x0f, 0x91, 0x34, 0x05, 0x50, 0x91, 0x6c, 0x25, 0x30, 0xa2,
		0x81, 0x8d, 0x39, 0x0f, 0x47, 0x4b, 0x3c, 0x13, 0xff, 0x06, 0xb0, 0x56,
		0xeb, 0x07, 0x80, 0x85, 0x91, 0xac, 0xc4, 0xf6, 0xe8, 0x78, 0xfd, 0xb3,
		0xff, 0xf6, 0xac, 0x6d, 0x36, 0x9a, 0x02, 0x6d, 0xc3, 0x39, 0xac, 0x6d,
		0xd3, 0x63, 0x10, 0x73, 0xb9, 0x43, 0x8d, 0x7f, 0xc8, 0x72, 0xae, 0xb5,
		0x6a, 0xd6, 0x77, 0xad, 0x0a, 0x4a, 0xb5, 0x7d, 0xb6, 0x85, 0xb4, 0xd5,
		0xc6, 0xa6, 0x5b, 0x50, 0xfb, 0x7b, 0x7b, 0xd2, 0x66, 0x54, 0x5c, 0xbd,
		0x1f, 0xeb, 0x2a, 0xa1, 0xeb, 0xb2, 0xa3, 0xb3, 0xa4, 0x72, 0xcb, 0xe7,
		0x74, 0xad, 0xcc, 0x2a, 0xd6, 0xb7, 0xec, 0x11, 0xb6, 0x9c, 0xb9, 0xda,
		0x14, 0xf7, 0x6a, 0xe3, 0xfc, 0xa1, 0x23, 0x6d, 0x27, 0x61, 0xcb, 0xa6,
		0x65, 0x81, 0x55, 0x2e, 0x9e, 0x5e, 0xb0, 0x6c, 0x27, 0x4d, 0x25, 0x68,
		0xa0, 0x5c, 0x0f, 0x86, 0x37, 0x56, 0xa2, 0x3f, 0xe1, 0xfa, 0x73, 0x19,
		0x15, 0xe7, 0x9d, 0x8e, 0x6f, 0xdf, 0x6a, 0x74, 0x46, 0x3c, 0x5c, 0xea,
		0x43, 0x4f, 0x67, 0xf2, 0xc2, 0xbe, 0xda, 0x28, 0x9e, 0x55, 0xe4, 0xda,
		0xcb, 0x0d, 0xb4, 0xd0, 0xf4, 0x76, 0x66, 0xbd, 0x8e, 0x54, 0x82, 0x27,
		0xe3, 0x5e, 0xb6, 0x4f, 0x22, 0x45, 0x56, 0xd1, 0x6d, 0x8c, 0x60, 0x87,
		0xa0, 0x23, 0x67, 0x14, 0x9f, 0x0c, 0xc3, 0x2e, 0x29, 0x74, 0x41, 0x7a,
		0x2d, 0x34, 0xc0, 0xfb, 0xee, 0xb0, 0xd5, 0x07, 0xe9, 0x9d, 0x87, 0x3a,
		0x88, 0x19, 0xed, 0xf8, 0xb3, 0x9e, 0x8d, 0xa4, 0xc1, 0x77, 0x62, 0x3a,
		0x62, 0x78, 0x62, 0xe7, 0xa2, 0x4b, 0x34, 0xcd, 0x24, 0x8f, 0x5f, 0x7a,
		0x14, 0xcb, 0x43, 0x77, 0x7c, 0xe3, 0x60, 0x7d, 0xa3, 0x64, 0x36, 0x57,
		0x26, 0x8e, 0x71, 0x05, 0xb5, 0x9c, 0xb1, 0x2e, 0xd1, 0x44, 0x12, 0x90,
		0xd1, 0x57, 0xfc, 0xfd, 0xe7, 0x23, 0x02, 0x86, 0xd1, 0x2e, 0xa1, 0xe1,
		0x84, 0x07, 0xe9, 0x4e, 0x4f, 0x5c, 0x7f, 0xe3, 0x35, 0x9a, 0x2b, 0x85,
		0x99, 0xa6, 0xb6, 0x69, 0xee, 0x04, 0x74, 0xa3, 0x88, 0x82, 0xbb, 0x2e,
		0xc9, 0xab, 0xe5, 0x7b, 0xae, 0x45, 0x25, 0x38, 0xde, 0x6e, 0xb9, 0x99,
		0x09, 0xd6, 0x48, 0x48, 0x93, 0x73, 0xde, 0x3c, 0xe1, 0xc4, 0x45, 0xd2,
		0xd6, 0x46, 0x73, 0x2a, 0xd5, 0x32, 0xc6, 0x20, 0x53, 0x2a, 0xc6, 0x51,
		0xd2, 0x86, 0x23, 0x78, 0x71, 0x34, 0x7b, 0x38, 0x46, 0xd2, 0x2c, 0xb7,
		0xeb, 0x8c, 0xda, 0x07, 0xcc, 0x9e, 0x39, 0x1f, 0x65, 0xcf, 0xb6, 0x5b,
		0x18, 0xb4, 0xbc, 0x09, 0x1a, 0x46, 0x9c, 0x18, 0x3a, 0xad, 0x9f, 0x69,
		0x00, 0x96, 0xdb, 0x54, 0xfd, 0xda, 0x30, 0x27, 0xb4, 0x7c, 0xea, 0x22,
		0x90, 0x9e, 0xfe, 0x7e, 0x1d, 0x6d, 0xf1, 0x70, 0xdc, 0x1b, 0xda, 0xd7,
		0x49, 0x3f, 0x83, 0xd9, 0x79, 0xb4, 0xde, 0x8d, 0x7b, 0xf3, 0xc0, 0xb9,
		0x6d, 0x49, 0x6c, 0xc0, 0x6c, 0x65, 0xae, 0xf4, 0x7e, 0x64, 0x5a, 0xee,
		0xd6, 0x2a, 0xc8, 0x1b, 0x72, 0x51, 0x04, 0x15, 0x33, 0xda, 0x4c, 0x8d,
		0x44, 0x1d, 0xa5, 0x27, 0x21, 0xbc, 0xb6, 0x07, 0xa4, 0x0a, 0xa5, 0x8f,
		0x84, 0xbd, 0x6a, 0x47, 0x6c, 0x9e, 0xd4, 0x44, 0xb5, 0x42, 0x77, 0x45,
		0x2f, 0xf8, 0x42, 0x76, 0xc9, 0x4b, 0x5c, 0x0b, 0x7c, 0x54, 0x36, 0xea,
		0xc7, 0x69, 0x32, 0x9f, 0xf2, 0x4c, 0x95, 0x02, 0xd6, 0x68, 0x1c, 0xe9,
		0xda, 0xf6, 0x58, 0x66, 0x08, 0xbb, 0xb7, 0x55, 0xc9, 0x58, 0x6b, 0x6b,
		0xdc, 0xc6, 0x8e, 0x83, 0xf7, 0x74, 0xe3, 0xb1, 0x1d, 0x07, 0xdb, 0x89,
		0x79, 0xbb, 0xfa, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x4f, 0xba,
		0xc3, 0x98, 0x15, 0x00, 0x00,
	},
		"_templates/console.html",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"_templates/console.html": templates_console_html,
}
