package main

import "strconv"

type ExNumberFormationException struct{
	errorMessage string
	errorIdx int
}

func (e ExNumberFormationException) Error() string{
	return e.errorMessage + " " +  strconv.Itoa(e.errorIdx)
}
