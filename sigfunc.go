package sigfunc

//License: MIT
//Copyright (c) 2015 Isagani Mendoza (itjumpstart.wordpress.com)

//Function signature for reading, writing and processing data (usually JSON)
//A read or write is a process (DAL, DML, etc)
//Business logic is also a process, etc.

type P00 interface {
	Process00()
}

type P01 interface {
	Process01() map[string]interface{}
}

type P10 interface {
	Process10(map[string]interface{})
}

type P11 interface {
	Process11(map[string]interface{}) map[string]interface{}
}
