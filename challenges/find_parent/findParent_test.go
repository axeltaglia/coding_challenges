package main

import (
	"testing"
)

func TestCases(t *testing.T) {
	t.Run("testGetParentFilePath", testGetParentFilePath)
	t.Run("testGetBasePath", testGetBasePath)
	t.Run("testFindParent", testFindParent)
}

func testGetParentFilePath(t *testing.T) {
	root := NewFile("root")
	a := NewFile("a")
	b := NewFile("b")
	c := NewFile("c")
	d := NewFile("d")
	e := NewFile("e")

	root.addChild(a)
	root.addChild(b)
	a.addChild(c)
	a.addChild(d)
	c.addChild(e)

	result := getParentFilePath(root, e, "")

	if result == nil {
		t.Errorf("Expected: '/a/c/' Got: [nil]")

	}
	expected := "/a/c/"
	if *result != expected {
		t.Errorf("Expected: %s Got: '%s'", expected, *result)
	}
}

func testGetBasePath(t *testing.T) {
	path1 := "/a/b"
	path2 := "/a/c"
	result := getBasePath(path1, path2)

	expected := "/a/"
	if result != expected {
		t.Errorf("Expected: %s Got: '%s'", expected, result)
	}
}

func testFindParent(t *testing.T) {
	root := NewFile("root")
	a := NewFile("a")
	b := NewFile("b")
	c := NewFile("c")
	d := NewFile("d")
	e := NewFile("e")

	root.addChild(a)
	root.addChild(b)
	a.addChild(c)
	a.addChild(d)
	c.addChild(e)

	result := findParent(root, b, d)

	expected := "/"
	if result != expected {
		t.Errorf("Expected: %s Got: '%s'", expected, result)
	}
}
