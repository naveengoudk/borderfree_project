package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)
func enableCors(w *http.ResponseWriter) {
(*w).Header().Set("Access-Control-Allow-Origin", "*")
(*w).Header().Set("Access-Control-Allow-Headers", "authentication, Content-Type")
}

// database connection
func db() *mongo.Client{
	clientOptions := options.Client().ApplyURI("mongodb+srv://naveengoud:Naveen@cluster0.fkhjj.mongodb.net/Borderfree?retryWrites=true&w=majority")
	client ,err := mongo.Connect(context.TODO(), clientOptions)
	if err!=nil{
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(),nil)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Print("mongodb connected")
	return client
}

// product model
type product struct{
	User   string `json:"user"`
	Name 	 string `json:"name"`
	Price  string `json:"price"`
	Description string `json:"description"`
}
type users struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

var productCollection = db().Database("Borderfree").Collection("products")
var userCollection = db().Database("Borderfree").Collection("users")

// signup route
func signup(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user users
	json.NewDecoder(r.Body).Decode(&user)
	var result	users
	err := userCollection.FindOne(context.TODO(),bson.M{"email":user.Email}).Decode(&result)
	if err != nil{
					BcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
			user.Password = string(BcryptPassword)
			cur ,err := userCollection.InsertOne(context.TODO(),user)
			if err	!= nil{
				log.Fatal(err)
			}
			json.NewEncoder(w).Encode(cur.InsertedID)
			} else {
		json.NewEncoder(w).Encode("User already exists with this email")
	}}

	// login route
	func login(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user users
	json.NewDecoder(r.Body).Decode(&user)
	var	result users
	err := userCollection.FindOne(context.TODO(),bson.M{"email":user.Email}).Decode(&result)
	if err  != nil{
		json.NewEncoder(w).Encode("User not found")
	}else{
	err = bcrypt.CompareHashAndPassword([]byte(result.Password),[]byte(user.Password))
	if err != nil{
		json.NewEncoder(w).Encode("Wrong password")
	}else{
	json.NewEncoder(w).Encode(result)
	}
}
}

// Forget password

type forgetuser struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

func forgetpassword(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	var user forgetuser
	json.NewDecoder(r.Body).Decode(&user)
	var olduser users
	err := userCollection.FindOne(context.TODO(),bson.D{{"email",user.Email}}).Decode(&olduser)
	if err	!= nil{
		json.NewEncoder(w).Encode("User not found")
	}else{
		BcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	json.NewEncoder(w).Encode(olduser)
	filter :=  bson.D{{"email",user.Email}}
		update :=  bson.M{"$set":bson.D{{"password",BcryptPassword}}}
	_,err := userCollection.UpdateOne(context.TODO(),filter,update)
	if err != nil{
		json.NewEncoder(w).Encode("Error updating password")
	}
	}
}
// getallproducts route

func getproducts(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	user :=  params["user"]
	fmt.Print(user)
	var products []primitive.M
	cur , err := productCollection.Find(context.TODO(),bson.D{{"user",user}})
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

func addproduct(w http.ResponseWriter , r *http.Request){
enableCors(&w)
w.Header().Set("Content-Type", "application/json")
var prod product
json.NewDecoder(r.Body).Decode(&prod)
fmt.Print(prod)
cur , err := productCollection.InsertOne(context.TODO(),prod)
if err!=nil{
	json.NewEncoder(w).Encode("Unknown Error")
}else{
json.NewEncoder(w).Encode(cur.InsertedID)
}
}

// Delete route

func deleteproduct(w http.ResponseWriter , r *http.Request){
enableCors(&w)
w.Header().Set("Content-Type", "application/json")
params	:= mux.Vars(r)
id := params["id"] 
_id , err := primitive.ObjectIDFromHex(id)
if err != nil{
	log.Fatal(err)
}
deleted , err := productCollection.DeleteOne(context.TODO(),bson.D{{"_id",_id}})
json.NewEncoder(w).Encode(deleted)
}

// get single product

func getoneproducts(w http.ResponseWriter , r *http.Request){
		enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id  := params["id"]
	_id,err := primitive.ObjectIDFromHex(id)
	if err!=nil{
		log.Fatal(err)
	}
	var result product
	err = productCollection.FindOne(context.TODO(),bson.D{{"_id",_id}}).Decode(&result)
	json.NewEncoder(w).Encode(result)
}

// Update product

func updateproduct(w http.ResponseWriter, r *http.Request){
		enableCors(&w)
w.Header().Set("Content-Type", "application/json")
params := mux.Vars(r)
id  := params["id"]
_id, _ := primitive.ObjectIDFromHex(id)
var prod product
json.NewDecoder(r.Body).Decode(&prod)
filter := bson.D{{"_id",_id}}
update := bson.D{{	"$set", bson.D{{"name",prod.Name},{"price",prod.Price}, {"description",prod.Description}}}}
productCollection.UpdateOne(context.TODO(),filter,update)
json.NewEncoder(w).Encode(prod)
}


func main() {
	os.Setenv("PORT", "8000")
	port := os.Getenv("PORT")
	fmt.Print(port)
	r := mux.NewRouter()
		r.HandleFunc("/", login).Methods("POST")
	r.HandleFunc("/signup",	signup).Methods("POST")
	r.HandleFunc("/forgetpassword",forgetpassword).Methods("POST")
	r.HandleFunc("/getproducts/{user}", getproducts).Methods("GET")
	r.HandleFunc("/getoneproducts/{id}", getoneproducts).Methods("GET")
	r.HandleFunc("/updateproduct/{id}", updateproduct).Methods("PUT")
	r.HandleFunc("/addproduct", addproduct).Methods("POST")
	r.HandleFunc("/deleteproduct/{id}", deleteproduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}