package main

import "fmt"

func main() {
	root := NewFile("root")
	a := NewFile("a")
	b := NewFile("b")
	c := NewFile("c")
	d := NewFile("d")

	root.addChild(a)
	root.addChild(b)
	a.addChild(c)
	a.addChild(d)

	fmt.Println(findParent(root, a, b))
	//-> /

	fmt.Println(findParent(root, c, d))
	//-> /a/
}

type File struct {
	children []*File
	name     string
}

func (f *File) addChild(child *File) {
	f.children = append(f.children, child)
}

func NewFile(name string) *File {
	file := new(File)
	file.name = name
	return file
}

func findParent(root, a, b *File) string {
	pathAChan := make(chan string)
	pathBChan := make(chan string)

	go func() {
		pathA := getParentFilePath(root, a, "")
		pathAChan <- *pathA
	}()

	go func() {
		pathB := getParentFilePath(root, b, "")
		pathBChan <- *pathB
	}()

	pathA := <-pathAChan
	pathB := <-pathBChan

	close(pathAChan)
	close(pathBChan)

	return getBasePath(pathA, pathB)
}

func getParentFilePath(root, target *File, path string) *string {
	for _, child := range root.children {
		if child == target {
			newPath := path + "/"
			return &newPath
		}
	}

	for _, child := range root.children {
		return getParentFilePath(child, target, path+"/"+child.name)
	}

	return nil
}

func getBasePath(path1, path2 string) string {
	basePath := ""
	for i := 0; i < len(path1); i++ {
		if path1[i] == path2[i] {
			basePath = basePath + string(path1[i])
		} else {
			break
		}

	}
	return basePath
}
