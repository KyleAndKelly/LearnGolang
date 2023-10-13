### Step 1. Download and Install 

```
cd $your_worspace_dir

go mod init 

go get -u github.com/gin-gonic/gin

```

### Step 2.Import  

-  example.go

```
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.Run()
}

```

### Step 3. Build and Run

- Build and run

```
go run example.go
```

- Then open http://127.0.0.1:8080/ping, you will get message like this

```
{"message":"hello world"}

```