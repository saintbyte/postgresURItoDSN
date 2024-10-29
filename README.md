# postgresURItoDSN
GORM отличился тем что использует DSN ( Data Source Name ). А сейчас модно и молодежно использовать database url  ( даже консольный psql поддерживает ) 
и это модуль содержит одну фукнцию UriToDSN с одним параметром. Задача этой функции из строки ```"postgresql://user:password@localhost:5432/dbname?param1=value1"```,
получить строку ```"user=user password=password host=localhost port=5432 dbname=dbname param1=value1"``` которую уже потом можно использовать в качестве параметра для ```postgres.Open``` 
в примерно таком кода ```gorm.Open(postgres.Open(dsn), &gorm.Config{})``` . Пока только postgres:// ( C остальным работает судя по документации и так нормально )

## install
```
go get github.com/saintbyte/postgresURItoDSN
```

Буду рад вашим пулл реквестам. 
