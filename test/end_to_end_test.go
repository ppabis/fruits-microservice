package test

import (
	"fruits_microservice/test/e2e"
	"testing"
)

func TestEndToEnd(t *testing.T) {
	err := e2e.Setup()
	defer e2e.Teardown()

	if err != nil {
		t.Errorf("Error while setting up: %v", err)
	}

	e2e.Get_All_Fruits_Default(t)
	e2e.Get_One_Users_Fruit(t)
	e2e.Get_Inexistent_Users_Fruit(t)
	e2e.Set_Fruit_With_Good_Key(t)
	e2e.Get_One_Users_Fruit_After_Set(t)
	e2e.Set_Fruit_With_Bad_Key(t)
	e2e.Set_Fruit_With_Bad_Audience(t)
	e2e.Get_All_Fruits_After_Set(t)

}
