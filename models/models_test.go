package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestDbStruct(t *testing.T) {

	tests := []struct {
		Req []byte
		Res DbStruct
	}{
		{
			Req: []byte(`{"ID":1,"FirstName":"Иванов","LastName":"Иван","MiddleName":"Иванович","BDate":"31.12.1922","Addres":"Москва","Department":"HR","AboutMe":"я","Tnumber":"71234567890","Email":"exaple@ex.ex"}`),
			Res: DbStruct{
				1,
				"Иванов",
				"Иван",
				"Иванович",
				"31.12.1922",
				"Москва",
				"HR",
				"я",
				"71234567890",
				"exaple@ex.ex",
			},
		},
		{
			Req: []byte(`{"FirstName":"Иванов","LastName":"Иван","BDate":"31.12.1922","Addres":"Москва","Department":"HR","AboutMe":"я","Tnumber":"71234567890","Email":"exaple@ex.ex"}`),
			Res: DbStruct{
				FirstName:  "Иванов",
				LastName:   "Иван",
				BDate:      "31.12.1922",
				Address:    "Москва",
				Department: "HR",
				AboutMe:    "я",
				Tnumber:    "71234567890",
				Email:      "exaple@ex.ex",
			},
		},
	}

	for i, r := range tests {
		var st DbStruct
		if err := json.Unmarshal(r.Req, &st); err != nil {
			t.Errorf("Error test %d\nJson unmarshal error: %v", i, err)
		}
		if !reflect.DeepEqual(st, r.Res) {
			t.Errorf("Error test %d\n\n\tres:%v\n\t\tvalid res:%v\n", i, st, r.Res)
		}
	}

}

func TestDbStruct_Validate(t *testing.T) {
	tests := [][]struct {
		Req DbStruct
		Res error
	}{
		{
			{
				Req: DbStruct{
					FirstName: "Иванов",
					Email:     "exaple@ex.ex",
				},
				Res: nil,
			},
			{
				Req: DbStruct{
					FirstName: "Ivanov",
					Email:     "exaple@ex.ex",
				},
				Res: fmt.Errorf("FirstName ERROR"),
			},
		},
		{
			{
				Req: DbStruct{
					LastName: "Иван",
				},
				Res: nil,
			},
			{
				Req: DbStruct{
					LastName: "Ivan",
				},
				Res: fmt.Errorf("LastName ERROR"),
			},
		},
		{
			{
				Req: DbStruct{
					MiddleName: "Иванович",
				},
				Res: nil,
			},
			{
				Req: DbStruct{
					MiddleName: "Ivanovich",
				},
				Res: fmt.Errorf("MiddleName ERROR"),
			},
		},
		{
			{
				Req: DbStruct{
					BDate: "12.31.1998",
				},
				Res: nil,
			},
			{
				Req: DbStruct{
					BDate: "02.29.1999",
				},
				Res: fmt.Errorf("Date ERROR "),
			},
		},
		{
			{
				Req: DbStruct{
					Address: "Измайловское ш., 73 Ж, Москва, 105122",
				},
				Res: nil,
			},
			{
				Req: DbStruct{
					Address: "Moscow",
				},
				Res: fmt.Errorf("Address ERROR "),
			},
		},
		{
			{
				Req: DbStruct{
					Department: "HR",
				},
				Res: nil,
			},
			{
				Req: DbStruct{
					Department: "12323",
				},
				Res: fmt.Errorf("Department ERROR "),
			},
		},
		{
			{
				Req: DbStruct{
					AboutMe: "в свободное время",
				},
				Res: nil,
			},
			{
				Req: DbStruct{
					AboutMe: "in free time",
				},
				Res: fmt.Errorf("AboutMe ERROR"),
			},
		},
		{
			{
				Req: DbStruct{
					Tnumber: "71111111111",
				},
				Res: nil,
			},
			{
				Req: DbStruct{
					Tnumber: "+71111111111",
				},
				Res: fmt.Errorf("Phone number ERROR "),
			},
		},
		{
			{
				Req: DbStruct{
					Email: "exaple@ex.ex",
				},
				Res: nil,
			},
			{
				Req: DbStruct{
					Email: "exapleex.ex",
				},
				Res: fmt.Errorf("Email ERROR "),
			},
		},
	}

	for i, r := range tests {
		for i2, r2 := range r {
			err := r2.Req.Validate()
			if r2.Res != nil && err != nil {
				if err.Error() != r2.Res.Error() {
					t.Errorf("Test %d.%d filed\n\tres:%v\n\t\tvalid res:%v\n", i+1, i2+1, err, r2.Res)
				}
			} else {
				if !errors.Is(r2.Res, err) {
					t.Errorf("Test %d.%d filed\n\tres:%v\n\t\tvalid res:%v\n", i+1, i2+1, err, r2.Res)
				}
			}
		}
	}
}
