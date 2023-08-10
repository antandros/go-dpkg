
# dpkg Go Package

## Overview

The `dpkg` package is a Go library designed to list software installed on Debian distributions using `dpkg`. The package provides functionality to parse package details and represent them in structured Go data types.

## Data Structures

### Package

Contains detailed information about a Debian package. Key fields include:

- `PackageName`: Name of the package.
- `Status`, `Priority`, `Section`, and more: Various details about the package.
- `Maintainer`, `OriginalMaintainer`: Contact information of the package maintainers.
- `Extra`: A map that can hold any additional information about the package not covered by the other fields.

### PackageContact

Represents the contact information of a package. It can be either an email address or a website.

- `Name`: The name of the contact.
- `Contact`: The actual contact information - email or website.
- `Type`: Specifies if the contact is an "email" or "website".

## Usage

```
package main

import (
	"encoding/json"
	"fmt"

	"gitlab.com/antandros/go-dpkg"
)

func main() {
	packages, err := dpkg.GetPackages()
	fmt.Println(err)
	resp, err := json.MarshalIndent(packages, "", "\t")
	fmt.Println(err)
	fmt.Println(string(resp))
}
```

## License

Refer to the `LICENSE` file in the repository.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue on the project's repository.