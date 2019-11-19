# gokreta

## Nyelvek/Languages
* [English](https://github.com/thegergo02/blob/master/gokreta/README.md)
* [Magyar](https://github.com/thegergo02/blob/master/gokreta/README.hu_HU.md)

## A csomag célja
Egy egyszerű absztrakció létrehozása a Kréta API-hoz.

## Példa
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

## Közremüködők
Név | Közremüködés
--- | ---
[thegergo02](https://github.com/thegergo02) | Maintainer
[boapps](https://github.com/boapps) | [API dokumentáció](https://github.com/boapps/e-kreta-api-docs)
