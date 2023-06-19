package main

import (
	"log"
	"net/http"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/joho/godotenv"
)

// Example Protobuf structure
type Person struct {
	Name  string `protobuf:"bytes,1,opt,name=name"`
	Email string `protobuf:"bytes,2,opt,name=email"`
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/json2protobuff", convertJSON2ProtobuffHandler)

	port := os.Getenv("PORT") // Heroku provides the port to bind to
	log.Println("Server started on http://0.0.0.0:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (p *Person) ProtoMessage()  {}
func (p *Person) Reset()         { *p = Person{} }
func (p *Person) String() string { return proto.CompactTextString(p) }
