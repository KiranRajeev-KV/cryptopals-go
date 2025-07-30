# Notes :)

### Hex Encoding

Hex encoding means you convert your string to binary, then take 4-bit chunks and convert each chunk to hexadecimal.

- Example: The character **'H'**  
  - ASCII code: 72  
  - Binary: `01001000`  
  - Split into 4-bit chunks: `0100` and `1000`  
  - Hexadecimal: `4` and `8`  
  - So, hex encoding of `'H'` is `"48"`

### Base64 Encoding

Base64 encoding means converting your string to binary, then taking 6-bit chunks and converting each chunk to a Base64 character according to the encoding table.

- Example: The string **"KiR"**  
  - ASCII codes: 75 105 82  
  - Binary: `01001011 01101001 01010010`  
  - 6-bit chunks: `010010` `110110` `100101` `010010`  
  - Decimal: 18, 54, 37, 18  
  - Corresponding Base64 chars: `S3I=`

### XOR Comparison

XOR is a bitwise operator (used to compare two bits) that returns 1 if one bit is 1 and the other is 0, otherwise it returns 0.

| Input A | Input B | Result |
| ------- | ------- | ------ |
| 0       | 0       | 0      |
| 0       | 1       | 1      |
| 1       | 0       | 1      |
| 1       | 1       | 0      |

### Single Byte XOR cipher

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

| Plaintext Byte |  Binary  | Key (K)  | XOR Result | Decimal | Ciphertext Char |
| :------------: | :------: | :------: | :--------: | :-----: | :-------------: |
|     H (72)     | 01001000 | 01001011 |  00000011  |    3    |     `\x03`      |
|    e (101)     | 01100101 | 01001011 |  00101110  |   46    |       `.`       |
|    l (108)     | 01101100 | 01001011 |  00100111  |   39    |       `'`       |
|    l (108)     | 01101100 | 01001011 |  00100111  |   39    |       `'`       |
|    o (111)     | 01101111 | 01001011 |  00100100  |   36    |       `$`       |

- **Resulting ciphertext (binary):** `00000011 00101110 00100111 00100111 00100100`
- **Ciphertext (ASCII):** `\x03.''$`

The key is easy to recover using frequency analysis or brute force, since there are only 256 possible keys (0-255).

### Repeated Key XOR cipher

A Repeated Key XOR cipher is a more complex encryption technique where you XOR each byte of the plaintext with a byte from a repeating key.

- **Plaintext:** `"Hello"`
  - ASCII codes: `72 101 108 108 111`
  - Binary: `01001000 01100101 01101100 01101100 01101111`

*- **Key:** `"Key"`
  - ASCII codes: `75 101 121`
  - Binary: `01001011 01100101 01111001`
  - Repeats: `01001011 01100101 01111001 01001011 01100101`

Now, XOR each byte of the plaintext with the corresponding byte from the repeating key:

| Plaintext Byte |  Binary  | Key (K)  | XOR Result | Decimal | Ciphertext Char |
| :------------: | :------: | :------: | :--------: | :-----: | :-------------: |
|     H (72)     | 01001000 | 01001011 |  00000011  |    3    |     `\x03`      |
|    e (101)     | 01100101 | 01100101 |  00000000  |    0    |     `\x00`      |
|    l (108)     | 01101100 | 01111001 |  00010101  |   21    |     `\x15`      |
|    l (108)     | 01101100 | 01001011 |  00100111  |   39    |       `'`       |
|    o (111)     | 01101111 | 01100101 |  00001010  |   10    |      `\n`       |

- **Resulting ciphertext (binary):** `00000011 00000000 00010101 00100111 00001010`
- **Ciphertext (ASCII):** `\x03\x00\x15'\n`

### Hamming Distance

Hamming distance counts how many bits are different between two binary strings of the same length.

- **Example:**  
  - String A: `1011101`  
  - String B: `1001001`  
  - Hamming distance: 3 (the bits differ in positions 2, 4, and 5)

XORing two strings and counting the number of 1s in the result gives you the Hamming distance.

### Cryptoanalysis of Repeated Key XOR cipher
Suppose you have the following ciphertext (hex-encoded):

```
0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f
```

You can break the ciphertext using the following steps:

**1. Estimate Key Length**

To guess the key length, split the ciphertext into blocks of different sizes (from 2 to 40 bytes). For each size, compare the first few blocks, count how many bits are different (Hamming distance), average the results, and divide by the block size. The key length with the lowest score is probably correct.

Let's estimate the key length by comparing normalized Hamming distances for key sizes 2 and 3.

