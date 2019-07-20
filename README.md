## Sherry Authorize tool

### Installation
```
go get github.com/asccclass/sherryauthorize
```

### Usage
```
import(
   "os"
   "fmt"
   "github.com/asccclass/sherryauthorize"
)

func main() {
   test, err := InitialAuthorize(os.Getenv("DBMS"), os.Getenv("DBLOGIN"), os.Getenv("DBPASSWORD"), os.Getenv("DBSERVER"), os.Getenv("DBPORT"),os.Getenv("DBNAME"))

   if err != nil {
      fmt.Printf("Initial Error: %v", err)
      return
   }

   token, err := test.chkLoginFromJSON(username, password)
   if err != nil {
      fmt.Printf("%v", err)
      return
   }
   fmt.Printf("%v\n", token)
}
```
