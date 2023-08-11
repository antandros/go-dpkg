package dpkg

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/antandros/go-pkgparser"
	"github.com/antandros/go-pkgparser/model"
)

func Parse(p *pkgparser.Parser) error {
	data, err := os.ReadFile("/var/lib/dpkg/status")
	if err != nil {
		return err
	}
	items := strings.Split(string(data), "\n\n")
	for _, item := range items {
		item = strings.Trim(item, " ")
		if len(item) == 0 {
			continue
		}
		r := regexp.MustCompile(`(?mi)^([\w-]+):\s*(.*)`)
		lines := strings.Split(item, "\n")
		var currentKey string
		var currenValue string
		packageItem := p.CreateModel()
		for _, line := range lines {

			match := r.FindStringSubmatch(line)
			if len(match) > 0 {
				if !strings.EqualFold(match[1], currentKey) {
					if currenValue != "" && currentKey != "" {
						packageItem, err = p.SetValue(currentKey, currenValue, packageItem)
						if err != nil {
							fmt.Println("Error", err.Error(), currenValue)
						}
					}
				}
				currentKey = match[1]
				currenValue = match[2]
			} else {
				currenValue += line
			}

		}
		p.Packages = append(p.Packages, packageItem)
	}

	return nil
}

func GetPackages() ([]model.Package, error) {
	var packages []model.Package
	p := new(pkgparser.Parser)
	p.Model = model.Package{}
	err := p.StructParse()
	if err != nil {
		return nil, err
	}
	err = Parse(p)
	if err != nil {
		return nil, err
	}
	for _, i := range p.Packages {
		item, ok := i.(*model.Package)
		if !ok {
			return nil, errors.New("struct conversion failed")
		}
		packages = append(packages, *item)
	}
	return packages, nil
}
