package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	database "server/Database"
	"server/models"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// signup route
func Signup(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user models.Users
	json.NewDecoder(r.Body).Decode(&user)
	var result	models.Users
	err := database.UserCollection.FindOne(context.TODO(),bson.M{"email":user.Email}).Decode(&result)
	if err != nil{
					BcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
			user.Password = string(BcryptPassword)
			cur ,err := database.UserCollection.InsertOne(context.TODO(),user)
			if err	!= nil{
				log.Fatal(err)
			}
			json.NewEncoder(w).Encode(cur.InsertedID)
			} else {
		json.NewEncoder(w).Encode("User already exists with this email")
	}}

	func enableCors(w *http.ResponseWriter) {
(*w).Header().Set("Access-Control-Allow-Origin", "*")
(*w).Header().Set("Access-Control-Allow-Headers", "authentication, Content-Type")
}

// create JWT

	var jwtSecretKey = []byte("secret-key")
	type Claims struct {
    Email string
				jwt.StandardClaims
}

func CreateJWT(email string) (response string, err error) {
    expirationTime := time.Now().Add(1 * time.Hour)
    claims := &Claims{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecretKey)
    if err == nil {
        return tokenString, nil
    }
    return "", err
}

// VerifyToken func will used to Verify the JWT Token 


func VerifyToken(tokenString string) (email string, err error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecretKey, nil
    })

    if token != nil {
        return claims.Email, nil
    }
    return "", err
}

	// login route
	func Login(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user models.Users
	json.NewDecoder(r.Body).Decode(&user)
	var	result models.Users
	err := database.UserCollection.FindOne(context.TODO(),bson.M{"email":user.Email}).Decode(&result)
	if err  != nil{
		json.NewEncoder(w).Encode("User not found")
	}else{
	err = bcrypt.CompareHashAndPassword([]byte(result.Password),[]byte(user.Password))
	if err != nil{
		json.NewEncoder(w).Encode("Wrong password")
	}else{
		token,_:=CreateJWT(user.Email)
	type response struct{
		Token	string
		Name string
	}
	res :=	response{token,result.Name}
	json.NewEncoder(w).Encode(res)
	}
}
}

// Forget password

type forgetuser struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

func Forgetpassword(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	var user forgetuser
	json.NewDecoder(r.Body).Decode(&user)
	var olduser models.Users
	err := database.UserCollection.FindOne(context.TODO(),bson.D{{"email",user.Email}}).Decode(&olduser)
	if err	!= nil{
		json.NewEncoder(w).Encode("User not found")
	}else{
		BcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	json.NewEncoder(w).Encode(olduser)
	filter :=  bson.D{{"email",user.Email}}
		update :=  bson.M{"$set":bson.D{{"password",BcryptPassword}}}
	_,err := database.UserCollection.UpdateOne(context.TODO(),filter,update)
	if err != nil{
		json.NewEncoder(w).Encode("Error updating password")
	}
	}
}
  
// getallproducts route

func Getproducts(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	token := r.Header.Get("Authorization")
	tokenString := strings.Split(token," ")[1]
	email,_	:= VerifyToken(tokenString)
	w.Header().Set("Content-Type", "application/json")
	var products []primitive.M
	cur , err := database.ProductCollection.Find(context.TODO(),bson.D{{"user",email}})
	if err != nil {
		json.NewEncoder(w).Encode("No products found")
	}

	for cur.Next(context.TODO()){
		var prod primitive.M
		err := cur.Decode(&prod)
		if err!= nil{
			log.Fatal(err)
		}
		products = append(products, prod)
	}
	cur.Close(context.TODO())
	json.NewEncoder(w).Encode(products)
}

// addProduct route

func Addproduct(w http.ResponseWriter , r *http.Request){
enableCors(&w)
token :=	r.Header.Get("Authorization")
tokenString := strings.Split(token," ")[1]
email,_	:= VerifyToken(tokenString)
w.Header().Set("Content-Type", "application/json")
var prod models.Product
prod.User=email
json.NewDecoder(r.Body).Decode(&prod)
fmt.Print(prod)
_ , err := database.ProductCollection.InsertOne(context.TODO(),prod)
if err!=nil{
	json.NewEncoder(w).Encode("Unknown Error")
}else{
json.NewEncoder(w).Encode(prod)
}
}

// Delete route

func Deleteproduct(w http.ResponseWriter , r *http.Request){
enableCors(&w)
w.Header().Set("Content-Type", "application/json")
params	:= mux.Vars(r)
id := params["id"] 
_id , err := primitive.ObjectIDFromHex(id)
if err != nil{
	log.Fatal(err)
}
deleted , err := database.ProductCollection.DeleteOne(context.TODO(),bson.D{{"_id",_id}})
json.NewEncoder(w).Encode(deleted)
}

// get single product

func Getoneproducts(w http.ResponseWriter , r *http.Request){
		enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id  := params["id"]
	_id,err := primitive.ObjectIDFromHex(id)
	if err!=nil{
		log.Fatal(err)
	}
	var result models.Product
	err = database.ProductCollection.FindOne(context.TODO(),bson.D{{"_id",_id}}).Decode(&result)
	json.NewEncoder(w).Encode(result)
}

// Update product


func Updateproduct(w http.ResponseWriter, r *http.Request){
enableCors(&w)
w.Header().Set("Content-Type", "application/json")
params := mux.Vars(r)
id  := params["id"]
_id, _ := primitive.ObjectIDFromHex(id)
var prod models.Product
json.NewDecoder(r.Body).Decode(&prod)
filter := bson.M{"_id":_id}
update := bson.M{	"$set": bson.M{"name":prod.Name,"price":prod.Price,"description":prod.Description}}
database.ProductCollection.UpdateOne(context.TODO(),filter,update)
json.NewEncoder(w).Encode(prod)
}