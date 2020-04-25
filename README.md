# Berlin Parking Extractor

## Frameworks & libraries
* [Gin Web Framework](https://github.com/gin-gonic/gin)
* [Colly](https://github.com/gocolly/colly)

## Extractor
Extractor is a program, which extracts parking data from https://www.hipark.de/parkplatzbelegung/.

### Instructions  
1. Execute ```$ go run main.go```  in extractor folder.
 

## API server
A backend API, which has a POST endpoint that returns a JSON response.

### Instructions
2. Execute ```$ go run main.go```, which will start the server at port ```8080```.

## Example
### JSON file
```
[
   {
      "name":"Alt. Stobenstr.",
      "status":"offen",
      "free_spaces":0,
      "time":1574673961000
   },
   {
      "name":"Andreaspassage",
      "status":"offen",
      "free_spaces":103,
      "time":1574673961000
   },
   ...
]
```
### JSON Fail Response Body
```
[
    {
        "status": "error",
        "data": "Parking not found",
        "message": "Error: unable to get parking!"
    }
]
```
### JSON Request Body
```
{
  "name":"Andreaspassage",
},
```
### JSON Successful Response Body
```
{
  "name":"Andreaspassage",
  "status":"offen",
  "free_spaces":103,
  "time":1574673961000
}
```