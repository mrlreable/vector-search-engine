package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mrlreable/vector-search-engine/cmd/config"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/schema"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	c, err := config.ReadConfig(path + "/config.json")
	if err != nil {
		log.Fatal(err)
	}

	host := fmt.Sprintf("%s:%d", c.WeaviateHost, c.WeaviatePort)

	cfg := weaviate.Config{
		Host:   host,
		Scheme: "http",
	}

	wclient, schema, err := initWeaviateClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", schema)
	fmt.Printf("Client info: %#v\n", wclient)

}

func initWeaviateClient(cfg weaviate.Config) (*weaviate.Client, *schema.Dump, error) {
	wclient, err := weaviate.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	schema, err := wclient.Schema().Getter().Do(context.Background())

	return wclient, schema, err
}
