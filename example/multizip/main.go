// Copyright 2016 the Go-FUSE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/hanwen/go-fuse/fs"
	"github.com/hanwen/go-fuse/zipfs"
)

func main() {
	// Scans the arg list and sets up flags
	debug := flag.Bool("debug", false, "debug on")
	flag.Parse()
	if flag.NArg() < 1 {
		_, prog := filepath.Split(os.Args[0])
		fmt.Printf("usage: %s MOUNTPOINT\n", prog)
		os.Exit(2)
	}

	root := &zipfs.MultiZipFs{}
	sec := time.Second
	opts := fs.Options{
		EntryTimeout:       &sec,
		AttrTimeout:        &sec,
		DefaultPermissions: true,
	}
	opts.Debug = *debug
	server, err := fs.Mount(flag.Arg(0), root, &opts)
	if err != nil {
		fmt.Printf("Mount fail: %v\n", err)
		os.Exit(1)
	}

	server.Serve()
}
