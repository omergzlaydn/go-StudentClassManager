package main

import "fmt"

// Sınıfları tanımla
var class1 = map[string]bool{
	"Ali": true,
	"Burak":   true,
}

var class2 = map[string]bool{
	"Cenk": true,
	"Davut":   true,
}

var class3 = map[string]bool{
	"Elif":    true,
	"Furkan":  true,
}

var class4 = map[string]bool{
	"Gül":   true,
	"Helin":   true,
}

// Öğrencileri bir üst sınıfa taşıyan fonksiyon
func promoteStudents(currentClass, nextClass *map[string]bool) {
	for student := range *currentClass {
		(*nextClass)[student] = true
		delete(*currentClass, student)
	}
}

// Belirli bir sınıftaki öğrencileri listeleyen fonksiyon
func listStudents(class *map[string]bool) {
	for student := range *class {
		fmt.Println(student)
	}
}

// Tüm sınıfları ve içindeki öğrencileri listeleyen fonksiyon
func listAllClasses(classes ...*map[string]bool) {
	for i, class := range classes {
		fmt.Printf("Class %d students:\n", i+1)
		listStudents(class)
		fmt.Println()
	}
}

// Pointer'ın işaret ettiği sınıfı değiştiren fonksiyon (Güncellenmiş)
func changeClassPointer(class *map[string]bool, newClass *map[string]bool) {
	*class = *newClass
}

func main() {
	fmt.Println("Before Promotion:")
	listAllClasses(&class1, &class2, &class3, &class4)

	// Öğrencileri bir üst sınıfa taşıma
	promoteStudents(&class1, &class2)
	promoteStudents(&class2, &class3)
	promoteStudents(&class3, &class4)

	fmt.Println("After Promotion:")
	listAllClasses(&class1, &class2, &class3, &class4)

	// Double pointer kullanımına gerek kalmadı, bu yüzden normal pointer kullanıyoruz.
	fmt.Println("Changing class pointer of class1 to class3:")
	changeClassPointer(&class1, &class3)
	listAllClasses(&class1, &class2, &class3, &class4)
}
