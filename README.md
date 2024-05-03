# Let's Build a Simple Database

[View rendered tutorial](https://cstack.github.io/db_tutorial/) (with more details on what this is.)

## Notes to myself

Run site locally:
```
bundle exec jekyll serve
```

## flowchart

```plantuml
@startuml
start
:input sql;
:parse sql;
switch(type)
    case(ddl)
        :execute ddl;
        stop
    case(dml)
        :execute dml;
        :prepare statement;
        :execute statement;
        :cursor open;
        :table open;
        :page open;
        :disk read;
        :output result;
endswitch
end
@enduml
```

## commands

```bash
# go test ./...
# cd gov && go test -v
```
