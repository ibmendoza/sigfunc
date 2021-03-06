## sigfunc

Go function signature for reading, writing and processing arbitrary data

### Rationale

When writing business apps, much of reading, writing and processing functions 
at the business logic tier have function parameters that vary from zero to any 
arbitrary number. I'd like to put a stop to this function signature zoo, and so 
I have created a generic building block for function signature.

Since JSON can represent any business data structure, I opted for 
`map[string]interface{}` type so your function can easily consume or generate 
JSON data if desired.

### Type MapSI

A shorthand convenience type for `map[string]interface{}`


### Function signature

Function signature          |        Result              
----------------------------|-------------------------
func Process00()            | none                      
func Process01() MapSI      | map[string]interface{}  
func Process10(MapSI)       | none                    
func Process11(MapSI) MapSI | map[string]interface{}   

### Why 00, 01, 10 and 11 suffixes?

Elementary, my dear Watson.

Suffix |  With parameter? | With result?
-------|------------------|-------------
00     | No               | No 
01     | No               | Yes
10     | Yes              | No
11     | Yes              | Yes


### Limitation

This package is limited to business data that can be represented as JSON, so 
keep in mind that it was designed for business applications.

### License

MIT

Copyright (c) 2015 Isagani Mendoza
http://itjumpstart.wordpress.com
