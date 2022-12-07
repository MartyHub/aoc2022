package main

import (
	"fmt"
	"strings"
)

type NodeType int

const (
	DirType NodeType = iota
	FileType
)

type Node interface {
	Name() string
	Size() int
	String() string
	Type() NodeType
}

type Dir struct {
	name     string
	children []Node
	parent   *Dir
}

func (d *Dir) AddDir(name string) *Dir {
	result := &Dir{
		name:   name,
		parent: d,
	}

	d.children = append(d.children, result)

	return result
}

func (d *Dir) AddFile(name string, size int) {
	file := &File{
		name: name,
		size: size,
	}

	d.children = append(d.children, file)
}

func (d *Dir) Do(f func(*Dir)) {
	for _, child := range d.children {
		if child.Type() == DirType {
			dir := child.(*Dir)

			f(dir)

			dir.Do(f)
		}
	}
}

func (d *Dir) Name() string {
	return d.name
}

func (d *Dir) Size() int {
	size := 0

	for _, child := range d.children {
		size += child.Size()
	}

	return size
}

func (d *Dir) String() string {
	sw := strings.Builder{}

	sw.WriteString(fmt.Sprintf("%v (dir, size=%v)\n", d.name, d.Size()))

	for _, child := range d.children {
		sw.WriteString(child.String())
	}

	return sw.String()
}

func (d *Dir) Type() NodeType {
	return DirType
}

type File struct {
	name string
	size int
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Size() int {
	return f.size
}

func (f *File) String() string {
	return fmt.Sprintf("%v (file, size= %v)\n", f.name, f.size)
}

func (f *File) Type() NodeType {
	return FileType
}
