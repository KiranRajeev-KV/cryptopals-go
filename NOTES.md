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

- Example:  
  - Plaintext: "Hello"
  - ASCII codes: 72 101 108 108 111
  - Binary: `01001000 01100101 01101100 01101100 01101111`
    
  - Key: "K"
  - Key ASCII code: 75
  - Key Binary: `01001011`

  - XOR each byte of plaintext with key:
  - `01001000 ^ 01001011 = 00000011` (3)
  - `01100101 ^ 01001011 = 00101110` (46)
  - `01101100 ^ 01001011 = 00100111` (39)
  - `01101100 ^ 01001011 = 00100111` (39)
  - `01101111 ^ 01001011 = 00100100` (36)
  - Resulting ciphertext (in binary): `00000011 00101110 00100111 00100111 00100100`
  - ASCII codes of ciphertext: 3 46 39 39 36
  - Ciphertext characters: `\x03.''$`

The key can be easily found using frequency analysis or brute force, as there are only 256 possible keys (0-255).