**Key Length: 2**

- Block 1: `0b36` (binary: `00001011 00110110`)
- Block 2: `3727` (binary: `00110111 00100111`)
- Block 3: `2a2b` (binary: `00101010 00101011`)
- Block 4: `2e63` (binary: `00101110 01100011`)

| Block Pair | Block 1 (binary) | Block 2 (binary) | XOR Result (binary) | Hamming Distance | Normalized |
|:-----------|:-----------------|:-----------------|:--------------------|:-----------------|:-----------|
| Block 1 vs Block 2 | 00001011 00110110 | 00110111 00100111 | 00111100 00010001 | 6 | 3.0 |
| Block 1 vs Block 3 | 00001011 00110110 | 00101010 00101011 | 00100001 00011101 | 6 | 3.0 |
| Block 1 vs Block 4 | 00001011 00110110 | 00101110 00101100 | 00100101 00001010 | 7 | 3.5 |
| Block 2 vs Block 3 | 00110111 00100111 | 00101010 00101011 | 00011101 00001100 | 6 | 3.0 |
| Block 2 vs Block 4 | 00110111 00100111 | 00101110 00101100 | 00011001 00001011 | 5 | 2.5 |
| Block 3 vs Block 4 | 00101010 00101011 | 00101110 00101100 | 00000100 00000111 | 3 | 1.5 |

Average normalized Hamming distance for key length 2: **2.75**

**Key Length: 3**

- Block 1: `0b3637` (binary: `00001011 00110110 00110111`)
- Block 2: `272a2b` (binary: `00100111 00101010 00101011`)
- Block 3: `2e6362` (binary: `00101110 01100011 01100010`)
- Block 4: `2c2e69` (binary: `00101100 00101110 01101001`)

| Block Pair | Block 1 (binary) | Block 2 (binary) | XOR Result (binary) | Hamming Distance | Normalized |
|:-----------|:-----------------|:-----------------|:--------------------|:-----------------|:-----------|
| Block 1 vs Block 2 | 00001011 00110110 00110111 | 00100111 00101010 00101011 | 00101100 00111000 00011100 | 9 | 3.0 |
| Block 1 vs Block 3 | 00001011 00110110 00110111 | 00101110 01100011 01100010 | 00100101 01000001 01010101 | 11 | 3.67 |
| Block 1 vs Block 4 | 00001011 00110110 00110111 | 00101100 00101110 01101001 | 00100111 00001000 01011110 | 11 | 3.67 |
| Block 2 vs Block 3 | 00100111 00101010 00101011 | 00101110 01100011 01100010 | 00001001 01001000 00001001 | 8 | 2.67 |
| Block 2 vs Block 4 | 00100111 00101010 00101011 | 00101100 00101110 01101001 | 00001011 00011100 00000010 | 6 | 2.0 |
| Block 3 vs Block 4 | 00101110 01100011 01100010 | 00101100 00101110 01101001 | 00000010 01001101 00001011 | 8 | 2.67 |

Average normalized Hamming distance for key length 3: **2.94**

**2. Break into Blocks**

To break the ciphertext into blocks, use the first few key lengths you estimated (like 2 and 3) and split the ciphertext into chunks of those sizes.

Assume the key length is 3 (for this example). Break the ciphertext into consecutive blocks of 3 bytes each:

```
Block 1: 0b 36 37
Block 2: 27 2a 2b
Block 3: 2e 63 62
Block 4: 2c 2e 69
```

**3. Transpose Blocks**

Transpose the blocks so that each new block contains all bytes encrypted with the same key byte.

Transposed blocks for key length 3:

```
Transposed Block 1: 0b 27 2e 2c
Transposed Block 2: 36 2a 63 2e
Transposed Block 3: 37 2b 62 69
```

**4. Single Byte XOR Cipher**

For each transposed block, use single-byte XOR and frequency analysis to find the most likely key byte.  
For example:  
- Transposed Block 1 (`0b 27 2e 2c`): key byte is likely `0x49` (`I`)  
- Transposed Block 2 (`36 2a 63 2e`): key byte is likely `0x43` (`C`)  
- Transposed Block 3 (`37 2b 62 69`): key byte is likely `0x45` (`E`)

**5. Combine Key Bytes**

Combine the key bytes to get the full key: `"ICE"`.

**Decrypt the Ciphertext**

XOR each byte of the ciphertext with the corresponding byte of the repeating key:

```
plaintext = ciphertext_byte ^ key_byte
```

**Resulting Plaintext:**

```
Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal
```
