package main

import "fmt"

type User struct {
	Username string
	Password string
	Role     string
	Email    string
}

type UserARR [50]User

type Forum struct {
	Pertanyaan string
	Jawaban    string
}

type ArrForum [50]Forum

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
	var F ArrForum

	for inputPilihanMenu != "5" {
		menu()
		fmt.Scan(&inputPilihanMenu)
		if inputPilihanMenu == "1.Registrasi" || inputPilihanMenu == "1" {
			Registrasi(*&A, *&n)
		} else if inputPilihanMenu == "2.Login" || inputPilihanMenu == "2" {
			login(*&A, *&n)
		} else if inputPilihanMenu == "3.Lihat Forum Konsultasi" || inputPilihanMenu == "3" {
			forum(&F)
		}
	}
}

func search(A UserARR, username string, password string, n int) bool {
	var found bool = false
	var i int
	//menggunakan sequential
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
	var F ArrForum

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
		forum(&F)
	} else {
		forum(&F)
	}
}

func forum(F *ArrForum) {
	var pertanyaan, inputPilihan string
	var n int
	var A UserARR

	fmt.Println("Bagaimana Cara Mencegah Kanker?")
	fmt.Println("Mencegah kanker melibatkan berbagai langkah sehat yang dapat diambil dalam kehidupan sehari-hari. Berikut ini adalah beberapa langkah yang dapat membantu dalam mencegah kanker:")
	fmt.Println("1. Hentikan kebiasaan merokok: Merokok diketahui menjadi salah satu faktor risiko utama untuk berbagai jenis kanker, termasuk kanker paru-paru, mulut, tenggorokan, dan pankreas.")
	fmt.Println("2. Konsumsi makanan sehat: Makan pola makan yang seimbang dan kaya akan buah-buahan, sayuran, biji-bijian utuh, dan sumber protein sehat dapat membantu mengurangi risiko kanker.")
	fmt.Println("3. Pertahankan berat badan yang sehat: Obesitas dikaitkan dengan peningkatan risiko kanker tertentu, seperti kanker payudara, usus besar, rahim, dan ginjal.")
	fmt.Println(" ")
	fmt.Println("Bagaimana Cara Mengatasi Sesak Nafas?")
	fmt.Println("Sesak nafas adalah kondisi yang serius dan memerlukan perhatian medis segera. Namun, jika Anda mengalami sesak nafas ringan atau sementara. Beberapa langkah yang dapat diambil:")
	fmt.Println("1. Bernafas perlahan dan dalam: Cobalah untuk bernafas secara perlahan dan dalam melalui hidung. ")
	fmt.Println("2. Hindari pemicu yang diketahui: Jika Anda memiliki alergi atau sensitivitas tertentu yang diketahui memicu sesak nafas, hindari paparan terhadap pemicu tersebut sesuai kemampuan Anda.")
	fmt.Println(" ")
	fmt.Println("Apakah Hepatitis Dapat Diobati? Jika Bisa, Apa Hal Yang Harus Dilakukan?")
	fmt.Println("Penyakit hepatitis dapat diobati tergantung pada jenis hepatitis yang Anda miliki. Ada beberapa jenis hepatitis, Berikut ini adalah informasi mengenai pengobatan masing-masing jenis hepatitis:")
	fmt.Println("")
	fmt.Println("Silahkan Masukkan Pertanyaan: ")

	fmt.Scan(&pertanyaan)
	fmt.Println(" ")
	fmt.Println("Silahkan Memilih")
	fmt.Println("1.Registrasi, 2.Keluar")

	fmt.Scan(&inputPilihan)

	if inputPilihan == "1" || inputPilihan == "1.Registrasi" {
		Registrasi(&A, &n)
	} else if inputPilihan == "2" || inputPilihan == "2.Keluar" {
		menu()
	}
}

func main() {
	var A UserARR
	var n int

	LoopMenu(&A, &n)

}
