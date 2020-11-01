package main

import (
	"fmt"

	"protoimport/assignment"
	"protoimport/errors"
)

func main() {
	e := &errors.Error{
		Code:    errors.Codes_CODE_PERMISSION_DENIED,
		Message: "",
	}
	fmt.Println(e)

	apikey := &assignment.Apikey{
		ApikeyId:      "id",
		Custodian:     "custodian",
		ApikeyPreview: "asdf",
	}
	fmt.Println(apikey)
}
