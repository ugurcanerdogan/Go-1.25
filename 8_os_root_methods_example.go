package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	// The data/ folder and 01.txt file are assumed to have been created beforehand
	root, err := os.OpenRoot("/Users/uqi/Desktop/go125/data")
	if err != nil {
		panic(err)
	}

	// Chmod
	root.Chmod("01.txt", 0600)
	finfo, _ := root.Stat("01.txt")
	fmt.Println("Chmod:", finfo.Mode().Perm())

	// Chown
	root.Chown("01.txt", 1000, 1000)
	finfo, _ = root.Stat("01.txt")
	stat := finfo.Sys().(*syscall.Stat_t)
	fmt.Printf("Chown: uid=%d, gid=%d\n", stat.Uid, stat.Gid)

	// Chtimes
	mtime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	atime := time.Now()
	root.Chtimes("01.txt", atime, mtime)
	finfo, _ = root.Stat("01.txt")
	fmt.Println("Chtimes:", finfo.ModTime())

	// Link
	root.Link("01.txt", "hardlink.txt")
	finfo, _ = root.Stat("hardlink.txt")
	fmt.Println("Link:", finfo.Name())

	// MkdirAll
	const dname = "path/to/secret"
	root.MkdirAll(dname, 0750)
	finfo, _ = root.Stat(dname)
	fmt.Println("MkdirAll:", dname, finfo.Mode())

	// RemoveAll
	root.RemoveAll("hardlink.txt")
	finfo, err = root.Stat("hardlink.txt")
	fmt.Println("RemoveAll:", finfo, err)

	// Rename
	const oldname = "01.txt"
	const newname = "go.txt"
	root.Rename(oldname, newname)
	_, err = root.Stat(oldname)
	fmt.Println("Rename (old):", err)
	finfo, _ = root.Stat(newname)
	fmt.Println("Rename (new):", finfo.Name())

	// Symlink & Readlink
	const lname = "symlink.txt"
	root.Symlink(newname, lname)
	lpath, _ := root.Readlink(lname)
	fmt.Println("Symlink:", lname, "->", lpath)

	// WriteFile & ReadFile
	const fname = "go.txt"
	root.WriteFile(fname, []byte("go is awesome"), 0644)
	content, _ := root.ReadFile(fname)
	fmt.Printf("WriteFile/ReadFile: %s: %s\n", fname, content)
}
