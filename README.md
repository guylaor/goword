# goword   
![image](https://api.travis-ci.org/guylaor/goword.svg?branch=master)

Go Package to extract text from word docx files


### usage

```
import "github.com/guylaor/goword"

func main() {

    text, err := goword.ParseText("1.docx")
    if err != nil {
        log.Panic(err)
    }
    fmt.Printf("%s ", text)

}
```




