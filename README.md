# GetBlogger.v1

Clone post blog di [Blogger](https://blogger.com) , bisa di blog pribadi maupun di blog orang lain :warning:  .


### Opsi yang di butuhkan
| Flag      | Kegunaan | type |
|-----------|----------|------|
| `-url`    | memanggil url | `string` |
| `-result` | menentukan result dari post blogger (default 500) | `int` |
| `-out`    | membuat output dari result blogger, type output menggunakn `.xml` | `string` |




### Penggunaan
1. Melihat opsi bantuan penggunaan
    ```
    gopher@localhost:~# go run main.go --help
    Usage of /tmp/go-build778205588/b001/exe/main:
    -out string
        	Output File !!!, format output is .xml
    -result int
        	Result post from link blogger (default 500)
    -url string
        	URL for target Blogger !!!
    exit status 2
    gopher@localhost:~#
    ```
2. Contoh Penggunaan
    ```
    gopher@localhost:~# go run main.go -url https://example.blogspot.com -result 500 -out example-post.xml
    ```
3. Menjadikan file `.go` menjadi executable 
    ```
    gopher@localhost:~# go build -o getblogger main.go

    ``` 

### Import

1. Buka [Blogger](https://blogger.com) 
2. Login dengan akun anda
3. Pilih pada `Settings >> Other >> Import Content`

### Author

    Name : Erwin Kurniawan
    Mail : rootshaor@gmail.com
    Web  : https://rootshaxor.github.io
    Fb   : https://facebook.com/rootshaxor