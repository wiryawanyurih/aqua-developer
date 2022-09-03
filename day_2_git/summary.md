## Summary Day 2

Hola.. aku ingin membagikan beberapa pembelajaran di Day-2 aqua-developer. Kali ini pada hari ini membahas beberapa fitur\ 
dari github yang sering digunakan oleh kebanyakan developer yakni, Git Command, Commit Convention, Semantic Versioning, \
& Git Management. Kegunaan github sendiri ini adalah untuk membantu para developer dalam berkolaborasi, mengelola, tracking\
history, serta bisa digunakan sebagai tempat backup loh. Jadi, sudah kebayang kan kegunaannya cukup membantu buat para\
developer.

## Git Command

Git Command ini merupakan kumpulan command-command yang bisa digunakan oleh developer untuk mengelola file codingnya hingga\
terupload ke github. Berikut beberapa command yang biasa digunakan:

### **Init** 

Command ini memiliki fungsi untuk membuat repository di lokal  
Command: 
```
    git init <options>
    
    Example:
    1. git init --quiet (default)
    2. git init --bare
```

### **Add** 

Command ini memiliki fungsi untuk menambah perubahan pada file baru dari working directory ke stagging area
Command: 
```
    1. git add . (all file/folder/changes)
    2. git add <fileName>
    3. git add <directoryName>
```

### **Remote** 

Command ini memiliki fungsi untuk menyambungkan repository dengan remote
Command: 
```
    git remote <option>

    Example:
    1. git remote -v 
    2. git remote add <remote_name> <remote_url>
```

### **Status** 

Command ini memiliki fungsi untuk menampilkan status dari perubahan mana yang telah dilakukan, mana yang belum, dan file\
mana yang tidak dilacak oleh Git
Command: 
```
    1. git status
```

### **Commit** 

Command ini memiliki fungsi untuk menangkap snapshot dari perubahan yang dilakukan proyek saat ini
Command: 
```
    git commit <option>

    Example: 
    1. git commit -a
    2. git commit -m "message"
```

### **Push** 

Command ini memiliki fungsi untuk mengirim perubahan yang sudah di commit ke repositori melalui remote
Command: 
```
    git push <option>

    Example: 
    1. git push -u origin master
    2. git push -u origin 
```

### **Fetch** 

Command ini memiliki fungsi untuk mengambil info struktur meta-data terbaru dari github ke git lokal. Tetapi tidak\
melakukan transfer file atau lebih seperti memeriksa apakah ada perubahan yang tersedia
Command: 
```
    git fetch <remote> <option>

    Example: 
    1. git fetch --all
    2. git fetch origin
```

### **Pull** 

Command ini memiliki fungsi untuk mengambil dan men-download konten dari repositori remote dan  memperbarui repositori\
lokal agar sesuai dengan konten tersebut
Command: 
```
    git pull <remote> <branch_name>

    Example:
    1. git pull 
    2. git pull origin new_branch 
    3. git pull <remote repo>
```

### **Branch** 

Command ini memiliki fungsi abstraksi/sebagai identitas developer dalam melakukan development
Command: 
```
    git branch <branch_name>
    
    Example:
    1. git branch first_branch
    2. git branch -a
```

### **Merge** 

Command ini memiliki fungsi untuk menggabungkan beberapa commit menjadi satu
Command: 
```
    git merge <name-branch>

    Example:
    1. git merge first_branch
```

### **Stash** 

Command ini memiliki fungsi sebagai penyimpanan sementara ketika membuat perubahan, atau lebih mudahnya seperti draft\
diemail
Command: 
```
    git stash <command> - list, show, drop, pop, apply, clear, etc.

    Example: 
    1. git stash list 
    2. git stash apply <commit hash> 
```

## Commit Convention

Pada section ini menggunakan Git Karma Convention yang berfungsi untuk memberikan penamaan pada pesan commit agar lebih\
jelas dan mudah dipahami oleh developer

Commit values:
- feat     : digunakan untuk penambahan fitur
- fix      : digunakan untuk memperbaiki bug
- perf     : digunakan untuk kebutuhan peningkatan performance
- docs     : digunakan untuk penambahan dokumentasi
- style    : digunakan untuk menambahkan style
- refactor : digunakan untuk aktifitas perubahan di script
- test     : digunakan untuk menambahkan aktivitas testing
- build    : digunakan untuk update konfigurasi

## Semantic Versioning

Pada section ini merupakan rules dalam memberikan penamaan version dari apps, dalam aturannya seperti ini:

```
    v<major>.<minor>.<patch>

    Example:
    v1.2.3
```

Fungsi:
- Major : melakukan perubahan yang besar atau mengubah API
- Minor : feat
- Patch : fix, perf

Rules:
- Ketika ada update di major, maka minor dan patch harus direset menjadi 0
- Ketika ada update di minor, maka major tidak perlu reset namun patch harus direset menjadi 0
- Ketika ada update di patch, maka major dan minor tidak perlu reset 

## Git Management

Pada section ini membahas Git Management yakni salah satunya Trunk Base Development. Trunk Base Development ini adalah\ 
modeel pengontrol atau sebuah manajemen yang memudahkan developer dalam berkolaborasi disingle branch