# gokreta

## Languages/Nyelvek
* [Magyar](https://github.com/thegergo02/blob/master/gokreta/README.hu_HU.md)
* [English](https://github.com/thegergo02/blob/master/gokreta/README.md)

## The package's goal
A simple abstraction to the KRÉTA API.

## Example
```golang
package main

import (
	"fmt"
	"github.com/thegergo02/gokreta"
)

func main() {
	instituteCode := "<instituteCode>"
	userName := "<userName>"
	password := "<password>"
	user, err := gokreta.NewUser(instituteCode,
		userName,
		password,
	)
	student, err := user.GetStudentDetails()
	if err != nil {
		panic(err)
	}
	fmt.Println(student.Name)
}
```

## Supporters
Name | Act
--- | ---
[thegergo02](https://github.com/thegergo02) | Maintainer
[boapps](https://github.com/boapps) | [API documentation](https://github.com/boapps/e-kreta-api-docs)
