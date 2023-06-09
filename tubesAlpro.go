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

type menuBerForm struct {
	beranda  string
	forumArr [50]Forum
}

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
			Registrasi(A, n)
		} else if inputPilihanMenu == "2.Login" || inputPilihanMenu == "2" {
			login(*&A, *&n)
		} else if inputPilihanMenu == "3.Lihat Forum Konsultasi" || inputPilihanMenu == "3" {
			forum1(&F)
		}
	}
}

func menuBeranda(A *UserARR, M *menuBerForm, n *int, IndexUser int) {
	var inputMenu string
	var F ArrForum
	var i int
	var found int = -1

	for found == -1 && i < *n {
		if A[IndexUser].Role == "Pasien" || A[IndexUser].Role == "pasien" {
			found = i
			for inputMenu != "1" && inputMenu != "1.Beranda" && inputMenu != "2" && inputMenu != "2.Forum" {
				fmt.Println(" ")
				fmt.Println("Pasien")
				fmt.Println("========== MENU ==========")
				fmt.Println("1. Beranda")
				fmt.Println("2. Forum")
				fmt.Scan(&inputMenu)

				if inputMenu == "1.Beranda" || inputMenu == "1" {
					LoopMenu(A, n)
				} else if inputMenu == "2. Forum" || inputMenu == "2" {
					forum2(A, &F, n, IndexUser)
				}
			}
		} else if A[IndexUser].Role == "Dokter" || A[IndexUser].Role == "dokter" {
			found = i
			for inputMenu != "1" && inputMenu != "1.Beranda" && inputMenu != "2" && inputMenu != "2.Forum" && inputMenu != "3" && inputMenu != "1.TrendingTopik" {
				fmt.Println(" ")
				fmt.Println("Dokter")
				fmt.Println("========== MENU ==========")
				fmt.Println("1. Beranda")
				fmt.Println("2. Forum")
				fmt.Println("3. TrendingTopik")
				fmt.Scan(&inputMenu)
				if inputMenu == "1.Beranda" || inputMenu == "1" {
					LoopMenu(A, n)
				} else if inputMenu == "2. Forum" || inputMenu == "2" {
					forum2(A, &F, n, IndexUser)
				} else if inputMenu == "3.TrendingTopik" || inputMenu == "3" {
					tagTrend(M)
				}
			}
		}
		i++
	}
}

func search(A UserARR, username string, password string, n int) int {

	//menggunakan sequential
	for i := 0; i < n; i++ {
		if A[i].Username == username && A[i].Password == password {
			return i
		}
	}
	return -1
}

func login(A *UserARR, n *int) {
	var username, pswrd string
	//var kn, kr int

	var M menuBerForm

	fmt.Println("Masukan Sesuai Registrasi")
	fmt.Print("Username : ")
	fmt.Scan(&username)
	fmt.Print("Password : ")
	fmt.Scan(&pswrd)

	for search(*A, username, pswrd, *n) == -1 {
		fmt.Println("Maaf Username dan Password tidak sesuai")
		fmt.Println(" ")
		fmt.Println("Masukan Sesuai Registrasi")
		fmt.Print("Username : ")
		fmt.Scan(&username)
		fmt.Print("Password : ")
		fmt.Scan(&pswrd)
	}
	menuBeranda(A, &M, *&n, search(*A, username, pswrd, *n))
}

func InputMultipleString(value *string) {
	var inputRune rune

	*value = ""

	for string(inputRune) != "\n" {
		fmt.Scanf("%c", &inputRune)

		if string(inputRune) != "\n" {
			*value = *value + string(inputRune)
		}
	}
}

func forum1(F *ArrForum) {
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
	fmt.Println("Penyakit hepatitis dapat diobati tergantung pada jenis hepatitis yang dimiliki. Berikut ini adalah informasi mengenai pengobatan masing-masing jenis hepatitis:")
	fmt.Println("")
	fmt.Println("Silahkan Masukkan Pertanyaan: ")
	fmt.Scan(&pertanyaan)
	InputMultipleString(&pertanyaan)

	fmt.Println(" ")
	fmt.Println("Maaf Anda Harus Registrasi Dulu")
	fmt.Println("1.Registrasi, 2.Keluar")
	fmt.Scan(&inputPilihan)

	if inputPilihan == "1" || inputPilihan == "1.Registrasi" {
		Registrasi(&A, &n)
		login(&A, &n)
	} else if inputPilihan == "2" || inputPilihan == "2.Keluar" {
		LoopMenu(&A, &n)
	}
}

func forum2(A *UserARR, F *ArrForum, n *int, indexUser int) {
	var inpertanyaan, pilihMenu, inputPilih, ulang string
	// var A UserARR
	var M menuBerForm

	done := false
	for !done {
		fmt.Println("----------------")
		fmt.Println(" ")
		fmt.Println("1. Isiforum")
		fmt.Println("2. Trending")
		fmt.Println("3. Kembali")
		fmt.Println("----------------")
		fmt.Scan(&pilihMenu)

		if pilihMenu == "1" || pilihMenu == "1.Isiforum" {
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
			fmt.Println("Penyakit hepatitis dapat diobati tergantung pada jenis hepatitis yang dimiliki. Berikut ini adalah informasi mengenai pengobatan masing-masing jenis hepatitis:")
			fmt.Println("")

			doneQuestion := false
			for ! doneQuestion {
				fmt.Println("Silahkan Masukkan Pertanyaan: ")
				fmt.Scan(&inpertanyaan)
				InputMultipleString(&inpertanyaan)
				fmt.Println("----------------------------")
				fmt.Println("1. Submit, 2.Batal")
				fmt.Println("Pilihan :")
				fmt.Scan(&inputPilih)
				if inputPilih == "1" {
					F[*n].Pertanyaan = inpertanyaan
					*n++
					fmt.Println("Pertanyaan Tersimpan")
					fmt.Println(" ")
				} else if inputPilih == "2" {
					fmt.Println("pertanyaan Dibatalkan")
				}

				fmt.Print("Ingin masukkan pertanyaan lagi (Ya/Tidak)? ")
				fmt.Scan(&ulang)
				doneQuestion = ulang == "Tidak"
			}

		} else if pilihMenu == "2" || pilihMenu == "2.Trending" {
			searchTrend(&M)

		} else if pilihMenu == "3" || pilihMenu == "3.Kembali" {
			done = true
		}
	}

	menuBeranda(A, &M, *&n, indexUser)
}



func searchTrend(M *menuBerForm) {
	//menggunakan search binary
	fmt.Println("masuk")
}

func tagTrend(M *menuBerForm) {
	//membutuhkan sorting insertion atau selection
	fmt.Println("berhasil")
}
func main() {
	var A UserARR
	var n int
	LoopMenu(&A, &n)
}
