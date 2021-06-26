# Seakun Backend Engineer Test

### 1. 

```
If n=3 then make a code that produces a line like below

*
**
***
**
*
```

running

```
go to folder 1
go run main.go
```


### 2. 

```
If n=10 then make a code that generates a series of numbers like below

1,2,3,5,8,13,21,34,55

```

running

```
go to folder 2
go run main.go
```

### 3.

```
n=3 is 3 x 2 x 1 = 6. If n=5 then make a code that generates value like below

120
```

running

```
go to folder 3
go run main.go
```

### 4.

```
Create database and query with the case below

School system have 3 table (user, teacher, class). Each table must have
Primary Key (auto increment). class table have a relationship to teacher
with many to many relationship, each teacher have credential for the
system and the credential is held inside user table with one to one
relationship.
Create a query for Analyst to get data, with specification below:
- Teacher data with his/her username.
- Classes with the teacher name on each class.
```

running

```
run sql file in folder 4 to DB
```

### 5.

```
Create an API (REST) based on case no.4, with specifications below

- As an Administrator, can create teacher data.
- As an Administrator, can edit teacher data.
- As an Administrator, can delete teacher data.
- As an Administrator, can read teacher data.
- As an Administrator, can filter teacher data based on name and
birthdate

```

running

```
go to folder 5
go run main.go
```

I also prepared cUrl for each requirements

Create new Teacher's Data

```
curl --location --request POST 'http://localhost:8001/teacher/' \
--header 'Authorization: KillingOfaSacredDeer' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'teachername=Gonzales'
```

Get Teacher's Data

```
curl --location --request GET 'http://localhost:8001/teacher/' \
--header 'Authorization: KillingOfaSacredDeer'
```

Edit Teacher's Data

```
curl --location --request POST 'http://localhost:8001/updateteacher/' \
--header 'Authorization: KillingOfaSacredDeer' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'teachername=Buana Sentosa' \
--data-urlencode 'teacherid=3'
```

Delete Teacher's Data

```
curl --location --request DELETE 'http://localhost:8001/teacher/4' \
--header 'Authorization: KillingOfaSacredDeer'
```

Get Filtered Teacher's Data

```
curl --location --request GET 'http://localhost:8001/teacher/go' \
--header 'Authorization: KillingOfaSacredDeer'
```




DB using [Postgres](https://www.postgresql.org) database (internally port `5432`)

