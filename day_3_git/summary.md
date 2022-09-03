# Summary Day 3

Halo, pada summary day-3 ini membahas database. Database ini memiliki banyak jenisnya, yakni Relational, Non-Relational, dan Rational.\
Kali ini yang akan dibahas adalah Database Relational, ciri-ciri utamanya adalah memiliki entitas-entitas database yang berhubungan\
satu sama lain. Selain itu, Database Relational terdiri dari kolom dan baris pada tablenya. Berikut beberapa teknologi yang digunakan\
pada Database Relational yakni, mysql, oracle, dan postgresql. Yang paling membedakan adalah postgresql bisa melakukan auto increment\
daripada oracle. 

## Data Definition Language (DDL)

DDL ini berfungsi untuk pendefinisian struktur database, table, dan lain sebagainya. Adapun beberapa query pada DDL ini, yakni sebagai\
berikut: 
1. Create : Berfungsi untuk create database maupun table
```
    - create database database_name;
    - create table table_name (column_name tipe_data, column_name tipe_data);
```

2. Alter  : Berfungsi untuk mengedit definisi dari table
```
    - alter table table_name rename to table_name_change;
    - alter table table_name add column column_name tipe_data;
```

3. Drop : Berfungsi untuk menghapus database, table, bahkan function
```
    - drop database database_name;
    - drop table table_name;
    - drop function function_name;
```

## Data Manipulation Language (DML)

DML ini berfungsi untuk memanipulasi data yang ada didalam database. Adapun beberapa query pada DML ini, yakni sebagai berikut: 
1. Select : Berfungsi untuk memilih/menampilkan data dari sebuah table
```
    - select * from public.table_name;
    - select column_name from public.table_name;
```

2. Insert : Berfungsi untuk menginput data ke dalam table
```
    insert into pubilc.table_name (column_name, column_name) values (column_input), (column_input);
```

3. Update : Berfungsi untuk mengupdate/memperbarui data yang sudah terinput ditable
```
    update public.table_name set column_name = Inputan where id = inputan_yang_dituju;
```

4. Delete : Berfungsi untuk menghapus data yang sudah terinput ditable
```
    delete from table_name where id = inputan_yang_dituju;
```

## Data Control Language (DCL)

DCL ini berfungsi untuk mengontrol isi dari database, contohnya:
```
    - grant
    - revoke
    - skip
    - limit
```

## Join 

Join ini berfungsi sebagai query penghubung antara table satu ke table yang lain. Query join pun banyak macamnya tergantung\
dengan kebutuhannya. Berikut query join yang biasa digunakan:
```
    - inner join
    - left join
    - right join
    - full join
```

## Aggregation

Aggregation ini semacam query yang bisa digunakan sebagai operasi hitung atau bisa disebut sebagai operand pada query database.\
Berikut query aggregation:
```
    - COUNT : digunakan untuk menghitung jumlah row database
    - SUM   : digunakan untuk menghitung jumlah total dari column yang dipilih 
    - MAX   : digunakan untuk mengembalikan nilai terbesar dari column yang dipilih
    - MIN   : digunakan untuk mengembalikan nilai terkecil dari column yang dipilih
    - AVG   : digunakan untuk mengembalikan nilai dari rata-rata column yang dipilih
```

## Subquery

Subquery adalah query yang bersarang/nested query atau bisa disebut query yang ada didalam query. Berikut contohnya:
```
    update products set stock = subquery.stock - 2 from (select id, stock from products where id = 1) as subquery where
    products.id = 1;
```

## Function

Function adalah semacam stored procedures yang bisa digunakan untuk menghitung atau melakukan pengoperasian query saat dipanggil,\
atau bisa disebut dengan membuat semacam rumus yang bisa digunakan oleh kita untuk mempercepat melakukan query. Berikut contohnya:
```
    CREATE FUNCTION kurangi_stock(INT, INT) RETURNS products AS 'update products set stock = subquery.stock - $2 from (select id, 
    stock from products where id = $1) AS subquery where products.id = $1; select * from products where id = $1' LANGUAGE 'sql';
```