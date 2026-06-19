# go-log-analyzer
This project is for learning the basics of the golang. Not a single line is copy pasted in this project, even this README

Copy the repository
```bash
git clone git@github.com:Griffinxd/go-log-analyzer.git
```


Build it 
```bash
go build -o loganalyzer src/main.go
```

Run it with empty arguments. It will inform you about the arguments
```bash
./loganalyzer
```
```bash
Usage of ./loganalyzer:
  -file string
        path for the target log file, required
  -format string
        Output format (text, json, csv). Default: text (default "text")
  -from string
        Start timestamp format: YYYY-MM-DD HH:MM:SS
  -level string
        INFO, WARNING, ERROR or empty for all
  -stats
        Print statistics instead of entries
  -to string
        End timestamp format: YYYY-MM-DD HH:MM:SS
```

Run it on tests
```bash
./loganalyzer -file=test/app_normal.log
```

```bash
./loganalyzer -file=test/app_large.log -stats
```



