package main

import (
	"fmt"
	"sort"
)

type User struct {
	ID       int
	Name     string
	Role     string
	Questions []Question
}

type Question struct {
	ID       int
	Content  string
	Tags     []string
	Replies  []Reply
}

type Reply struct {
	ID      int
	Content string
}

type Forum struct {
	Questions []Question
}

func RegisterPatient(users *[]User, name string) {
	userID := len(*users) + 1
	patient := User{
		ID:   userID,
		Name: name,
		Role: "patient",
	}
	*users = append(*users, patient)
	fmt.Println("Pasien", name, "telah terdaftar.")
}

func PostQuestion(forum *Forum, userID int, content string, tags []string) {
	questionID := len(forum.Questions) + 1
	question := Question{
		ID:      questionID,
		Content: content,
		Tags:    tags,
		Replies: []Reply{},
	}
	forum.Questions = append(forum.Questions, question)
	fmt.Println("Pasien telah memposting pertanyaan.")
}

func ReplyToQuestion(forum *Forum, questionID int, userID int, content string) {
	for i := range forum.Questions {
		if forum.Questions[i].ID == questionID {
			replyID := len(forum.Questions[i].Replies) + 1
			reply := Reply{
				ID:      replyID,
				Content: content,
			}
			forum.Questions[i].Replies = append(forum.Questions[i].Replies, reply)
			fmt.Println("Dokter atau pasien telah memberikan tanggapan.")
			return
		}
	}
	fmt.Println("Pertanyaan tidak ditemukan.")
}

func ViewForum(forum Forum) {
	fmt.Println("=== Forum Konsultasi ===")
	for _, question := range forum.Questions {
		fmt.Println("Pertanyaan:", question.Content)
		fmt.Println("Tanggapan:")
		for _, reply := range question.Replies {
			fmt.Println(reply.Content)
		}
		fmt.Println()
	}
}

func SearchQuestionByTag(forum Forum, tag string) {
	fmt.Println("=== Pencarian berdasarkan Tag:", tag, "===")
	for _, question := range forum.Questions {
		for _, t := range question.Tags {
			if t == tag {
				fmt.Println("Pertanyaan:", question.Content)
				fmt.Println("Tanggapan:")
				for _, reply := range question.Replies {
					fmt.Println(reply.Content)
				}
				fmt.Println()
				break
			}
		}
	}
}

func SortTagsByQuestionCount(forum Forum) {
	tagCounts := make(map[string]int)
	for _, question := range forum.Questions {
		for _, tag := range question.Tags {
			tagCounts[tag]++
		}
	}

	type TagCount struct {
		Tag   string
		Count int
	}

	tagCountsSlice := []TagCount{}
	for tag, count := range tagCounts {
		tagCountsSlice = append(tagCountsSlice, TagCount{Tag: tag, Count: count})
	}

	// Menggunakan selection sort untuk mengurutkan berdasarkan jumlah pertanyaan
	for i := 0; i < len(tagCountsSlice)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(tagCountsSlice); j++ {
			if tagCountsSlice[j].Count > tagCountsSlice[minIndex].Count {
				minIndex = j
			}
		}
		tagCountsSlice[i], tagCountsSlice[minIndex] = tagCountsSlice[minIndex], tagCountsSlice[i]
	}

	fmt.Println("=== Topik/Tag Terurut berdasarkan Jumlah Pertanyaan ===")
	for _, tc := range tagCountsSlice {
		fmt.Println("Tag:", tc.Tag, "Jumlah Pertanyaan:", tc.Count)
	}
}

func SequentialSearchTag(forum Forum, tag string) bool {
	for _, question := range forum.Questions {
		for _, t := range question.Tags {
			if t == tag {
				return true
			}
		}
	}
	return false
}

func BinarySearchTag(tags []string, tag string) bool {
	low := 0
	high := len(tags) - 1

	for low <= high {
		mid := (low + high) / 2
		if tags[mid] == tag {
			return true
		} else if tags[mid] < tag {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	forum := Forum{
		Questions: []Question{},
	}
	users := []User{}

	// Mendaftarkan pengguna sebagai pasien
	RegisterPatient(&users, "John")

	// Pasien membuat pertanyaan
	PostQuestion(&forum, 1, "Bagaimana cara mengurangi stres?", []string{"stres", "kesehatan mental"})

	// Dokter atau pasien memberikan tanggapan
	ReplyToQuestion(&forum, 1, 2, "Anda bisa mencoba melakukan yoga atau meditasi untuk mengurangi stres.")

	// Melihat forum konsultasi
	ViewForum(forum)

	// Mencari pertanyaan berdasarkan tag
	SearchQuestionByTag(forum, "stres")

	// Menampilkan topik/tag terurut berdasarkan jumlah pertanyaan
	SortTagsByQuestionCount(forum)

	// Sequential Search Tag
	fmt.Println("=== Sequential Search ===")
	fmt.Println("Hasil Pencarian (Sequential):", SequentialSearchTag(forum, "kesehatan mental"))
	fmt.Println("Hasil Pencarian (Sequential):", SequentialSearchTag(forum, "kesehatan fisik"))

	// Binary Search Tag (setelah dilakukan pengurutan)
	tags := []string{"kesehatan fisik", "kesehatan mental", "stres"}
	sort.Strings(tags)
	fmt.Println("=== Binary Search ===")
	fmt.Println("Hasil Pencarian (Binary):", BinarySearchTag(tags, "kesehatan mental"))
	fmt.Println("Hasil Pencarian (Binary):", BinarySearchTag(tags, "stres"))
	fmt.Println("Hasil Pencarian (Binary):", BinarySearchTag(tags, "kesehatan spiritual"))
}
