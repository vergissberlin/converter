package main

import (
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"
)

// Example Protobuf structure
type Person struct {
	Name  string `protobuf:"bytes,1,opt,name=name"`
	Email string `protobuf:"bytes,2,opt,name=email"`
}

func main() {

	// Use viper to read environment variables
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	http.HandleFunc("/json2protobuff", convertJSON2ProtobuffHandler)

	port := viper.GetString("PORT") // Heroku provides the port to bind to
	log.Println("Server started on http://0.0.0.0:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (p *Person) ProtoMessage()  {}
func (p *Person) Reset()         { *p = Person{} }
func (p *Person) String() string { return proto.CompactTextString(p) }
