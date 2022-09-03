# Summary Day 10

## Design System

Desain system adalah standarisasi desain reusable yang dapat digunakan dalam mengembangkan frontend dari aplikasi dalam bentuk component. Kelebihan dalam menggunakan design system adalah lebih cepat, konsisten, dan dapat menjadi referensi oleh junior engineer

## OneFish Design System

OneFish design merupakan design system yang digunakan oleh tim internal eFishery dalam mengembangkan aplikasi internal. Design system ini terinspirasi dari ChakraUI dan berbasis ReactJS.

## ReactJS

ReactJS merupakan library Javascript yang digunakan dalam pengembangan UI suatu aplikasi. Salah satu cara instalasinya adalah dengan 
```npx create-react-app app-name```

## Cara Penggunaan

Tidak semua OneFish Design System bisa diimplementasikan, karena membutuhkan tim token internal eFishery agar bisa digunakan. Namun penggunaannya mirip dengan ChakraUI. Kemudian jika ada design system yang belum teraplikasikan ke OneFish, maka bisa menggunakan ChakraUI dengan penyesuaian design system eFishery.


## instalasi OneFish Design System:
```
- install reactjs
  npx create-react-app my-app

- install ChakraUI
  npm i @chakra-ui/react @emotion/react @emotion/styled framer-motion

- post install OneFish
  npm config set //balong.efishery.com/:_authToken="<token>"
  npm config set registry=https://balong.efishery.com
  npm config set always-auth=true

- install
  yarn add @efishery/onefish --registry https://balong.efishery.com
  yarn add @chakra-ui/react@^1.8.8 @chakra-ui/icons @chakra-ui/system @emotion/react@^11 @emotion/styled@^11 framer-motion@^4.1.17 @fontsource/poppins
```