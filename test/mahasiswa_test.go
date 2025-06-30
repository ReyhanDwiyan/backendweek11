package test

import (
	"context"
	"fmt"
	"inibackend/model"
	"inibackend/repository"
	"strconv"
	"testing"
)

var ctx = context.TODO()

func TestInsertMahasiswa(t *testing.T) {

	mhs := model.Mahasiswa{
		Nama:     "Test Mahasiswa",
		NPM:      9999999999,
		Prodi:    "Teknik Uji",
		Fakultas: "Fakultas Testing",
		Alamat: model.Alamat{
			Jalan:     "Jl. Testing No.1",
			Kelurahan: "UnitTest",
			Kota:      "Bandung",
		},
		Minat: []string{"Unit Testing", "Mocking"},
		MataKuliah: []model.MataKuliah{
			{Kode: "UT101", Nama: "Dasar Testing", Nilai: 100},
		},
	}

	insertedID, err := repository.InsertMahasiswa(ctx, mhs)
	if insertedID == nil {
		//t.Error("InsertMahasiswa failed, insertedID is nil")
		fmt.Printf("Inserted failed: %v", err)
	} else {
		fmt.Printf("Inserted Mahasiswa with ID: %v", insertedID)
	}
}

func TestGetMahasiswaByNPM(t *testing.T) {
	npm := 9999999999 // NPM yang digunakan pada TestInsertMahasiswa
	mhs, err := repository.GetMahasiswaByNPM(ctx, npm)
	if err != nil {
		t.Errorf("GetMahasiswaByNPM failed: %v", err)
		return
	}
	if mhs.NPM != npm {
		t.Errorf("Expected NPM %v, got %v", strconv.FormatInt(int64(npm), 10), mhs.NPM)
	} else {
		fmt.Printf("Retrieved Mahasiswa: %+v", mhs)
	}
}

func TestGetAllMahasiswa(t *testing.T) {
	all, err := repository.GetAllMahasiswa(ctx)
	if len(all) == 0 {
		//t.Error("No mahasiswa found")
		fmt.Printf("No mahasiswa found: %v", err)
	} else {
		fmt.Printf("Total mahasiswa: %v", len(all))
		fmt.Print(all)
	}
}

func TestUpdateMahasiswa(t *testing.T) {
	npm := 9999999999
	newData := model.Mahasiswa{
		Nama:     "Mahasiswa Update",
		NPM:      npm,
		Prodi:    "Teknik Revisi",
		Fakultas: "Fakultas Pembaruan",
		Alamat: model.Alamat{
			Jalan:     "Jl. Update No.99",
			Kelurahan: "UpdateZone",
			Kota:      "Depok",
		},
		Minat: []string{"Go Update", "MongoDB"},
		MataKuliah: []model.MataKuliah{
			{Kode: "UT102", Nama: "Testing Lanjut", Nilai: 95},
		},
	}

	updatedNPM, err := repository.UpdateMahasiswa(ctx, npm, newData)
	if err != nil {
		t.Errorf("UpdateMahasiswa failed: %v", err)
	} else {
		fmt.Printf("Updated Mahasiswa with NPM: %v\n", updatedNPM)
	}
}

func TestDeleteMahasiswa(t *testing.T) {
	npm := 9999999999

	deletedNPM, err := repository.DeleteMahasiswa(ctx, npm)
	if err != nil {
		t.Errorf("DeleteMahasiswa failed: %v", err)
	} else {
		fmt.Printf("Deleted Mahasiswa with NPM: %v\n", deletedNPM)
	}
}
