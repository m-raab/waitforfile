Wait For File
=========================
This tool checks the availability of a file with a special content string.

Usage
-------------------------

It is possible to configure the following parameters:

```
Usage of waitforfile:
  -file filepath
        The path for the file.
  -version string
        The content of the file.
  -timeout int
    	Timeout for waiting in seconds
  -timeperiod int
    	Time between checks in seconds  
```
    	
The program will wait `timeperiod` seconds between each check. 'waitforfile' will finish with an exit code after the `timeout`, if no file is available.

Exit Codes
-------------------------
| Code | Description |
|----:|------------------------|
| 20 | Content is not correct |
| 10 | No file available |
|  0 | File with version found |
| 101 | File path for file is not specified |
| 106 | Parameter timeperiod is bigger or equal to timeout. The timeout configuration must be bigger than the timeperiod. |
| 107 | Parameter timeout must be bigger than 0. |
| 108 | Parameter timeperiod must be bigger than 0. |