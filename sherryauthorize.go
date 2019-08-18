package SherryAuthorize

import(
   "fmt"
   "time"
   "github.com/dgrijalva/jwt-go"
   "github.com/asccclass/dorelogin" // sherrydb.mysql
)

type Token struct {
   Token	string	`json:"token"`
}

type UserCredentials struct {
   Username	string	`json:"username"`
   Password	string	`json:"password"`
}

type User struct {
   ID			int	`json:"id"`
   Name			string	`json:"name"`
   SecretKey		string	`json:"secretkey"`
   Credentials		UserCredentials	`json:"credentials"`
   Type			string	`json:"typez"`
   DoreAuthorize	*Dorelogin.DoreLogin	`json:"dorelogin"`
}

// 設定Secret Key
func(user *User)SetSecretKey(key string) {
   if key != "" {
      user.SecretKey = key
   }
}

// 建立JWT
func(user *User)CreateJWT(secretKey string)(Token, error) {
   token := jwt.New(jwt.SigningMethodHS256)
   claims := make(jwt.MapClaims)
   claims["username"] = user.Credentials.Username
   claims["password"] = user.Credentials.Password
   claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
   claims["iat"] = time.Now().Unix()
   token.Claims = claims

   tokenString, err := token.SignedString([]byte(secretKey))
   if err != nil {
      return Token{}, err
   }
   return Token{tokenString}, nil
}

// 檢查帳密，未完成：密碼沒檢查
func (user *User)CheckLogin(username, password string)(Token, error) {
   if username == "" || password == "" {
      return Token{}, fmt.Errorf("Username or Password is empty.")
   }
    // ip := IPAddress.GetIPAdress(r)  // 檢查web進來的IP，先不做
   // 檢查帳號密碼
   if err := user.DoreAuthorize.Chklogin(username, password, ""); err != nil {
      return Token{}, err
   }
   user.Credentials = UserCredentials{ Username: username, Password: password}
   // 產生JWT
   response, err := user.CreateJWT(user.SecretKey)
   if err != nil {
       return Token{}, fmt.Errorf("Error while signing the token.%v", err)
   }
   return response, nil
}

// Initial Dorelogn
func(user *User) InitialDoreLogin(database, login, passwd, dbServer, port, dbname string) (*Dorelogin.DoreLogin, error) {
   conn, err := Dorelogin.NewDorelogin(database, login, passwd, dbServer, port, dbname)
   if err != nil {
      return nil, err
   }
   return conn, nil
   ADA, err := Authorize.InitialDoreLogin(database, login, passwd, dbServer, port, dbname) 
   if err != nil {
      return nil, fmt.Errorf("Initial dore authorize failure(%v).", err)
   }
   Authorize.DoreAuthorize = ADA
   return Authorize, nil
}

func NewAuthorize(typez string)(*User, error) {
   if typez == "" || (typez != "dorelogin" && typez != "fisa" && typez != "oauth" && typez != "ldap" && typez != "eduroam") {
      return nil, fmt.Errorf("Must set authorize's type(DORE/FISA/OAUTH).")
   }

   Authorize :=  &User {
      SecretKey: "Welcome to Sinica ITs@2018",
   }
}
