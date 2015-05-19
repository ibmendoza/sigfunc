package sigfunc

//Author: Gani Mendoza (itjumpstart.wordpress.com)

//Function signature for reading, writing and processing data (usually JSON)
//A read or write is a process (DAL, DML, etc)
//Business logic is also a process, etc.

type MapSI map[string]interface{}

type P00 interface {
	Process00()
}

type P01 interface {
	Process01() MapSI
}

type P10 interface {
	Process10(MapSI)
}

type P11 interface {
	Process11(MapSI) MapSI
}
