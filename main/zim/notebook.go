package zim

import "strconv"

type Notebook struct {
	Version              string
	Name                 string
	InterWiki            string
	Home                 string
	Icon                 string
	DocumentRoot         string
	Shared               bool
	EOL                  string
	DisableTrash         bool
	ShortRelativeLinks   bool
	DefaultFileFormat    string
	DefaultFileExtension string
	NotebookLayout       string
}

func NotebookFromMap(notebookMap map[string]string) Notebook {
	mapGetter := func(m map[string]string, key string) string {
		if val, ok := m[key]; ok {
			return val
		}
		return ""
	}

	shared, _ := strconv.ParseBool(mapGetter(notebookMap, "shared"))
	disableTrash, _ := strconv.ParseBool(mapGetter(notebookMap, "disable_trash"))
	shortLinks, _ := strconv.ParseBool(mapGetter(notebookMap, "short_relative_links"))

	n := Notebook{
		Version:              mapGetter(notebookMap, "version"),
		Name:                 mapGetter(notebookMap, "name"),
		InterWiki:            mapGetter(notebookMap, "interwiki"),
		Home:                 mapGetter(notebookMap, "home"),
		Icon:                 mapGetter(notebookMap, "icon"),
		DocumentRoot:         mapGetter(notebookMap, "document_root"),
		Shared:               shared,
		EOL:                  mapGetter(notebookMap, "endofline"),
		DisableTrash:         disableTrash,
		ShortRelativeLinks:   shortLinks,
		DefaultFileFormat:    mapGetter(notebookMap, "default_file_format"),
		DefaultFileExtension: mapGetter(notebookMap, "default_file_extension"),
		NotebookLayout:       mapGetter(notebookMap, "notebook_layout"),
	}

	return n
}
