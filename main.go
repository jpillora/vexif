package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"sort"

	"github.com/jpillora/opts"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

var c = struct {
	Files []string `opts:"mode=arg, help=file to view exif info"`
}{}

func main() {
	//parse cli
	opts.New(&c).Parse()
	//view file info
	for _, fp := range c.Files {
		if err := view(fp); err != nil {
			log.Fatalf("%s", err)
		}
	}
	fmt.Print(">>> done\n")
}

func view(fp string) error {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()
	e, err := exif.Decode(f)
	if err != nil {
		return fmt.Errorf("exif decode failed: %w", err)
	}
	t := tags{}
	e.Walk(t)
	fmt.Printf(">>> %s\n", fp)
	keys := []string{}
	for k := range t {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("  %s = %s\n", k, t[k])
	}
	return nil
}

type tags map[string]string

func (t tags) Walk(name exif.FieldName, tag *tiff.Tag) error {
	t[string(name)] = tag.String()
	return nil
}
