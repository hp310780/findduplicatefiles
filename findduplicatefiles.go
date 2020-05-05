package findduplicatefiles

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var dir string
	var chunk int

	flag.StringVar(&dir, "directory", "", "The directory to search for duplicate files.")
	flag.IntVar(&chunk, "chunk", 1, "Size of initial hash to check. 1 indicates the full file hash. 2 is half etc.")
	flag.Parse()

	if dir == "" {
		fmt.Println("No given directory to search.")
		os.Exit(0)
	}
	duplicates := FindDuplicateFiles(dir, chunk)
	fmt.Printf("Found the following duplicate files in directory %s :\n", dir)
	fmt.Println(duplicates)
}

func FindDuplicateFiles(dir string, chunk int) [][]string {
	// First group files by size, eliminating unique files.
	// Then, find duplicates by hash.

	// [[dup1, dup2], [dupA1, dupA2]] etc.
	duplicates := make([][]string, 0)

	fileSizes := findDuplicatesBySize(dir)

	for size, files := range fileSizes {
		if size == 0 {
			if len(files) > 1 {
				duplicates = append(duplicates, files)
			}
			continue
		}
		size := int(size)
		duplicateFiles := findDuplicatesByHash(files, chunk, size)
		if duplicateFiles != nil {
			for _, df := range duplicateFiles {
				duplicates = append(duplicates, df)
			}
		}
	}
	return duplicates
}

func findDuplicatesBySize(dir string) map[int64][]string {
	// map[size] -> [files of this size]
	fileSizes := make(map[int64][]string)
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			check(err)
			if !info.IsDir() {
				fileSizes[info.Size()] = append(fileSizes[info.Size()], path)
			}
			return nil
		})
	check(err)
	return fileSizes
}

func findDuplicatesByHash(files []string, fileChunk int, fileSize int) [][]string {
	// Given files of the same size, we now check the hash of the files (up to the chunk)
	// to see which are equal, then check the full hash if they are.

	hashes := make(map[string][]string)
	duplicates := make([][]string, 0)
	seekAt := fileSize / fileChunk

	for _, f := range files {
		fileHash := generateHash(f, seekAt)
		hashes[fileHash] = append(hashes[fileHash], f)
	}

	fullFileHashes := make(map[string][]string)

	// fileChunk 1 means we've already checked the full file hash
	if fileChunk == 1 {
		fullFileHashes = hashes
	} else {
		for _, files := range hashes {
			for _, f := range files {
				// Setting chunk to fileSize means we
				// generate a hash for the full file
				fullHash := generateHash(f, fileSize)
				fullFileHashes[fullHash] = append(fullFileHashes[fullHash], f)
			}
		}
	}

	for _, files := range fullFileHashes {
		if len(files) > 1 {
			duplicates = append(duplicates, files)
		}
	}

	return duplicates
}

func generateHash(fp string, chunk int) string {
	f, err := os.Open(fp)
	check(err)
	defer f.Close()

	// Generate hash for chunk only
	b := make([]byte, chunk)
	r := bufio.NewReader(f)

	_, err = r.Read(b)
	check(err)

	hasher := sha256.New()
	_, err = hasher.Write(b)
	check(err)

	checksum := hasher.Sum(nil)
	hexDigest := hex.EncodeToString(checksum)
	return hexDigest
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
