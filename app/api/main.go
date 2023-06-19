package main

import (
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
)

// Example Protobuf structure
type Person struct {
	Name  string `protobuf:"bytes,1,opt,name=name"`
	Email string `protobuf:"bytes,2,opt,name=email"`
}

func main() {
	http.HandleFunc("/json2protobuff", convertJSON2ProtobuffHandler)
	log.Println("Server started on http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (p *Person) ProtoMessage()  {}
func (p *Person) Reset()         { *p = Person{} }
func (p *Person) String() string { return proto.CompactTextString(p) }
