# Summary Day 5

## Pointer
Pointer adalah variabel yang digunakan untuk menyimpan alamat memory dari variabel\
yang lain. Dalam bahasa pemrograman Golang juga variabel ini digunakan untuk menyimpan\
beberapa data dialamaat memory tertentu disistem.

contoh implementasi:
```
    var a int = 5523
    
    //deklarasi pointer
    var p *int

    //inisialisai pointer dengan alamat dari variable a
    p = &a

    //inisialisasi pointer dengan value dari variable a
    p = *a
```

Merubah value dari pointer
```
    var a int = 3
    var b *int = &a

    //memasukkan value baru
    a = 10

    //variabel b akan terbarui valuenya menjadi 10
    fmt.Println("value baru b: ", *b)
```

## method
Fungsi method ini sama seperti difungsi-fungsi method dibahasa yang lain.

Implementasi dari define method:
```
    func (a string) sayHello() {
        fmt.Println("halo kamu, " a)
    }
```

Memanggil method
```
    func main() {
        var a1 = "Yurih"
        a1.sayHello()
    }
```

Method dengan pointer
```
    func (s *int) rubahAngka(angka int){
        fmt.Println("rubah angka: ", angka)
        s.num = angka
    }
```

## Public & Private Method
Pada bab ini memerlukan untuk membuat library untuk bisa menggunakan private maupun \
public. Untuk yang mendefinisikan public ditandai dengan huruf awal dari nama functionnya, \
sedangkan untuk private ditandai dengan huruf kecil.

Create library
```
    package library

    import "fmt"

    func SayHello(){
        fmt.Println("hello")
    }

    func introduce(name string) {
        fmt.Println("nama saya ", name)
    }
```

Memanggil function
```
    package main

    import "library"

    func main(){
        library.SayHello()
        library.introduce("Yurih")
    }
```

## Interface
Interface adalah custom type berupa kumpulan dari 1 atau lebih method signatures. Interface \
adalah abstract, tidak dapat membuat instance dari interface.

Define Interface
```
    type menhitung interface{
        luas() float64
        keliling() float64
    }
```

Buat method
```
    type lingkaran struct{
        diameter float64
    }

    func (l lingkaran) jariJari() float64{
        return l.diameter / 2
    }
    
    func (l lingkaran) luas() float64{
        return math.Pi * math.Pow(l.jariJari(), 2)
    }

    func (l lingkaran) keliling() float64 {
        return math.Pi * math.Pow(l.jariJari*(), 2)
    }

    type persegi struct {
        sisi float64
    }

    func (p persegi) luas() float64 {
        return math.Pow(p.sisi, 2)
    }

    func (p persegi) keliling() float64 {
        return p.sisi * 4
    }
```

Memanggil interface ke function main
```
    var bangunDatar = persegi{10, 0}
    fmt.Println(bangunDatar.luas())
    fmt.Println(bangunDatar.keliling())

    bangunDatar = lingkaran(13.0)
    fmt.Println(bangunDatar.luas())
    fmt.Println(bangunDatar.keliling())
    fmt.Println(bangunDatar.(lingkaran).jariJari())
```

## GoRoutine
Goroutine merupakan salah satu bagian paling penting dalam concurrent programming di Go. Salah\
satu yang membuat goroutine sangat istimewa adalah eksekusi-nya dijalankan di multi core processor.\
Kita bisa tentukan berapa banyak core yang aktif, makin banyak akan makin cepat.

Implementasi Goroutine
```
    import {
        "fmt"
        "time"
    }

    func number(){
        for i := 1; i <=5; i++{
            time.Sleep(250 * time.Millisecond)
            fmt.Printf("%d ", i)
        }
    }

    func alphabets(){
        for i := a; i <=e; i++{
            time.Sleep(400 * time.Millisecond)
            fmt.Printf("%c ", i)
        } 
    }

    func main(){
        go number()
        go alphabets()
        time.Sleep(3000 * time.Millisecond)
        fmt.Println("main terminated")
    }

    //output:
    // 1 a 2 3 b 4 c 5 d e main terminated
```

Implementasi Goroutin dengan WaitGroup
```
    import {
        "fmt"
        "sync"
    }

    func main(){
        wg := sync.WaitGroup{}
        wg.Add{2}
        go func(){
            defer wg.Done{}
            fmt.Println("hallo")
        }()

        go func(){
            defer.wg.Done{}
            fmt.Println("world")
        }()

        go func(){
            defer.wg.Done{}
            fmt.Println("Yurih")
        }()

        wg.Wait()
    }
```
