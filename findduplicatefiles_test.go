package main

import (
	"reflect"
	"testing"
)

const dir string = "test/test_data"

func TestFindDuplicateFiles(t *testing.T) {
	expected := make([][]string, 0)
	expected = append(expected, []string{"test/test_data/test1.txt", "test/test_data/test2.txt"})

	duplicates := FindDuplicateFiles(dir, 1)

	if len(duplicates) != len(expected) {
		t.Errorf("Differing slice length returned. Expected %v got %v", len(expected), len(duplicates))
	}

	for _, dups := range expected {
		if !inSlice(duplicates, dups) {
			t.Errorf("FindDuplicateFiles failed: Expected %v got %v", expected, duplicates)
		}
	}
}

func TestFindDuplicatesBySize(t *testing.T) {
	expected := make(map[int64][]string)
	expected[24] = []string{"test/test_data/test1.txt", "test/test_data/test2.txt", "test/test_data/test3.txt"}
	expected[201706] = []string{"test/test_data/out.gif"}

	fileSizes := findDuplicatesBySize(dir)

	if len(fileSizes) != len(expected) {
		t.Errorf("Differing map length returned. Expected %v got %v", len(expected), len(fileSizes))
	}

	if !reflect.DeepEqual(expected[24], fileSizes[24]) {
		t.Errorf("Differing file sizes returned. Expected %v got %v", expected[24], fileSizes[24])
	}

	if !reflect.DeepEqual(expected[201706], fileSizes[201706]) {
		t.Errorf("Differing file sizes returned. Expected %v got %v", expected[201706], fileSizes[201706])
	}
}

func TestFindDuplicatesByHash(t *testing.T) {
	expected := make([][]string, 0)
	expected = append(expected, []string{"test/test_data/test1.txt", "test/test_data/test2.txt"})

	files := []string{"test/test_data/test1.txt", "test/test_data/test2.txt", "test/test_data/test3.txt"}
	duplicates := findDuplicatesByHash(files, 1, 24)

	if len(duplicates) != len(expected) {
		t.Errorf("Differing slice length returned. Expected %v got %v", len(expected), len(duplicates))
	}

	for _, dups := range expected {
		if !inSlice(duplicates, dups) {
			t.Errorf("findDuplicatesByHash failed: Expected %v got %v", expected, duplicates)
		}
	}
}

func TestGenerateHash(t *testing.T) {
	file := "test/test_data/test1.txt"
	expected := "525c67bc7b921a252bf84b04d6bf97c60cea73689b185fe0f0c309efd53be0f0"

	generated := generateHash(file, 24)

	if generated != expected {
		t.Errorf("generateHash failed: Expected %v got %v", expected, generated)
	}

	expected = "4a1e67f2fe1d1cc7b31d0ca2ec441da4778203a036a77da10344c85e24ff0f92"

	generated = generateHash(file, 12)

	if generated != expected {
		t.Errorf("generateHash failed: Expected %v got %v", expected, generated)
	}
}

func inSlice(s1 [][]string, s2 []string) bool {

	for _, s := range s1 {
		if reflect.DeepEqual(s, s2) {
			return true
		}
	}
	return false
}
