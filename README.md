# geistel
A fun little implementation of Feistel Cipher in Go

Online Demo: 
[https://repl.it/repls/SerpentineUnitedInstance](https://repl.it/repls/SerpentineUnitedInstance)

> Check out `main.go` for usage example.

### Run Demo Locally
- Clone repo.
- `cd` into it.
- `go run main.go geistel.go`

### Implementation Details

- Takes input and key as a `[]byte` slice and outputs enciphered/deciphered bytes as a `[]byte` slice.
- If the byte-length of complete input is odd, it appends the first byte of the key to make the input even. If the key is empty, then the null byte is appended. After enciphering, a null byte is again appended to the enciphered output to tell whether the actual input was odd or even when deciphering.
- The number of rounds depends on the byte-length of the key e.g. if the key is `abc`, the number of rounds will be `3`. Or if you use `SHA256` hash of your key, then the rounds will be `32` which is the byte-length of a SHA256 hash.
- The round function is very simple/dumb. It takes the `n`<sup>th</sup> byte of the key on `n`<sup>th</sup> round and takes `mod` against each byte of the current right-half of the input. See `roundFunction` in `geistel.go` for more details.
- The output from round function gets `XOR`ed with the current left-half of the input.
- The `XOR`ed output goes on to become the next right-half and the current right-half without round function and XOR, becomes the next left-half.

> Check out [This Diagram](https://en.wikipedia.org/wiki/Feistel_cipher#/media/File:Feistel_cipher_diagram_en.svg) and [this Wikipedia page](https://en.wikipedia.org/wiki/Feistel_cipher) and [this Computerphile YouTube video](https://www.youtube.com/watch?v=FGhj3CGxl8I) to understand the working of a Feistel Cipher.

> Note: Certainly not for production use.

author: [ahmednooor](https://github.com/ahmednooor)
