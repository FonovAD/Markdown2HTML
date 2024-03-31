### Правила оформления документа

Каждое изменение в сторону производительности должно сопровождаться следующим текстом:

#### Комментарий

#### Изменения в коде
.cmd/main.go#[54-57;62-94]     // [54-57;62-94] - строки файла

#### Тесты
    ```
    goos: darwin
    goarch: arm64
    pkg: Markdown_Processor/cmd
    BenchmarkMain-8                         Totat op: 100,   errors 0
            Totat op: 321,   errors 0
         321           3718841 ns/op         5576091 B/op      63467 allocs/op
            Totat op: 1,     errors 0
    BenchmarkHEADING-8                      Totat op: 100,   errors 0
            Totat op: 10000,         errors 41
            Totat op: 131205,        errors 529
      131205              9085 ns/op           15341 B/op        172 allocs/op
            Totat op: 1,     errors 0
    BenchmarkWORD-8                         Totat op: 100,   errors 0
            Totat op: 10000,         errors 0
            Totat op: 101533,        errors 0
      101533             11700 ns/op           17449 B/op        197 allocs/op
            Totat op: 1,     errors 0
    BenchmarkLIST-8                         Totat op: 100,   errors 0
            Totat op: 8817,          errors 0
        8817            122657 ns/op          185243 B/op       2131 allocs/op
            Totat op: 1,     errors 0
    BenchmarkNUMBEREDLIST-8                 Totat op: 100,   errors 0
            Totat op: 8215,          errors 0
        8215            140440 ns/op          214644 B/op       2455 allocs/op
            Totat op: 1,     errors 0
    BenchmarkBOLT-8                         Totat op: 100,   errors 0
            Totat op: 10000,         errors 0
            Totat op: 52081,         errors 0
       52081             22954 ns/op           35453 B/op        412 allocs/op
    PASS
    ok      Markdown_Processor/cmd  8.056s
    ```