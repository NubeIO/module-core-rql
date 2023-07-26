package docs

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var separator string = string(filepath.Separator)

func Get(path string, deep, asTriple bool) (doc string, err error) {
	defer func() {
		if err != nil {
			et := errors.New(fmt.Sprintf("error with input data: %v %v", path, deep))
			err = et
		}
	}()
	{
		var apath string
		apath, err = filepath.Abs(path)
		if err != nil {
			return "", fmt.Errorf("cannot get absolute path: %v", err)
		}
		var st os.FileInfo
		st, err = os.Stat(apath)
		if os.IsNotExist(err) {
			return "", fmt.Errorf("cannot find: `%s`", path)
		}
		if !st.IsDir() {
			return "", fmt.Errorf("is not a folder: `%s`", path)
		}
	}
	var files []string
	{
		// List of ignore folders: vendor, .git
		ignore := []string{"vendor", ".git"}
		folderList := []string{path}
		// For avoid infinite loop added limits of search iterations(cycles).
		for iter := 0; iter < 1000000; iter++ {
			var findFolders []string
			for _, folder := range folderList {
				fileInfo, err := ioutil.ReadDir(folder)
				if err != nil {
					return "", fmt.Errorf("cannot read dir `%s`: %v", folder, err)
				}
				for _, f := range fileInfo {
					if f.IsDir() {
						isIgnore := false
						for _, ig := range ignore {
							if f.Name() == ig {
								isIgnore = true
								break
							}
						}
						if isIgnore {
							continue
						}
						findFolders = append(findFolders, f.Name())
						continue
					}
					if name := f.Name(); strings.HasSuffix(name, ".go") {
						files = append(files, folder+separator+name)
					}
				}
			}
			folderList = findFolders
			if !deep || len(folderList) == 0 {
				break
			}
		}
	}

	if len(files) == 0 {
		return "", fmt.Errorf("cannot find any files")
	}

	sort.Strings(files)
	// Reading files one by one.
	for _, filename := range files {
		var content []byte
		content, err = ioutil.ReadFile(filename)
		if err != nil {
			// If cannot read a file content, then return the error.
			return "", fmt.Errorf("cannot read file content: %v", filename)
		}
		// Read file line by line.
		lines := bytes.Split(content, []byte("\n"))
		for i := range lines {
			line := lines[i]
			var ts string = "//"
			if asTriple {
				ts = "///"
			}
			index := bytes.Index(line, []byte(ts))
			if index < 0 {
				// that line haven`t triplet-slash
				continue
			}
			if index > 0 {
				// Before triplet-slash is not acceptable any characters,
				// except `\t` or space.
				isAcceptableLine := true
				for pos := 0; pos < index; pos++ {
					if !(line[pos] == ' ' || line[pos] == '\t') {
						isAcceptableLine = false
						break
					}
				}
				if !isAcceptableLine {
					continue
				}
			}
			newLine := string(line[index+len(ts):]) + "\n"

			if strings.Contains(newLine, "Function") {
				doc += "## " + newLine
			} else if strings.Contains(newLine, "Example") {
				doc += "### " + newLine
			} else {
				doc += newLine
			}
		}
	}
	return
}
