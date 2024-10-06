# Cypher Tool

The **Cypher Tool** is a command-line application that allows users to encrypt and decrypt messages using three encryption methods: **ROT13**, **Reverse**, and  **Random Mapping**.

## Overview

The Cypher Tool consists of the following functions:

- `main`: The main entry point of the application.
- `getInput`: Handles user input for selecting operations and ciphers.
- `encrypt_rot13`: Encrypts messages using the ROT13 cipher.
- `decrypt_rot13`: Decrypts messages encrypted with the ROT13 cipher.
- `encrypt_reverse`: Encrypts messages by reversing the characters.
- `decrypt_reverse`: Decrypts messages  by reversing the characters.
- `encrypt_ mapping`:Encrypts messages custem maps for the characters.
- `decrypt_reverse`:Decrypts messages custem maps for the characters.

## What the Cypher Tool does

- Greets the user
- Allows the user to select the operation (encrypt or decrypt)
- Allows the user to select the encryption type
- Allows the user to input the message to encrypt/decrypt
- Outputs the result of the operation
- Offer contue choose 


## Features

- Greets the user.
- Allows the user to select the operation (encrypt, decrypt).
- Allows the user to select the encryption type (ROT13, Reverse, or Mapping).
- Accepts user input for messages to encrypt/decrypt.
- Outputs the result of the encryption or decryption operation.

## Usage

To use the Cypher Tool, follow these steps:

Launch the program:
```bash
$ ./cyphertool
# Welcome to the Cypher Tool!

## Select Operation (1/2/3):
- **1.** Encrypt
- **2.** Decrypt

```bash
$ 1
# Cypher Tool Instructions

## Select Cypher (1/2/3):
- **1.** ROT13
- **2.** Reverse
- **3.** Random Mapping

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
