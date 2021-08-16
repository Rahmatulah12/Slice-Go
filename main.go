package main

import "fmt"

func main() {

	/*
		Slice
		Slice adalah referensi elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi
		sebuah array ataupun slice lainnya. Karena merupakan referensi, menjadikan perubahan data di tiap
		elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama
	*/
	var names = [4]string{"Rahmatulah", "Sidik", "Siapa", "Ya"}
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Mangga"
	fruits[2] = "Sirsak"
	fruits[3] = "Jeruk"

	// inisialisasi slice
	var slice = []string{"apple", "grape"}
	fmt.Println(slice)

	/*
		inisialisasi slice juga dapat dilakukan dengan mengambil slice/array yg sudah ada
		contoh : sliceFruits [dari index keberapa : sebelum index ke berapa]
				 sliceFruits[0:1]
	*/
	var sliceFruits = fruits[0:2]
	var sliceFruits2 = slice[1:2]
	// mengambil semua elemen dari index ke 1
	var sliceFruits3 = slice[1:]
	// mengambil semua elemen dari index ke 3
	var sliceNames = names[3:]
	fmt.Println(sliceFruits)
	fmt.Println(sliceFruits2)
	fmt.Println(sliceFruits3)
	fmt.Println(sliceNames)
	/*
		note slice merupakan tipe data reference
		Slice merupakan tipe reference. Artinya jika ada slice baru yang terbentuk dari slice lama,
		maka elemen slice baru memiliki referensi yang sama dengan elemen slice lama.
		Setiap perubahan yang terjadi di elemen slice baru, akan berdampak juga pada elemen slice lama
		yang memiliki referensi yang sama.
	*/

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]
	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	/*
		len()
		len() digunakan untuk menghitung lebar slice yang ada.
		Sebagai contoh jika sebuah variabel adalah slice dengan data 4 buah,
		maka fungsi ini pada variabel tersebut akan mengembalikan angka 4,
		yang angka tersebut didapat dari jumlah elemen yang ada.
	*/
	fmt.Println(len(slice))

	/*
		cap()
		cap() digunakan untuk menghitung lebar maksimum/kapasitas slice.
		Nilai kembalian fungsi ini awalnya sama dengan len ,
		tapi bisa berubah tergantung dari operasi slice yang dilakukan.
	*/
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	/*
		append()
		append() digunakan untuk menambahkan elemen pada slice.
		Elemen baru tersebut diposisikan setelah indeks paling akhir.
		Nilai balik fungsi ini adalah slice yang sudah ditambahkan nilai barunya.
	*/
	var cSlice = append(slice, "banana")
	fmt.Println(slice)
	fmt.Println(cSlice)

	/*
		copy()
		copy() digunakan untuk men-copy elemen slice tujuan (parameter ke-2), untuk digabungkan dengan slice target (parameter ke-1).
		Fungsi ini mengembalikan jumlah elemen yang berhasil di-copy (yang nilai tersebut merupakan nilai terkecil antara len(sliceTarget)
		dan len(sliceTujuan) )
	*/

	var childs = []string{"Aylin"}
	var nameChilds = []string{"Fariha", "Aylin Fariha"}

	var copy = copy(childs, nameChilds)

	fmt.Println(childs)
	fmt.Println(nameChilds)
	fmt.Println(copy)

	/*
		Pengaksesan Elemen Slice Dengan 3 Indeks
		3 index adalah teknik slicing elemen yang sekaligus menentukan kapasitasnya.
		Cara menggunakannnya yaitu dengan menyisipkan angka kapasitas di belakang, seperti fruits[0:1:1] .
		Angka kapasitas yang diisikan tidak boleh melebihi kapasitas slice yang akan di slicing
	*/
	var hoods = []string{"Brothers", "Hood", "Ok", "No"}
	var twoIndex = hoods[0:3]

	var threeIndex = twoIndex[0:2:2]

	fmt.Println(hoods)
	fmt.Println(len(hoods))
	fmt.Println(cap(hoods))

	fmt.Println(twoIndex)
	fmt.Println(len(twoIndex))
	fmt.Println(cap(twoIndex))

	fmt.Println(threeIndex)
	fmt.Println(len(threeIndex))
	fmt.Println(cap(threeIndex))

}
