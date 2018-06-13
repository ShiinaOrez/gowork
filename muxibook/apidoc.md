-------
## URL: /signup
## Methods: POST

|Headers| |
|-------|-|
|None| |

|URL Params|
|----------|
|None|

POST Data:

```
{
	"username": string
}
```

RETURN Data(200 OK):

```
{
	"msg": "successful"
}
```

|Status Code| |
|-----------|-|
|200|OK|
|401|Username has been registered|

-------




-------
## URL: /login
## Methods: POST

|Headers| |
|-------|-|
|None| |

|URL Params|
|----------|
|None|

POST Data:

```
{
	"username": string
}
```

RETURN Data(200 OK):

```
{
	"token": user.ID+user.Username
}
```

|Status Code| |
|-----------|-|
|200|OK|

-------



-------
## URL: /addbook
## Methods: POST

|Headers| |
|-------|-|
|token|string|

|URL Params|
|----------|
|None|

POST Data:

```
{
	"kind": int     //the kind of this book
	"book": string  //book's name
	"no": string    //book's number
}
```

RETURN Data(200 OK):

```
{
	"msg": "add successful!"
}
```

|Status Code| |
|-----------|-|
|200|OK|
|400|Bad Request|
|401|Confirm Failed|
|402|Book Re-Add|

-------



-------
## URL: /lendbook
## Methods: POST

|Headers| |
|-------|-|
|token|string|

|URL Params|
|----------|
|None|

POST Data:

```
{
	"book": string      //book's name
	"username": string  //username
	"realname": string  //user's realname
}
```

RETURN Data(200 OK):

```
{
	"book":      string     //book's name
	"kind":      int        //book's kind
	"available": int        //book is available or not
	"who":       string     //who lend this book
	"when":      time.Time  //when lend this book
	"realname":  string     //user's realname
}
```

|Status Code| |
|-----------|-|
|200|OK|
|400|Bad Request|
|401|Confirm Failed|
|402|User/Book Not Found|
|403|User's Book Limited|
|407|Book Not Available|

-------



-------
## URL: /returnbook
## Methods: POST

|Headers| |
|-------|-|
|None| |

|URL Params|
|----------|
|None|

POST Data:

```
{
	"no": string       //book's number
	"username": string //username
}
```

RETURN Data(200 OK):

```
{
	"msg": "successful"
}
```

|Status Code| |
|-----------|-|
|200|OK|
|400|Bad Request|
|402|User/Book Not Found|
|403|Book Still Available|
|407|Book Not Belong To You|

-------



-------
## URL: /renewbook
## Methods: POST

|Headers| |
|-------|-|
|None| |

|URL Params|
|----------|
|None|

POST Data:

```
{
	"no": string       //book's number
	"username": string //username
}
```

RETURN Data(200 OK):

```
{
	"msg": "successful!"
}
```

|Status Code| |
|-----------|-|
|200|OK|
|400|Bad Request|
|401|Beyond Time Limit|
|402|User/Book Not Found|

-------



-------
## URL: /mybook
## Methods: POST

|Headers| |
|-------|-|
|token|string|

|URL Params|
|----------|
|None|

POST Data:

```
{
	"username": string
}
```

RETURN Data(200 OK):

```
{
	lend:[
		{
			"no":         string    //book's number
			"bookname":   string    //book's name
			"returntime": time.Time //book's return time
		}
		...
	]
}
```

|Status Code| |
|-----------|-|
|200|OK|
|401|Confirm Failed|
|402|User Not Found|

-------



-------
## URL: /searchbook
## Methods: POST

|Headers| |
|-------|-|
|None| |

|URL Params|
|----------|
|None|

POST Data:

```
{
	"partten": string  //partten string
	"page":    int     //page number
}
```

RETURN Data(200 OK):

```
{
	"num": int    //how many match record
	"page": int   //page number
	"books":[
		{
			"book":      string    //book's name
			"kind":      int       //book's kind
			"available": int       //book is available or not
			"who":       string    //who lend it
			"when":      time.Time //when lend it
			"realname":  string    //user's realname
		}
	]
}
```

|Status Code| |
|-----------|-|
|200|OK|
|401|No Match Record|

-------



-------
## URL: /getbook
## Methods: GET

|Headers| |
|-------|-|
|None| |

|URL Params|
|----------|
|kind|
|page|

POST Data:

```
{
	"username": string
}
```

RETURN Data(200 OK):

```
{
	"num": int    //how many match record
	"page": int   //page number
	"books":[
		{
			"book":      string    //book's name
			"kind":      int       //book's kind
			"available": int       //book is available or not
			"who":       string    //who lend it
			"when":      time.Time //when lend it
			"realname":  string    //user's realname
		}
	]
}
```

|Status Code| |
|-----------|-|
|200|OK|
|401|No Match Record|

-------