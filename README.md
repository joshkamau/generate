# generate
This package generates text files given a template and a json file containing the data. 

Usage of 
`./generate 
  -json string
        json file path
  -np string
        output file name property
  -out string
        output dir
  -tpl string
        template file path
 `       
The json file must contain an array. Every element in the array will generate a file. 
The -np flag specify with template in the json object will contain the file name. This property can be 
a golang text template and can contain functions. e.g  {"name":"Jane", "file":"{{.name | ToUpper }}.html"},
