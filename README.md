# Classic-Cryptography-Simulator-Backend
Back-end side of simulator for well-known classic cipher algorithms using Go languange

You can also check the front-end side repository by clicking [here](https://github.com/mikeleo03/Classic-Cryptography-Simulator-Frontend)

## Getting Started
Make sure you already install Go language on your computer before running the code

## ⚙️ &nbsp;How to Run
1. Clone this repository from terminal using this following command
    ``` bash
    git clone https://github.com/GoDillonAudris512/Classic-Cryptography-Simulator-Backend.git
    ```
2. Create a .env file inside the repository directory using .env.example file as the template. You can keep the variables blank. The server should automatically use port 8080 as the default port 
3. Navigate to this repository directory
4. Run the server using this following command
    ``` bash
    go run main.go
    ```
5. The back-end server should be running. You can also check the server by opening http://localhost:8080/api


## 🔑 &nbsp;List of Endpoints
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