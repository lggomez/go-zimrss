package zim

import (
	"strings"
	"time"
)

type PageMetadata struct {
	Title        string
	ContentType  string
	WikiFormat   string
	CreationDate time.Time
	ContentDate  *time.Time
	Path         string
}

/* PathToURL translates a page path into a valid HTML URL. This makes a couple of assumptions:
* The link provided to the program is the root domain of the generated site (i.e.: https://www.foo.com)
* The file structure of the local notebook is the same as the one of the website
 */
func (p *PageMetadata) PathToURL(rootPath string, baseURL string, fileExtension string) string {
	urlPath := strings.ReplaceAll(p.Path, rootPath, "")

	urlPrefix := baseURL
	if !strings.HasSuffix(baseURL, "/") {
		urlPrefix = urlPrefix + "/"
	}

	urlSuffix := urlPath
	if strings.HasPrefix(urlSuffix, "/") && len(urlSuffix) > 0 {
		urlSuffix = urlSuffix[1:]
	}

	urlSuffix = strings.ReplaceAll(urlSuffix, fileExtension, ".html")

	return urlPrefix + urlSuffix
}

type PageMetadataByCreationDate []*PageMetadata

func (p PageMetadataByCreationDate) Len() int {
	return len(p)
}
func (p PageMetadataByCreationDate) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PageMetadataByCreationDate) Less(i, j int) bool {
	return p[i].CreationDate.Before(p[j].CreationDate)
}
