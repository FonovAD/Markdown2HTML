# Benchmarks

### 1
#### Комментарий
Код до каких-либо изменений

#### Изменения в коде
-

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

### 2
#### Комментарий
Чтобы уменьшить количество аллокаций, создам массив определенного размера. На тесты это не повлияет, т.к. они не оценивают функции чтения.

#### Изменения в коде
.cmd/main.go#[30]

`FileWithoutByte0 := make([]byte, len(file))`

#### Тесты
    ```
    goos: darwin
    goarch: arm64
    pkg: Markdown_Processor/cmd
    BenchmarkMain-8                         Totat op: 100,   errors 0
            Totat op: 302,   errors 0
         302           3762007 ns/op         5626021 B/op      64346 allocs/op
            Totat op: 1,     errors 0
    BenchmarkHEADING-8                      Totat op: 100,   errors 0
            Totat op: 10000,         errors 41
            Totat op: 117484,        errors 481
      117484              9979 ns/op           16437 B/op        186 allocs/op
            Totat op: 1,     errors 0
    BenchmarkWORD-8                         Totat op: 100,   errors 0
            Totat op: 10000,         errors 0
            Totat op: 108308,        errors 0
      108308             10688 ns/op           16152 B/op        180 allocs/op
            Totat op: 1,     errors 0
    BenchmarkLIST-8                         Totat op: 100,   errors 0
            Totat op: 9514,          errors 0
        9514            122735 ns/op          186077 B/op       2138 allocs/op
            Totat op: 1,     errors 0
    BenchmarkNUMBEREDLIST-8                 Totat op: 100,   errors 0
            Totat op: 7696,          errors 0
        7696            145074 ns/op          219286 B/op       2529 allocs/op
            Totat op: 1,     errors 0
    BenchmarkBOLT-8                         Totat op: 100,   errors 0
            Totat op: 10000,         errors 0
            Totat op: 50112,         errors 3
       50112             24622 ns/op           36790 B/op        429 allocs/op
    PASS
    ok      Markdown_Processor/cmd  8.199s
    ```

### 3
#### Комментарий
Есть гипотеза, что длина строки и сложность разбора коррелируют. Поэтому решил разбить большой текст по строкам. По памяти это не даст прирост, но из-за меньшего количества проходов, необходимых для правильного лексического анализа, программа должна потреблять всех ресурсов меньше

#### Изменения в коде
.internal/processing/Lexer.go#[17]

`if count > (len(L.Code)*2)/3 {`

.cmd/bench_test.go#[65-73]

```
tokens := strings.Split(GeneralTest, "\n")
	for i := 0; i < b.N; i++ {
		for _, j := range tokens {
			TotalOperations += 1
			TT := TestedToken{TestTokenType: "main", TestToken: j}
			if err := TMain(TT); err != nil {
				errors += 1
			}
		}
```

#### Тесты
        ```
        goos: darwin
        goarch: arm64
        pkg: Markdown_Processor/cmd
        BenchmarkMain-8                         Totat op: 2000,          errors 2
                Totat op: 5360,          errors 2
                268           3857790 ns/op         4988687 B/op      56923 allocs/op
                Totat op: 1,     errors 1
        BenchmarkHEADING-8                      Totat op: 100,   errors 4
                Totat op: 10000,         errors 654
                Totat op: 124326,        errors 7527
                124326              9260 ns/op           14281 B/op        162 allocs/op
                Totat op: 1,     errors 0
        BenchmarkWORD-8                         Totat op: 100,   errors 0
                Totat op: 10000,         errors 0
                Totat op: 105986,        errors 0
                105986             11035 ns/op           16370 B/op        181 allocs/op
                Totat op: 1,     errors 0
        BenchmarkLIST-8                         Totat op: 100,   errors 0
                Totat op: 9554,          errors 0
                9554            120824 ns/op          178536 B/op       2042 allocs/op
                Totat op: 1,     errors 0
        BenchmarkNUMBEREDLIST-8                 Totat op: 100,   errors 0
                Totat op: 8535,          errors 0
                8535            134738 ns/op          202298 B/op       2313 allocs/op
                Totat op: 1,     errors 0
        BenchmarkBOLT-8                         Totat op: 100,   errors 0
                Totat op: 10000,         errors 0
                Totat op: 53166,         errors 2
                53166             22129 ns/op           34497 B/op        397 allocs/op
        PASS
        ok      Markdown_Processor/cmd  7.833s
        ```

