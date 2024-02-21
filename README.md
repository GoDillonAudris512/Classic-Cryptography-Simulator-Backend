# 🔐 Classic Cryptography Simulator
> Back-end side of simulator for well-known classic cipher algorithms using Go languange

## General Information
This program is created to simulate some popular classic cryptography algorithms built under the web. Dive into the captivating world of classic cryptography with this interactive simulator! Experiment with renowned algorithms like Caesar and Vigenere Ciphers, encrypt messages beyond the alphabet with extended ASCII support, and handle both text snippets and entire files. Decrypt hidden messages, craft your own secret codes, and gain a deeper understanding of this fascinating field. Unleash your inner codebreaker and embark on this intriguing cryptographic journey!

## Project Structure
```bash
.
├─── lib
│   └─── matrix.go
├─── middleware
│   ├─── affineHandlers.go
│   ├─── autoVigenereHandlers.go
│   ├─── enigmaHandlers.go
│   ├─── extendedVigenereHandlers.go
│   ├─── generalHandlers.go
│   ├─── hillHandlers.go
│   ├─── playfairHandlers.go
│   ├─── superHandlers.go
│   └─── vigenereHandlers.go
├─── model
│   ├─── playfair.go
│   └─── token.go
├─── router
│   └─── router.go
├─── .env
├─── .env.example
├─── .gitignore
├─── go.mod
├─── go.sum
├─── main.go
└─── README.md
```

## User Interfaces
User Interface is designed and implemented on the front-end side. Further implementation stated on [this repository](https://github.com/mikeleo03/Classic-Cryptography-Simulator-Frontend)

## ⚙️ &nbsp;How to Run the Program

Clone this repository from terminal with this command
``` bash
$ git clone https://github.com/GoDillonAudris512/Classic-Cryptography-Simulator-Backend.git
```

### Run the application on development server
1. Create a .env file inside the repository directory using .env.example file as the template. You can keep the variables blank. The server should automatically use port 8080 as the default port 
2. Run the server using this following command
    ``` bash
    go run main.go
    ```

If you do it correctly, the back-end development server should be running. You can also check the server by opening http://localhost:8080/api. To use back-end side functionalities, don't forget to also run the front-end side. Further explanation on how to run the front-end development server stated on [this repository](https://github.com/mikeleo03/Classic-Cryptography-Simulator-Frontend)


## 🔑 &nbsp;Endpoints
| Endpoint                             |  Method  |   Usage  |
| ------------------------------------ | :------: | -------- |
| /api/vigenere                        | POST     | Users can perform encryption and decryption using Vigenere Cipher
| /api/auto-vigenere                   | POST     | Users can perform encryption and decryption using Auto-Key Vigenere Cipher
| /api/extended-vigenere               | POST     | Users can perform encryption and decryption using Extended Vigenere Cipher
| /api/playfair                        | POST     | Users can perform encryption and decryption using Playfair Cipher
| /api/affine                          | POST     | Users can perform encryption and decryption using Affine Cipher
| /api/hill                            | POST     | Users can perform encryption and decryption using Hill Cipher
| /api/super                           | POST     | Users can perform encryption and decryption using Super Encryption of Extended Vigenere Cipher and Columnar Transposition Cipher
| /api/enigma                          | POST     | Users can perform encryption and decryption using _Enigma Cipher_


## Authors
| Name                     |   Role   |  
| ------------------------ | -------- |
| Go Dillon Audris         | Back-end Engineer
| Michael Leon Putra Widhi | Front-end Engineer