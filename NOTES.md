# Notes :)

#### Hex Encoding

Hex encoding means you convert your string to binary, then take 4-bit chunks and convert each chunk to hexadecimal.

- Example: The character **'H'**  
  - ASCII code: 72  
  - Binary: `01001000`  
  - Split into 4-bit chunks: `0100` and `1000`  
  - Hexadecimal: `4` and `8`  
  - So, hex encoding of `'H'` is `"48"`

#### Base64 Encoding

Base64 encoding means converting your string to binary, then taking 6-bit chunks and converting each chunk to a Base64 character according to the encoding table.

- Example: The string **"KiR"**  
  - ASCII codes: 75 105 82  
  - Binary: `01001011 01101001 01010010`  
  - 6-bit chunks: `010010` `110110` `100101` `010010`  
  - Decimal: 18, 54, 37, 18  
  - Corresponding Base64 chars: `S3I=`

#### XOR Comparison

XOR is a bitwise operator (used to compare two bits) that returns 1 if one bit is 1 and the other is 0, otherwise it returns 0.

| Input A | Input B | Result |
|---------|---------|--------|
|    0    |    0    |   0    |
|    0    |    1    |   1    |
|    1    |    0    |   1    |
|    1    |    1    |   0    |

#### Single Byte XOR cipher

A Single Byte XOR cipher is a simple encryption technique where you XOR each byte of the plaintext with a single byte key.

```text
ciphertext = plaintext ^ key
plaintext  = ciphertext ^ key
```

- **Plaintext:** `"Hello"`
  - ASCII codes: `72 101 108 108 111`
  - Binary: `01001000 01100101 01101100 01101100 01101111`

- **Key:** `"K"`
  - ASCII code: `75`
  - Binary: `01001011`

Now, XOR each byte of the plaintext with the key:

| Plaintext Byte | Binary      | Key (K)     | XOR Result | Decimal | Ciphertext Char |
|:--------------:|:-----------:|:-----------:|:----------:|:-------:|:---------------:|
| H (72)         | 01001000    | 01001011    | 00000011   | 3       | `\x03`          |
| e (101)        | 01100101    | 01001011    | 00101110   | 46      | `.`             |
| l (108)        | 01101100    | 01001011    | 00100111   | 39      | `'`             |
| l (108)        | 01101100    | 01001011    | 00100111   | 39      | `'`             |
| o (111)        | 01101111    | 01001011    | 00100100   | 36      | `$`             |

- **Resulting ciphertext (binary):** `00000011 00101110 00100111 00100111 00100100`
- **Ciphertext (ASCII):** `\x03.''$`

> The key is easy to recover using frequency analysis or brute force, since there are only 256 possible keys (0-255).

