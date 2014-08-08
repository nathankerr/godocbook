package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file_2 := &embedded.EmbeddedFile{
		Filename:    `book.tex`,
		FileModTime: time.Unix(1407537166, 0),
		Content:     string([]byte{0x5c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5b, 0x61, 0x35, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2c, 0x31, 0x32, 0x70, 0x74, 0x2c, 0x74, 0x77, 0x6f, 0x73, 0x69, 0x64, 0x65, 0x2c, 0x6f, 0x6e, 0x65, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x2c, 0x6f, 0x70, 0x65, 0x6e, 0x72, 0x69, 0x67, 0x68, 0x74, 0x5d, 0x7b, 0x62, 0x6f, 0x6f, 0x6b, 0x7d, 0xa, 0xa, 0x5c, 0x75, 0x73, 0x65, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x7b, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x70, 0x65, 0x63, 0x7d, 0xa, 0x5c, 0x73, 0x65, 0x74, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x6f, 0x6e, 0x74, 0x7b, 0x41, 0x6c, 0x65, 0x67, 0x72, 0x65, 0x79, 0x61, 0x7d, 0xa, 0x5c, 0x73, 0x65, 0x74, 0x6d, 0x6f, 0x6e, 0x6f, 0x66, 0x6f, 0x6e, 0x74, 0x7b, 0x49, 0x6e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x61, 0x7d, 0xa, 0x5c, 0x73, 0x65, 0x74, 0x73, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x6e, 0x74, 0x7b, 0x41, 0x76, 0x65, 0x72, 0x69, 0x61, 0x20, 0x53, 0x61, 0x6e, 0x73, 0x7d, 0xa, 0xa, 0x5c, 0x75, 0x73, 0x65, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5b, 0xa, 0x9, 0x61, 0x35, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2c, 0xa, 0x9, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x68, 0x65, 0x61, 0x64, 0x66, 0x6f, 0x6f, 0x74, 0x2c, 0xa, 0x9, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x3d, 0x7b, 0x33, 0x32, 0x34, 0x62, 0x70, 0x2c, 0x20, 0x35, 0x32, 0x34, 0x62, 0x70, 0x7d, 0x2c, 0xa, 0x9, 0x25, 0x20, 0x73, 0x68, 0x6f, 0x77, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2c, 0xa, 0x9, 0x68, 0x65, 0x61, 0x64, 0x73, 0x65, 0x70, 0x3d, 0x5c, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x6b, 0x69, 0x70, 0x2c, 0xa, 0x9, 0x68, 0x65, 0x61, 0x64, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3d, 0x31, 0x35, 0x70, 0x74, 0x2c, 0xa, 0x9, 0x66, 0x6f, 0x6f, 0x74, 0x73, 0x6b, 0x69, 0x70, 0x3d, 0x32, 0x5c, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x6b, 0x69, 0x70, 0x2c, 0xa, 0x9, 0x74, 0x77, 0x6f, 0x73, 0x69, 0x64, 0x65, 0x2c, 0xa, 0x9, 0x76, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x3d, 0x37, 0x30, 0x3a, 0x39, 0x39, 0x2c, 0x20, 0x25, 0x20, 0x7e, 0x31, 0x3a, 0x73, 0x71, 0x72, 0x74, 0x28, 0x32, 0x29, 0xa, 0x9, 0x68, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x3d, 0x39, 0x39, 0x3a, 0x37, 0x30, 0x2c, 0xa, 0x9, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x70, 0x61, 0x72, 0x3d, 0x32, 0x37, 0x70, 0x74, 0x2c, 0xa, 0x5d, 0x7b, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x7d, 0xa, 0xa, 0x25, 0x20, 0x73, 0x65, 0x74, 0x75, 0x70, 0x20, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x73, 0xa, 0x5c, 0x75, 0x73, 0x65, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x7b, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x68, 0x64, 0x72, 0x7d, 0xa, 0x5c, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x70, 0x61, 0x67, 0x65, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x7b, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x7d, 0x7b, 0x25, 0xa, 0x9, 0x5c, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x68, 0x66, 0x7b, 0x7d, 0xa, 0x9, 0x5c, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x66, 0x6f, 0x6f, 0x74, 0x5b, 0x43, 0x5d, 0x7b, 0x5c, 0x73, 0x66, 0xe2, 0x80, 0x93, 0x7e, 0x5c, 0x74, 0x68, 0x65, 0x70, 0x61, 0x67, 0x65, 0x7e, 0xe2, 0x80, 0x93, 0x7d, 0xa, 0x9, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x68, 0x65, 0x61, 0x64, 0x72, 0x75, 0x6c, 0x65, 0x77, 0x69, 0x64, 0x74, 0x68, 0x7b, 0x30, 0x70, 0x74, 0x7d, 0xa, 0x9, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x66, 0x6f, 0x6f, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x77, 0x69, 0x64, 0x74, 0x68, 0x7b, 0x30, 0x70, 0x74, 0x7d, 0xa, 0x7d, 0xa, 0xa, 0x5c, 0x6c, 0x65, 0x74, 0x5c, 0x4f, 0x6c, 0x64, 0x70, 0x61, 0x72, 0x74, 0x5c, 0x70, 0x61, 0x72, 0x74, 0xa, 0x5c, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x70, 0x61, 0x72, 0x74, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x7b, 0x7d, 0xa, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x70, 0x61, 0x72, 0x74, 0x5b, 0x31, 0x5d, 0x7b, 0x5c, 0x4f, 0x6c, 0x64, 0x70, 0x61, 0x72, 0x74, 0x7b, 0x23, 0x31, 0x7d, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x70, 0x61, 0x72, 0x74, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x7b, 0x23, 0x31, 0x7e, 0xe2, 0x80, 0x93, 0x7e, 0x7d, 0x7d, 0xa, 0x5c, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x7b, 0x7d, 0xa, 0xa, 0x5c, 0x70, 0x61, 0x67, 0x65, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x7b, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x7d, 0xa, 0x5c, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x68, 0x66, 0x7b, 0x7d, 0xa, 0x5c, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x68, 0x65, 0x61, 0x64, 0x5b, 0x43, 0x4f, 0x5d, 0x7b, 0x5c, 0x73, 0x66, 0x5c, 0x6c, 0x65, 0x66, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x7d, 0xa, 0x5c, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x68, 0x65, 0x61, 0x64, 0x5b, 0x43, 0x45, 0x5d, 0x7b, 0x5c, 0x73, 0x66, 0x20, 0x5c, 0x70, 0x61, 0x72, 0x74, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x20, 0x5c, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x7d, 0xa, 0x5c, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x66, 0x6f, 0x6f, 0x74, 0x5b, 0x43, 0x5d, 0x7b, 0x5c, 0x73, 0x66, 0xe2, 0x80, 0x93, 0x7e, 0x5c, 0x74, 0x68, 0x65, 0x70, 0x61, 0x67, 0x65, 0x7e, 0xe2, 0x80, 0x93, 0x7d, 0xa, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x68, 0x65, 0x61, 0x64, 0x72, 0x75, 0x6c, 0x65, 0x77, 0x69, 0x64, 0x74, 0x68, 0x7b, 0x30, 0x70, 0x74, 0x7d, 0xa, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x66, 0x6f, 0x6f, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x77, 0x69, 0x64, 0x74, 0x68, 0x7b, 0x30, 0x70, 0x74, 0x7d, 0xa, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2a, 0x7b, 0x5c, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x7d, 0x5b, 0x31, 0x5d, 0x7b, 0x5c, 0x6d, 0x61, 0x72, 0x6b, 0x62, 0x6f, 0x74, 0x68, 0x7b, 0x5c, 0x74, 0x68, 0x65, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x7e, 0x23, 0x31, 0x7d, 0x7b, 0x7d, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x7b, 0x23, 0x31, 0x7d, 0x7d, 0xa, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2a, 0x7b, 0x5c, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x7d, 0x5b, 0x31, 0x5d, 0x7b, 0x5c, 0x6d, 0x61, 0x72, 0x6b, 0x62, 0x6f, 0x74, 0x68, 0x7b, 0x5c, 0x74, 0x68, 0x65, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7e, 0x23, 0x31, 0x7d, 0x7b, 0x7d, 0x7d, 0xa, 0xa, 0x5c, 0x75, 0x73, 0x65, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x7b, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x78, 0x7d, 0xa, 0x5c, 0x73, 0x65, 0x74, 0x6b, 0x65, 0x79, 0x73, 0x7b, 0x47, 0x69, 0x6e, 0x7d, 0x7b, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3d, 0x5c, 0x6c, 0x69, 0x6e, 0x65, 0x77, 0x69, 0x64, 0x74, 0x68, 0x2c, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3d, 0x30, 0x2e, 0x39, 0x31, 0x35, 0x5c, 0x74, 0x65, 0x78, 0x74, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2c, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x73, 0x70, 0x65, 0x63, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x7d, 0xa, 0xa, 0x5c, 0x75, 0x73, 0x65, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5d, 0x7b, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x7d, 0xa, 0x5c, 0x6c, 0x73, 0x74, 0x73, 0x65, 0x74, 0x7b, 0xa, 0x9, 0x62, 0x61, 0x73, 0x69, 0x63, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x5c, 0x66, 0x6f, 0x6f, 0x74, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x2c, 0xa, 0x9, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x3d, 0x6c, 0x65, 0x66, 0x74, 0x2c, 0xa, 0x9, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x3d, 0x74, 0x72, 0x75, 0x65, 0x2c, 0xa, 0x9, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x3d, 0x6c, 0x65, 0x66, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x2c, 0xa, 0x9, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x6f, 0x73, 0x3d, 0x62, 0x2c, 0xa, 0x9, 0x74, 0x61, 0x62, 0x73, 0x69, 0x7a, 0x65, 0x3d, 0x32, 0x2c, 0xa, 0x9, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x61, 0x74, 0x77, 0x68, 0x69, 0x74, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3d, 0x74, 0x72, 0x75, 0x65, 0x2c, 0xa, 0x7d, 0xa, 0x5c, 0x64, 0x65, 0x66, 0x5c, 0x63, 0x6f, 0x64, 0x65, 0x23, 0x31, 0x7b, 0x5c, 0x6d, 0x62, 0x6f, 0x78, 0x7b, 0x5c, 0x6c, 0x73, 0x74, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x7b, 0x23, 0x31, 0x7d, 0x7d, 0x7d, 0xa, 0xa, 0x5c, 0x75, 0x73, 0x65, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x7b, 0x61, 0x6d, 0x73, 0x6d, 0x61, 0x74, 0x68, 0x7d, 0xa, 0x5c, 0x75, 0x73, 0x65, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x7b, 0x61, 0x6d, 0x73, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x7d, 0xa, 0x5c, 0x6e, 0x65, 0x77, 0x74, 0x68, 0x65, 0x6f, 0x72, 0x65, 0x6d, 0x7b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x7b, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0xa, 0xa, 0x5c, 0x75, 0x73, 0x65, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x7b, 0x63, 0x69, 0x74, 0x65, 0x7d, 0xa, 0x5c, 0x75, 0x73, 0x65, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x7b, 0x75, 0x72, 0x6c, 0x7d, 0xa, 0x5c, 0x75, 0x73, 0x65, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x7b, 0x68, 0x79, 0x70, 0x65, 0x72, 0x72, 0x65, 0x66, 0x7d, 0xa, 0x5c, 0x75, 0x73, 0x65, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x7b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0xa, 0x5c, 0x75, 0x73, 0x65, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x7b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x7d, 0xa, 0xa, 0x25, 0x20, 0x66, 0x69, 0x78, 0x20, 0x27, 0x62, 0x6c, 0x61, 0x6e, 0x6b, 0x27, 0x20, 0x70, 0x61, 0x67, 0x65, 0x73, 0x20, 0x61, 0x66, 0x74, 0x65, 0x72, 0x20, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0xa, 0x5c, 0x6c, 0x65, 0x74, 0x5c, 0x6f, 0x72, 0x69, 0x67, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x70, 0x61, 0x67, 0x65, 0x5c, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x70, 0x61, 0x67, 0x65, 0xa, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x70, 0x61, 0x67, 0x65, 0x7b, 0x25, 0xa, 0x9, 0x5c, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x70, 0x61, 0x67, 0x65, 0xa, 0x9, 0x7b, 0x5c, 0x70, 0x61, 0x67, 0x65, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x7b, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x7d, 0x5c, 0x6f, 0x72, 0x69, 0x67, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x70, 0x61, 0x67, 0x65, 0x7d, 0xa, 0x7d, 0xa, 0xa, 0x25, 0x25, 0x20, 0x73, 0x65, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x73, 0xa, 0x5c, 0x6d, 0x61, 0x6b, 0x65, 0x61, 0x74, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0xa, 0x5c, 0x64, 0x65, 0x66, 0x5c, 0x40, 0x6d, 0x61, 0x6b, 0x65, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x68, 0x65, 0x61, 0x64, 0x23, 0x31, 0x7b, 0x25, 0xa, 0x9, 0x5c, 0x76, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2a, 0x7b, 0x32, 0x30, 0x5c, 0x70, 0x40, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x5c, 0x70, 0x61, 0x72, 0x69, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x20, 0x5c, 0x7a, 0x40, 0x20, 0x5c, 0x72, 0x61, 0x67, 0x67, 0x65, 0x64, 0x72, 0x69, 0x67, 0x68, 0x74, 0x20, 0x5c, 0x73, 0x66, 0x5c, 0x48, 0x75, 0x67, 0x65, 0xa, 0x9, 0x5c, 0x69, 0x66, 0x6e, 0x75, 0x6d, 0x20, 0x5c, 0x63, 0x40, 0x73, 0x65, 0x63, 0x6e, 0x75, 0x6d, 0x64, 0x65, 0x70, 0x74, 0x68, 0x20, 0x3e, 0x5c, 0x6d, 0x40, 0x6e, 0x65, 0xa, 0x9, 0x9, 0x5c, 0x69, 0x66, 0x40, 0x6d, 0x61, 0x69, 0x6e, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0xa, 0x9, 0x9, 0x9, 0x7b, 0x5c, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x70, 0x65, 0x63, 0x5b, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x3d, 0x34, 0x2e, 0x30, 0x5d, 0x7b, 0x41, 0x76, 0x65, 0x72, 0x69, 0x61, 0x20, 0x53, 0x61, 0x6e, 0x73, 0x7d, 0x5c, 0x74, 0x68, 0x65, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x7d, 0xa, 0x9, 0x9, 0x5c, 0x66, 0x69, 0xa, 0x9, 0x5c, 0x66, 0x69, 0xa, 0x9, 0x5c, 0x70, 0x61, 0x72, 0x62, 0x6f, 0x78, 0x7b, 0x30, 0x2e, 0x36, 0x5c, 0x6c, 0x69, 0x6e, 0x65, 0x77, 0x69, 0x64, 0x74, 0x68, 0x7d, 0x7b, 0x5c, 0x72, 0x61, 0x67, 0x67, 0x65, 0x64, 0x72, 0x69, 0x67, 0x68, 0x74, 0x5c, 0x76, 0x73, 0x6b, 0x69, 0x70, 0x2d, 0x31, 0x2e, 0x38, 0x5c, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x6b, 0x69, 0x70, 0x23, 0x31, 0x7d, 0xa, 0x9, 0x5c, 0x76, 0x73, 0x6b, 0x69, 0x70, 0x20, 0x34, 0x30, 0x70, 0x74, 0xa, 0x7d, 0x7d, 0xa, 0x5c, 0x64, 0x65, 0x66, 0x5c, 0x40, 0x6d, 0x61, 0x6b, 0x65, 0x73, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x68, 0x65, 0x61, 0x64, 0x23, 0x31, 0x7b, 0x25, 0xa, 0x9, 0x5c, 0x76, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2a, 0x7b, 0x31, 0x30, 0x5c, 0x70, 0x40, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x5c, 0x70, 0x61, 0x72, 0x69, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x20, 0x5c, 0x7a, 0x40, 0x20, 0x5c, 0x72, 0x61, 0x67, 0x67, 0x65, 0x64, 0x72, 0x69, 0x67, 0x68, 0x74, 0xa, 0x9, 0x9, 0x5c, 0x73, 0x66, 0xa, 0x9, 0x9, 0x5c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x5c, 0x40, 0x4d, 0xa, 0x9, 0x9, 0x5c, 0x48, 0x75, 0x67, 0x65, 0x20, 0x23, 0x31, 0xa, 0x9, 0x9, 0x5c, 0x76, 0x73, 0x6b, 0x69, 0x70, 0x20, 0x34, 0x30, 0x70, 0x74, 0xa, 0x7d, 0x7d, 0xa, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7b, 0x5c, 0x40, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x7b, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x7b, 0x31, 0x7d, 0x7b, 0x5c, 0x7a, 0x40, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x2d, 0x33, 0x2e, 0x35, 0x65, 0x78, 0x20, 0x5c, 0x40, 0x70, 0x6c, 0x75, 0x73, 0x20, 0x2d, 0x31, 0x65, 0x78, 0x20, 0x5c, 0x40, 0x6d, 0x69, 0x6e, 0x75, 0x73, 0x20, 0x2d, 0x2e, 0x32, 0x65, 0x78, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x32, 0x2e, 0x33, 0x65, 0x78, 0x20, 0x5c, 0x40, 0x70, 0x6c, 0x75, 0x73, 0x2e, 0x32, 0x65, 0x78, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x5c, 0x73, 0x66, 0x5c, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x7d, 0x7d, 0xa, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x73, 0x75, 0x62, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7b, 0x5c, 0x40, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7b, 0x73, 0x75, 0x62, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x7b, 0x32, 0x7d, 0x7b, 0x5c, 0x7a, 0x40, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x2d, 0x33, 0x2e, 0x32, 0x35, 0x65, 0x78, 0x5c, 0x40, 0x70, 0x6c, 0x75, 0x73, 0x20, 0x2d, 0x31, 0x65, 0x78, 0x20, 0x5c, 0x40, 0x6d, 0x69, 0x6e, 0x75, 0x73, 0x20, 0x2d, 0x2e, 0x32, 0x65, 0x78, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x31, 0x2e, 0x35, 0x65, 0x78, 0x20, 0x5c, 0x40, 0x70, 0x6c, 0x75, 0x73, 0x20, 0x2e, 0x32, 0x65, 0x78, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x5c, 0x73, 0x66, 0x5c, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x7d, 0x7d, 0xa, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x73, 0x75, 0x62, 0x73, 0x75, 0x62, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7b, 0x5c, 0x40, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7b, 0x73, 0x75, 0x62, 0x73, 0x75, 0x62, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x7b, 0x33, 0x7d, 0x7b, 0x5c, 0x7a, 0x40, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x2d, 0x33, 0x2e, 0x32, 0x35, 0x65, 0x78, 0x5c, 0x40, 0x70, 0x6c, 0x75, 0x73, 0x20, 0x2d, 0x31, 0x65, 0x78, 0x20, 0x5c, 0x40, 0x6d, 0x69, 0x6e, 0x75, 0x73, 0x20, 0x2d, 0x2e, 0x32, 0x65, 0x78, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x31, 0x2e, 0x35, 0x65, 0x78, 0x20, 0x5c, 0x40, 0x70, 0x6c, 0x75, 0x73, 0x20, 0x2e, 0x32, 0x65, 0x78, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x5c, 0x73, 0x66, 0x5c, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x73, 0x69, 0x7a, 0x65, 0x7d, 0x7d, 0xa, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x7b, 0x5c, 0x40, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7b, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x7d, 0x7b, 0x34, 0x7d, 0x7b, 0x5c, 0x7a, 0x40, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x33, 0x2e, 0x32, 0x35, 0x65, 0x78, 0x20, 0x5c, 0x40, 0x70, 0x6c, 0x75, 0x73, 0x31, 0x65, 0x78, 0x20, 0x5c, 0x40, 0x6d, 0x69, 0x6e, 0x75, 0x73, 0x2e, 0x32, 0x65, 0x78, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x2d, 0x31, 0x65, 0x6d, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x5c, 0x73, 0x66, 0x5c, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x73, 0x69, 0x7a, 0x65, 0x7d, 0x7d, 0x9, 0xa, 0x5c, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x73, 0x75, 0x62, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x7b, 0x5c, 0x40, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7b, 0x73, 0x75, 0x62, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x7d, 0x7b, 0x35, 0x7d, 0x7b, 0x5c, 0x70, 0x61, 0x72, 0x69, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x33, 0x2e, 0x32, 0x35, 0x65, 0x78, 0x20, 0x5c, 0x40, 0x70, 0x6c, 0x75, 0x73, 0x31, 0x65, 0x78, 0x20, 0x5c, 0x40, 0x6d, 0x69, 0x6e, 0x75, 0x73, 0x20, 0x2e, 0x32, 0x65, 0x78, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x2d, 0x31, 0x65, 0x6d, 0x7d, 0x25, 0xa, 0x9, 0x7b, 0x5c, 0x73, 0x66, 0x5c, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x73, 0x69, 0x7a, 0x65, 0x7d, 0x7d, 0xa, 0x5c, 0x6d, 0x61, 0x6b, 0x65, 0x61, 0x74, 0x6f, 0x74, 0x68, 0x65, 0x72, 0xa, 0xa, 0x5c, 0x64, 0x65, 0x66, 0x5c, 0x65, 0x6d, 0x70, 0x68, 0x61, 0x73, 0x69, 0x7a, 0x65, 0x23, 0x31, 0x7b, 0x7b, 0x5c, 0x62, 0x66, 0x20, 0x23, 0x31, 0x7d, 0x7d, 0xa, 0xa, 0x5c, 0x64, 0x65, 0x66, 0x5c, 0x6e, 0x6f, 0x74, 0x65, 0x23, 0x31, 0x7b, 0x5c, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x70, 0x61, 0x72, 0x7b, 0x5c, 0x74, 0x69, 0x6e, 0x79, 0x5c, 0x72, 0x61, 0x67, 0x67, 0x65, 0x64, 0x72, 0x69, 0x67, 0x68, 0x74, 0x20, 0x23, 0x31, 0x7d, 0x7d, 0xa, 0xa, 0x5c, 0x64, 0x65, 0x66, 0x5c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x70, 0x61, 0x72, 0x74, 0x23, 0x31, 0x23, 0x32, 0x23, 0x33, 0x23, 0x34, 0x7b, 0x7b, 0xa, 0x9, 0x5c, 0x74, 0x79, 0x70, 0x65, 0x6f, 0x75, 0x74, 0x7b, 0x23, 0x31, 0x20, 0x23, 0x32, 0x20, 0x23, 0x33, 0x20, 0x23, 0x34, 0x7d, 0xa, 0x9, 0x5c, 0x73, 0x75, 0x62, 0x73, 0x75, 0x62, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7b, 0x23, 0x32, 0x7d, 0xa, 0x9, 0x5c, 0x6c, 0x73, 0x74, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5b, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3d, 0x7b, 0x23, 0x32, 0x7d, 0x5d, 0x7b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2f, 0x23, 0x31, 0x2f, 0x23, 0x33, 0x2f, 0x23, 0x31, 0x23, 0x34, 0x2e, 0x73, 0x65, 0x65, 0x64, 0x7d, 0xa, 0x9, 0x5c, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x7b, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x7d, 0x5b, 0x48, 0x5d, 0xa, 0x9, 0x5c, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0xa, 0x9, 0x5c, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73, 0x7b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2f, 0x23, 0x31, 0x2f, 0x23, 0x33, 0x2f, 0x23, 0x31, 0x23, 0x34, 0x7d, 0xa, 0x9, 0x5c, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x7b, 0x23, 0x32, 0x7d, 0xa, 0x9, 0x5c, 0x65, 0x6e, 0x64, 0x7b, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x7d, 0xa, 0x7d, 0x7d, 0xa, 0xa, 0x5c, 0x64, 0x65, 0x66, 0x5c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x23, 0x31, 0x23, 0x32, 0x7b, 0x7b, 0xa, 0x9, 0x5c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x70, 0x61, 0x72, 0x74, 0x7b, 0x23, 0x31, 0x7d, 0x7b, 0x55, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x20, 0x23, 0x32, 0x7d, 0x7b, 0x6e, 0x6f, 0x6e, 0x65, 0x7d, 0x7b, 0x7d, 0xa, 0x9, 0x5c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x70, 0x61, 0x72, 0x74, 0x7b, 0x23, 0x31, 0x7d, 0x7b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x64, 0x20, 0x23, 0x32, 0x7d, 0x7b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x7d, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x7d, 0xa, 0x9, 0x5c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x70, 0x61, 0x72, 0x74, 0x7b, 0x23, 0x31, 0x7d, 0x7b, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x20, 0x23, 0x32, 0x7d, 0x7b, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x7d, 0x7b, 0x7d, 0xa, 0x9, 0x5c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x70, 0x61, 0x72, 0x74, 0x7b, 0x23, 0x31, 0x7d, 0x7b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x64, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x20, 0x23, 0x32, 0x7d, 0x7b, 0x62, 0x6f, 0x74, 0x68, 0x7d, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x7d, 0xa, 0x7d, 0x7d, 0xa, 0xa, 0x5c, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x7b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0xa, 0xa, 0x5b, 0x5b, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x2e, 0x5d, 0x5d, 0xa, 0x5b, 0x5b, 0x2e, 0x5d, 0x5d, 0xa, 0xa, 0x5b, 0x5b, 0x65, 0x6e, 0x64, 0x5d, 0x5d, 0xa, 0xa, 0x5c, 0x65, 0x6e, 0x64, 0x7b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file_3 := &embedded.EmbeddedFile{
		Filename:    `package.txt`,
		FileModTime: time.Unix(1407535603, 0),
		Content:     string([]byte{0x5b, 0x77, 0x69, 0x74, 0x68, 0x20, 0x2e, 0x50, 0x44, 0x6f, 0x63, 0x5d, 0x23, 0x20, 0x5b, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x61, 0x74, 0x68, 0x5d, 0xa, 0xa, 0x5b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x20, 0x2e, 0x44, 0x6f, 0x63, 0x20, 0x22, 0x22, 0x20, 0x22, 0x5c, 0x74, 0x22, 0x5d, 0xa, 0x5b, 0x65, 0x6e, 0x64, 0x5d, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}

	// define dirs
	dir_1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1407537166, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file_2, // book.tex
			file_3, // package.txt

		},
	}

	// link ChildDirs
	dir_1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`templates`, &embedded.EmbeddedBox{
		Name: `templates`,
		Time: time.Unix(1407537243, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir_1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"book.tex":    file_2,
			"package.txt": file_3,
		},
	})
}
