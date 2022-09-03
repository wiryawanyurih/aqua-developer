# Summary Day 4

## Sejarah
Go Lang atau biasa disebut Golang ini merupakan bahasa pemrograman yang dibuat oleh Google. \
Bahasa pemrograman ini didesain pada akhir 2007, kemudian baru mulai dibangun pada pertengahan\
tahun 2008. Setelah itu, baru merilis versi 1.0 pada tahun 2012, hingga saat ini go sudah merilis\
versi 1.19. Saat ini pemrograman ini cukup populer dikalangan perusahaan-perusahaan tech company besar,\
salah satunya adalah eFishery, Gojek, Tokopedia, Google, dan lain-lain.


## Tipe Data
Tipe Data yang bisa digunakan di GOlang ini beragam, yakni sebagai berikut:
```
    1. Numerical Non-Decimal
        - uint8
        - uint16
        - uint32
        - uint64
        - uint
        - byte
        - int8
        - int16
        - int 32
        - int64
        - int
        - rune
    2. Numerical Decimal
    3. Boolean
    4. String
```


## Variables 
Deklarasi variables 
```
    func main(){
        var firstName string = "Yurih"
        var lastName string
        lastName = "Kaloko"
    }
```

Deklarasi variables tanpa tipe data
```
    func main(){
        var firstName string = "Wiryawan"
        lastName := "Yurih"
    }
```

Deklarasi multi variables
```
    var satu, dua, tiga string
    satu, dua, tiga = "pertama", "kedua", "ketiga"
    lima, enam, tujuh := "five", "six", "seven"
```

Deklarasi underscore variables
```
    _ = "belajar"
    _ = "matematika"
    name, _ := "Yurih", "Kaloko"
```

Deklarasi constants
```
    const firstName string = "Yurih"
```


## Condition
if...else...then
```
    var angka = 4

    if point == 11 {
        fmt.println("lulus")
    } else if point > 5 {
        fmt.println("lulus nilai pas")
    } else {
        fmt.println("tidak lulus")
    }
```

Comparison operator
```
    1. &&
    2. ||
    3. !
```

Switch case
```
    var pointer = 7

    switch pointer {
        case 7:
            fmt.println("7")
        case 8: 
            fmt.println("8")
        default:
            fmt.println("not found")
    }
```


## Looping
for...range
```
    var makanan = [2]string{"Makaroni", "mie"}

    for i, makan := range makanan{
        fmt.println("makanan: ", makan)
    }
```

for...loop
```
    for i := 0; i < 5; i++{
        fmt.println("angka ", i)
    }
```

for...loop dengan break continue
```
    for i := 0; i < 5; i++{
        if i%2 == 1 {
            continue
        }

        if i > 4 {
            break
        }

        fmt.println("angka ", i)
    }
```


## function
Function
```
    func swap(x, y string) (string, string){
        return y, x
    }

    func main(){
        a, b := swap("hello", "world")
        fmt.println(a, b)
    }
```

Consecutive named functionn parameters
```
    //from
    func sum(x int, y int) int {
        return x + y
    }

    //to
    func sum(x, y int) int {
        return x + y
    }
```

Multiple return value
```
    func tukar(a, b string) (string, string){
        return b, a
    }
```


## struct
Struct
```
    type lecture struct {
	    name    string
	    address string
	    matkul  string
    }

    func main() {
	    var s1 lecture
	    s1.name = "Wiryawan"
	    s1.address = "Jl. Suhat"
	    s1.matkul = "Golang"
    }
```

Embedded struct
```
    type lecture struct {
	    name    string
	    address string
	    matkul  string
	    boot    academy
    }

    type academy struct {
	    name    string
	    address string
    }

    func main() {
	    var s1 lecture
	    s1.name = "Wiryawan"
	    s1.address = "Jl. Suhat"
	    s1.matkul = "Golang"
	    s1.boot = academy{
		    name:    "eFishery Academy",
		    address: "Jl. Kayu Tangan",
	}
```

Struct + Slice
```
    type lecture struct {
	    name    string
	    address string
	    matkul  string
    }

    func main(){
        var allLecture = []person{
            {name: "Yurih", address: "Jl. suhat", matkul: "Matematika"},
            {{name: "Kaloko", address: "Jl. sulfat", matkul: "Statistika"},}
        }
    }
```


## Data Structure
Map
```
    type academy struct {
	    name    string
	    address string
    }

    var kampus map[string]academy
	kampus = map[string]academy{}

	kampus["kampus_pusat"] = academy{
		name:    "Pusat",
		address: "Jl. Kayu tanggan, Kota Malang",
	}
```

Array
```
    var daftar_matkul = [3]string{"Golang", "PHP", "Java"}
```


## Slice
Slice
```
    var kurikulum = []string{"2018", "2019", "2020"}
```

Defer
```
    defer fmt.println("Yurih")
    fmt.println("Halo")

    /*output:
    Halo
    Yurih
    */
```