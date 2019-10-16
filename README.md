# gokreta

## Nyelvek/Languages
* [Magyar](https://github.com/thegergo02/gokreta/README.md)
* [English](https://github.com/thegergo02/gokreta/README.en_US.md)

## A csomag célja
Egy egyszerű absztrakció létrehozása a Kréta API-hoz.

##Példa
```
package main

import (
	"fmt"
	"github.com/thegergo02/gokreta"
)

func main() {
	instituteCode := "<instituteCode>"
	userName := "<userName>"
	password := "<password>"
	authDetails, err := gokreta.GetAuthDetailsByCredetinals(instituteCode,
		userName,
		password,
	)
	student, err := gokreta.GetStudentDetails(instituteCode, authDetails.AccessToken)
	fmt.Println(student.Name)
	if err != nil {
		fmt.Println("error")
	}
}
```

## Közremüködők
Név | Közremüködés
--- | ---
[thegergo02](https://github.com/thegergo02) | Maintainer
[boapps](https://github.com/boapps) | [API dokumentáció](https://github.com/boapps/e-kreta-api-docs)
