package test

import (
	"fruits_microservice/config"
	"fruits_microservice/fruits"
	"fruits_microservice/test/integration"
	"strconv"
	"testing"
)

func Test_Get_Fruits(t *testing.T) {
	// Arrange
	container, port, err := integration.RedisWithTestData()
	defer integration.DestroyContainer(container)
	if err != nil {
		t.Fatal(err)
	}

	config.RedisEndpoint = "localhost:" + strconv.Itoa(port)

	// Act
	username, fruit, err := fruits.GetFruit(1)

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	if username != "john" {
		t.Fatalf("Expected username to be john, got %s", username)
	}

	if fruit != "kiwi" {
		t.Fatalf("Expected fruit to be kiwi, got %s", fruit)
	}

	// Act
	userFruitMap, err := fruits.GetFruits()

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	if len(userFruitMap) != 4 {
		t.Fatalf("Expected 4 fruits, got %d", len(userFruitMap))
	}

	if _, ok := userFruitMap["john"]; !ok {
		t.Fatal("Expected john to be in the map")
		if userFruitMap["john"] != "kiwi" {
			t.Fatalf("Expected john's fruit to be kiwi, got %s", fruit)
		}
	}

	if _, ok := userFruitMap["damian"]; !ok {
		t.Fatal("Expected damian to be in the map")
		if userFruitMap["damian"] != "apple" {
			t.Fatalf("Expected damian's fruit to be apple, got %s", fruit)
		}
	}

}

func Test_Update_Fruits(t *testing.T) {
	// Arrange
	container, port, err := integration.RedisWithTestData()
	defer integration.DestroyContainer(container)

	if err != nil {
		t.Fatal(err)
	}

	config.RedisEndpoint = "localhost:" + strconv.Itoa(port)

	// Act
	err = fruits.UpdateFruit("user:3", "ariana", "orange")
	if err != nil {
		t.Fatal(err)
	}

	err = fruits.UpdateFruit("user:4", "damian", "banana")
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	userFruitMap, err := fruits.GetFruits()
	if err != nil {
		t.Fatal(err)
	}

	validMap := map[string]string{
		"john":      "kiwi",
		"johnathan": "apple",
		"ariana":    "orange",
		"damian":    "banana",
		"alexis":    "pineapple",
	}

	for k, v := range validMap {
		if _, ok := userFruitMap[k]; !ok {
			t.Fatalf("Expected %s to be in the map", k)
		}

		if userFruitMap[k] != v {
			t.Fatalf("Expected %s's fruit to be %s, got %s", k, v, userFruitMap[k])
		}
	}

}
