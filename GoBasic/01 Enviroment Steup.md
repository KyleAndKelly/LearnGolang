
- Env: Win11 + WSL + VsCode

### Step 1. Download 

[Golang tar file for Linux](https://studygolang.com/dl)

### Step 2. Extract 

- the downloaded Go package to a specified directory by using the following command

```
cd /home

tar zxvf go1.18.3.linux-amd64.tar.gz

```
### Step3. Modify /etc/profile

Use the following command to apply the changes to the current shell session

```
export GOROOT=/home/go
export GOPATH=/home/gopath
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
export GOPROXY="https://goproxy.io"
```

### Step4. Source  /etc/profile

```
source /etc/profile

```

### Step5. Test 

Type the following command in the terminal

```
root@DESKTOP-AHPD166:~# go version
```

If the following message is displayed, it means that the Golang environment has been set up correctly on your Linux WSL.

```
go version go1.21.3 linux/amd64
```

### Step6. Build hello world.go

- Create a new file named "hello_world.go" in any directory and add the following contents:

```
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}

```
- Then build the program by using the following commands:

```
go build hello_world.go

```
- Normally, an executable binary named "hello_world" should be created in the current directory, then run it

```
./hello_world

```

Everything is Ready, go and enjoy it 
:)
