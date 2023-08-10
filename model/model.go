package model

type PackageContact struct {
	Name    string `json:"name,omitempty"`
	Contact string `json:"contact,omitempty"`
	Type    string `json:"type,omitempty"`
}
type Package struct {
	Extra              map[string]interface{} `parser:"*,others" json:"extra,omitempty"`
	PackageName        string                 `parser:"Package" json:"packageName,omitempty"`
	Important          string                 `parser:"Important" json:"important,omitempty"`
	Protected          string                 `parser:"Protected" json:"protected,omitempty"`
	Status             string                 `parser:"Status" json:"status,omitempty"`
	Priority           string                 `parser:"Priority" json:"priority,omitempty"`
	Section            string                 `parser:"Section" json:"section,omitempty"`
	Installed          string                 `parser:"Installed" json:"installed,omitempty"`
	VisiblePkgName     string                 `parser:"Cnf-Visible-Pkgname" json:"visiblePkgName,omitempty"`
	Maintainer         PackageContact         `parser:"Maintainer" json:"maintainer,omitempty"`
	Architecture       string                 `parser:"Architecture" json:"architecture,omitempty"`
	MultiArch          string                 `parser:"Multi-Arch" json:"multiArch,omitempty"`
	InstalledSize      int                    `parser:"Installed-Size" json:"installedSize,omitempty"`
	Source             string                 `parser:"Source" json:"source,omitempty"`
	Version            string                 `parser:"Version" json:"version,omitempty"`
	Description        string                 `parser:"Description" json:"description,omitempty"`
	Vendor             string                 `parser:"Vendor" json:"vendor,omitempty"`
	Depends            string                 `parser:"Depends" json:"depends,omitempty"`
	License            string                 `parser:"License" json:"license,omitempty"`
	OriginalMaintainer PackageContact         `parser:"Original-Maintainer" json:"originalMaintainer,omitempty"`
	Homepage           string                 `parser:"Homepage" json:"homepage,omitempty"`
}
