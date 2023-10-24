package controller

import (
	"testing"
)

func TestPetController_UploadImage(t *testing.T) {
	//TODO
}

func TestPetController_CreatePet(t *testing.T) {
	tests := []struct {
		name    string
		expects []any
		wantErr bool
	}	{
		{
			name: "CreatePet",
			expects: []any{

			},
			wantErr: false,
		}
	}
}

func TestPetController_UpdatePet(t *testing.T) {

}

func TestPetController_GetPetsByStatuses(t *testing.T) {

}

func TestPetController_GetPetByID(t *testing.T) {

}

func TestPetController_UpdatePetByID(t *testing.T) {

}

func TestPetController_DeletePetByID(t *testing.T) {

}
