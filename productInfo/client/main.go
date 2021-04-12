package main

import (
	"context"
	"log"
	"time"

	pb "productInfo/client/ecommerce"

	"google.golang.org/grpc"
)

const (
	address = "localhost:13001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

	name := "I like, i love"
	description := "I like, i love, all is ok."
	price := float32(100.00)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatal("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Cound not get product: %v", err)
	}
	log.Printf("Product: %v", product.String())
}
