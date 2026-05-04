# NoteShareEFREI

## necessary preparation
Add a file : db.env
```
MYSQL_ROOT_PASSWORD="password"
MYSQL_DATABASE="NoteShareEFREI"
MYSQL_USER="mysql"
MYSQL_PASSWORD="password_user"
```
Add a file : go-app.env
```
DB_USER="root"
DB_PASSWORD="password"
DB_NAME="NoteShareEFREI"
DB_HOST="db"
DB_PORT=3306
```


## Build and run the project
Currently for development
```sh
# inside the project root
go env -w GOEXPERIMENT=jsonv2
go build . && ./NoteShareEFREI
```
With the docker :
```sh
docker-compose up --build
```

# Windows
```powershell
go env -w GOEXPERIMENT=jsonv2
go build . ; .\NoteShareEFREI.exe
```
bash / command prompt :
```bash
go env -w GOEXPERIMENT=jsonv2
go build . && .\NoteShareEFREI
```
