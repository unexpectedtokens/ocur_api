package util

import (
	"fmt"
	"strconv"
)



func IDStringToINT(s string) (int, error){
	id, err := strconv.Atoi(s)
	if err != nil{
		return 0, fmt.Errorf("error converting id to int: %s", err.Error())
	}
	return id, nil
}