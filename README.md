# Cypher Tool

The **Cypher Tool** is a command-line application that allows users to encrypt and decrypt messages using three encryption methods: **ROT13**, **Reverse**, and  **Random Mapping**.

## Overview

The Cypher Tool consists of the following functions:

- `main()`: The main entry point of the application.
- `getInput() (toEncrypt bool, encoding string, message string)`: Handles user input for selecting operations and ciphers.
- `readInput() string `: Reads the input from the keyboard
- `encrypt_rot13(s string) string`: Encrypts messages using the ROT13 cipher and returns encrypted string.
- `decrypt_rot13(s string) string`: Decrypts messages encrypted with the ROT13 cipher and returns decrypted string.
- `encrypt_reverse(s string) string`: Encrypts messages by reversing the characters and returns encrypted string.
- `decrypt_reverse(s string) string`: Decrypts messages  by reversing the characters and returns decrypted string.
- `createEncryptMap() map[rune]rune`: Stores the map/list for all standard printable ASCI characters needed for encryption
- `createDecryptMap() map[rune]rune`: Same as `createEncryptMap` but in revers
- `encrypt_mappin(plainText string) string`: Encrypt message by using encryptMap list for the characters and returns encrypted string.
- `decrypt_mapping(encryptedText string) string `: Decrypts messages custem maps for the characters and returns decrypted string.


## What the Cypher Tool does

The Cypher Tool allows users to encrypt and decrypt data using several encryption methods. The first two methods, **ROT13** and **Reverse**, are designed specifically for Latin alphabet characters. The third method, **Random Mapping**, can encrypt and decrypt all standard printable ASCII characters.


## Features

- Greets the user.
- Allows the user to select the operation (encrypt or decrypt).
- Allows the user to select the encryption type (ROT13, Reverse, or Mapping).
- Accepts user input for messages to encrypt/decrypt.
- Outputs the result of the encryption or decryption operation.

## Usage

To use the Cypher Tool, follow these steps:

Launch the program:
```bash
$ ./cyphertool
# "Welcome to the Cypher Tool 2024"

## Select Operation (1/2/3):
- 1. Encrypt
- 2. Decrypt

```bash
$ 1
# Cypher Tool Instructions

## Select Cypher (1/2/3):
- 1. ROT13
- 2. Reverse
- 3. Random Mapping

```bash
$ 3
Enter the message + Enter:
$ This is a secret message.

Decrypted message using reverse:
$ Gdre re M eqSCqf jqeeMNq.

## Enter the "Y" or "y" + Enter if you want to continue:

```bash
$ y

## The Cyphers Used:

### ROT13
ROT13 (rotate by 13 places) replaces a letter with the letter 13 letters after it in the alphabet.

### Reverse
The Reverse cypher takes a letter as input and returns its reverse letter in the Latin alphabet.

### Random Mapping
Our custom encryption takes a letter as input, returns coresponding symbols from the map list letter in the Latin alphabet.
