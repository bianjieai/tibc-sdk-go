package commitment

import "net/url"

func (mp MerklePath) String() string {
	pathStr := ""
	for _, k := range mp.KeyPath {
		pathStr += "/" + url.PathEscape(k)
	}
	return pathStr
}