### 4
#### Комментарий
Убрал ненужный компонент. Должно немного ускорить приложение.

#### Изменения в коде
.internal/processing/parser.go[7;100]
`var EmptyToken Token = Token{TokenType{"", ""}, ""}`

```listnode := Node{operator: Token{Type: SecondTokenTypes["GROUPNUMBEREDLIST"], Text: "GROUPNUMBEREDLIST"}}
```

.internal/processing/Token.go[3:6]
```
type Token struct {
	Type TokenType
	Text string
}
```


#### Тесты
        ```
        goos: darwin
        goarch: arm64
        pkg: Markdown_Processor/test
        BenchmarkMain-8                         Totat op: 2000,          errors 5
                Totat op: 8080,          errors 19
             404           2904436 ns/op         4600246 B/op      52454 allocs/op
                Totat op: 1,     errors 1
        BenchmarkHEADING-8                      Totat op: 100,   errors 5
                Totat op: 10000,         errors 633
                Totat op: 137887,        errors 8716
          137887              8460 ns/op           14716 B/op        165 allocs/op
                Totat op: 1,     errors 0
        BenchmarkWORD-8                         Totat op: 100,   errors 0
                Totat op: 10000,         errors 0
                Totat op: 118845,        errors 0
          118845             10234 ns/op           15753 B/op        178 allocs/op
                Totat op: 1,     errors 0
        BenchmarkLIST-8                         Totat op: 100,   errors 0
                Totat op: 10000,         errors 0
           10000            112036 ns/op          167888 B/op       1920 allocs/op
                Totat op: 1,     errors 0
        BenchmarkNUMBEREDLIST-8                 Totat op: 100,   errors 0
                Totat op: 9154,          errors 0
            9154            128578 ns/op          193625 B/op       2215 allocs/op
                Totat op: 1,     errors 0
        BenchmarkBOLT-8                         Totat op: 100,   errors 0
                Totat op: 10000,         errors 0
                Totat op: 56156,         errors 5
           56156             20666 ns/op           32434 B/op        370 allocs/op
        PASS
        ok      Markdown_Processor/test 7.855s
        ```


### 5
#### Комментарий
При конкатенации строк строка пересоздается. Работа с массивами выглядит иначе и должна быть эффективнее. Поэтому заменил все конкатенации строк на метод bytes.Buffer.WriteString

#### Изменения в коде
.internal/processing/dfs.go


#### Тесты
        ```
        goos: darwin
        goarch: arm64
        pkg: Markdown_Processor/test
        BenchmarkMain-8                         Totat op: 2000,          errors 6
                Totat op: 8060,          errors 32
             403           2891616 ns/op         4644179 B/op      52378 allocs/op
                Totat op: 1,     errors 0
        BenchmarkHEADING-8                      Totat op: 100,   errors 4
                Totat op: 10000,         errors 617
                Totat op: 138834,        errors 8603
          138834              8563 ns/op           15109 B/op        170 allocs/op
                Totat op: 1,     errors 0
        BenchmarkWORD-8                         Totat op: 100,   errors 0
                Totat op: 10000,         errors 0
                Totat op: 115190,        errors 0
          115190             10311 ns/op           16481 B/op        186 allocs/op
                Totat op: 1,     errors 0
        BenchmarkLIST-8                         Totat op: 100,   errors 0
                Totat op: 10000,         errors 0
           10000            106802 ns/op          170584 B/op       1947 allocs/op
                Totat op: 1,     errors 0
        BenchmarkNUMBEREDLIST-8                 Totat op: 100,   errors 0
                Totat op: 9810,          errors 0
            9810            117054 ns/op          189903 B/op       2147 allocs/op
                Totat op: 1,     errors 0
        BenchmarkBOLT-8                         Totat op: 100,   errors 0
                Totat op: 10000,         errors 1
                Totat op: 59487,         errors 10
           59487             19348 ns/op           31516 B/op        363 allocs/op
        PASS
        ok      Markdown_Processor/test 7.738s
        ```