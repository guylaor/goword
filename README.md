# goword
Go Package to extract text from word docx files


### usage

```
text, err := goword.ParseText("1.docx")
if err != nil {
    log.Panic(err)
}
fmt.Printf("%s ", text)
```




