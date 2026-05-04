# NoteShareEFREI
http://famogal.freeboxos.fr:18000/

## Necessary preparation
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


## Installation
Make sure you have Docker and Docker Compose installed on your machine. Then, clone the repository and navigate to the project directory:
```sh
git clone https://github.com/NoteShareEFREI/NoteShareEFREI.git
cd NoteShareEFREI
```

## Build and run the project
With the docker :
```sh
# From the project root directory
docker-compose up --build
sudo mysql -u ${DB_user} -p ${MYSQL_ROOT_PASSWORD} -P ${DB_PORT} --skip-ssl < sql/schema.sql
sudo mysql -u ${DB_user} -p ${MYSQL_ROOT_PASSWORD} -P ${DB_PORT} --skip-ssl < sql/view.sql 
sudo mysql -u ${DB_user} -p ${MYSQL_ROOT_PASSWORD} -P ${DB_PORT} --skip-ssl < sql/migration_add_indexes.sql 
```
Then you'll be able to access the application on http://localhost:8080/ 

## Team
- Alexis DELAVIS : Lead developer
- Guillaume BERNARD : Database Engineer
- Jude AYBALEN : UML engineer 
- Thomas BIETH-LEGOUT : Lead graphist
- Alexis LAFARGUE : Security engineer

## Technical stack
| Technology          | Usage                                         | GreenIt Justification                                                                                                                                         |
|---------------------|-----------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Golang              | Backend development                           | Compiled language so much more efficient for Backend development even under heavy load                                                                        |
| MySQL               | Database management                           | Didn't need more than MySQL for how little our database is. We could even have gone for sqlite, but it wouldn't have changed a lot the ecological performance |
| Docker              | Containerization                              | No greenIT justification just for easier deployment                                                                                                           |
| Docker Compose      | Orchestration of multi-container applications | Same as Docker                                                                                                                                                |
| HTML/CSS/JavaScript | Frontend development                          | Minimal javascript and minified css but there was not really a better choice aside from WebGL and WASM which is not adapted for anything which isn't a game.  |

## Project Structure

```
NoteShareEFREI/
├── api/                                    # API handlers for core features
│   ├── account.go                         # Account management endpoints
│   ├── admin.go                           # Admin-related endpoints
│   ├── comments.go                        # Comments feature endpoints
│   ├── create.go                          # Note/sheet creation endpoints
│   └── login.go                           # Login endpoints
├── backend/                                # Backend utilities and middleware
│   ├── cookie.go                          # Cookie management
│   ├── hash.go                            # Password hashing utilities
│   └── middleware.go                      # HTTP middleware
├── database/                               # Database operations
│   ├── addvalues.go                       # Insert/Update operations
│   ├── database.go                        # Database connection and init
│   └── scan.go                            # Result scanning utilities
├── files/                                  # Static files and course materials
│   ├── styles.css                         # Stylesheet
│   └── {Category}/                        # Category folders (e.g., Math, Physics)
│       └── {SubCategory}/                 # Subcategory folders (e.g., Algebra, Mechanics)
├── js/                                     # JavaScript files
│   └── script.js                          # Frontend scripts
├── routers/                                # HTTP route handlers
│   ├── account.go                         # Account routes
│   ├── admin.go                           # Admin routes
│   ├── create.go                          # Signup routes
│   ├── createsheet.go                     # Sheet creation routes
│   ├── default.go                         # Default routes
│   ├── home.go                            # Home page routes
│   ├── login.go                           # Login routes
│   └── template_cache.go                  # Template caching
├── sql/                                    # Database SQL files
│   ├── schema.sql                         # Initial database schema
│   ├── view.sql                           # Database views
│   ├── initial_data.sql                   # Optional Initial data
│   └── migration_add_indexes.sql          # Database indexes
├── templates/                              # HTML templates
│   ├── account                            # Account management template
│   ├── account_redirect                   # Account redirect template
│   ├── admin                              # Admin panel template
│   ├── create_account                     # Account creation template
│   ├── createsheetfailed                  # Failed sheet creation template
│   ├── createsheetsuccess                 # Successful sheet creation template
│   ├── home                               # Home page template
│   └── log_in                             # Login page template
├── main.go                                # Application entry point
├── go.mod                                 # Go modules definition
├── go.sum                                 # Go modules checksums
├── docker-compose.yml                     # Docker Compose configuration
├── Dockerfile                             # Docker image configuration
├── db.env                                 # Database environment variables
├── go-app.env                             # Go app environment variables
├── README.md                              # Project documentation
└── NoteShareEFREI                         # Compiled binary
```

