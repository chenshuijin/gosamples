package main

/*
#include <stdio.h>
#include <stdlib.h>
#include "test.h"
*/
import "C"

import "fmt"

func main() {
	fmt.Println("start....")
	C.print(C.CString("test c print"))
	C.print1(C.CString("test c print1"))
	C.print2(C.CString("test c print2"))
	C.ttargbyte(C.CString("abcd"), C.CString("123456789ef"))
	fmt.Println("isboolok:", C.isboolok())
}
