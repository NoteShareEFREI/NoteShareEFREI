# NoteShareEFREI

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
