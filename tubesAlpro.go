package main

import "fmt"

type User struct {
	Username string
	Password string
	Role     string
	Email    string
}

type UserARR [50]User

func menu() {
	fmt.Println("          =====SELAMAT DATANG=====           ")
	fmt.Println("        DI KONSULTASI KESEHATAN KAMI         ")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("  Atasi Segera Penyakit Anda,Sayangilah Tubuh Anda   ")
	fmt.Println("=====================================================")
	fmt.Println("1. Registrasi")
	fmt.Println("2. Login")
	fmt.Println("3. Lihat Forum Konsultasi")
	fmt.Println("4. Edit Profile")
	fmt.Println("5. Keluar")
	fmt.Println("-----------------------------------------------------")
}

func Registrasi(A *UserARR, n *int) {
	var username, pswrd, role, email string
	var inputPilihan int
	fmt.Println("=====Menu Registrasi=====")
	fmt.Print("Email :")
	fmt.Scan(&email)
	fmt.Print("Username :")
	fmt.Scan(&username)
	fmt.Print("Password :")
	fmt.Scan(&pswrd)
	fmt.Print("Role :")
	fmt.Scan(&role)
	fmt.Println("----------------------------")
	fmt.Println("1. Simpan, 2.Batal")
	fmt.Println("Pilihan :")
	fmt.Scan(&inputPilihan)
	if inputPilihan == 1 {
		A[*n].Email = email
		A[*n].Username = username
		A[*n].Password = pswrd
		A[*n].Role = role
		*n++
		fmt.Println("Data tersimpan yah:)")
	} else if inputPilihan == 2 {
		fmt.Println("Data tidak jadi disimpan:(")
	}
	// fmt.Println(" ")
	// for i = 0; i < *n; i++ {
	// 	fmt.Println("Email :", A[i].Email)
	// 	fmt.Println("Username :", A[i].Username)
	// 	fmt.Println("Password :", A[i].Password)
	// 	fmt.Println("Role :", A[i].Role)
	// }
}

func LoopMenu(A *UserARR, n *int) {
	var inputPilihanMenu string

	for inputPilihanMenu != "5" {
		menu()
		fmt.Scan(&inputPilihanMenu)
		if inputPilihanMenu == "1.Registrasi" || inputPilihanMenu == "1" {
			Registrasi(*&A, *&n)
		} else if inputPilihanMenu == "2.Login" || inputPilihanMenu == "2" {
			login(*&A, *&n)
		}
	}
}

func search(A UserARR, username string, password string, n int) bool {
	var found bool = false
	var i int
	i = 0
	for !found && i < n {
		found = A[i].Username == username && A[i].Password == password
		i++
	}
	return found
}
func login(A *UserARR, n *int) {
	var username, pswrd string
	//var kn, kr int
	var found bool

	fmt.Println("Masukan Sesuai Registrasi")
	fmt.Print("Username : ")
	fmt.Scan(&username)
	fmt.Print("Password : ")
	fmt.Scan(&pswrd)
	found = search(*A, username, pswrd, *n)
	if !found {
		for !found {
			fmt.Println("Maaf Username dan Password tidak sesuai")
			fmt.Println(" ")
			fmt.Println("Masukan Sesuai Registrasi")
			fmt.Print("Username : ")
			fmt.Scan(&username)
			fmt.Print("Password : ")
			fmt.Scan(&pswrd)
			found = search(*A, username, pswrd, *n)
		}
		forum(A)
	} else {
		forum(A)
	}
}

func forum(A *UserARR) {
	fmt.Println("test 1,2,3")

}

func pertanyaan() {
	fmt.Println("Silahkan Masukkan Pertanyaan")
	fmt.Scan()
}

func main() {
	var A UserARR
	var n int

	LoopMenu(&A, &n)

}
