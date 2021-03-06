package main

import (
	"crypto/sha1"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Asset struct {
	ID   string   // unique id, sha1 of relative path
	Path []string // relative path, split
	File string   // full path
	Date int64    // mtime, unix epoch
}

var (
	tmplFuncs = template.FuncMap{
		"literal": func(in interface{}) string {
			return fmt.Sprintf("%#v", in)
		},
		"load": func(filename string) ([]byte, error) {
			return ioutil.ReadFile(filename)
		},
	}

	tmplStr = `// generated by binfs

package {{.Package}}

import (
  "time"
  "go.guoyk.net/binfs"
)

var (
{{range .Assets}}
  binfs{{.ID}} = binfs.Chunk{
    Path: {{.Path | literal}},
    Date: time.Unix({{.Date}}, 0),
    Data: {{.File | load | literal }},
  }
{{end}}
)

func init() {
{{range .Assets}}
  binfs.Load(&binfs{{.ID}})
{{end}}
}
`
	tmpl = template.Must(template.New("binfs").Funcs(tmplFuncs).Parse(tmplStr))
)

var (
	optPackage string
)

func scan(dirs []string) (assets []Asset, err error) {
	for _, dir := range dirs {
		if strings.HasSuffix(dir, "/") {
			dir = dir[0 : len(dir)-1]
		}
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			_, file := filepath.Split(path)
			// skip hidden directories
			if info.IsDir() {
				if strings.HasPrefix(file, ".") {
					return filepath.SkipDir
				} else {
					return nil
				}
			}
			// skip hidden files
			if strings.HasPrefix(file, ".") {
				return nil
			}
			// relative path
			var rel string
			if rel, err = filepath.Rel(dir, path); err != nil {
				return err
			}
			comps := append([]string{dir}, strings.Split(rel, string(filepath.Separator))...)
			assets = append(assets, Asset{
				ID:   fmt.Sprintf("%02x", sha1.Sum([]byte(rel))),
				Date: info.ModTime().Unix(),
				File: path,
				Path: comps,
			})
			return nil
		})
	}
	return
}

func exit(err *error) {
	if *err != nil {
		_, _ = fmt.Fprintln(os.Stderr, (*err).Error())
		os.Exit(1)
	}
}

func main() {
	var err error
	defer exit(&err)

	flag.StringVar(&optPackage, "pkg", "main", "package name for generated go file")
	flag.Parse()

	dirs := flag.Args()
	if len(dirs) == 0 {
		err = errors.New("missing directories")
		return
	}

	var assets []Asset
	if assets, err = scan(dirs); err != nil {
		return
	}

	err = tmpl.Execute(os.Stdout, map[string]interface{}{
		"Package": optPackage,
		"Assets":  assets,
	})
}
