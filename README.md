# Cypher Tool

The **Cypher Tool** is a command-line application that allows users to encrypt and decrypt messages using three encryption methods: **ROT13**, **Reverse**, and  **Random Mapping**.

## Overview

The Cypher Tool consists of the following functions:

- `main`: The main entry point of the application.
- `getInput`: Handles user input for selecting operations and ciphers.
- `readInput`: Reads the input from the keyboard
- `encrypt_rot13`: Encrypts messages using the ROT13 cipher and returns encrypted string.
- `decrypt_rot13`: Decrypts messages encrypted with the ROT13 cipher and returns decrypted string.
- `encrypt_reverse`: Encrypts messages by reversing the characters and returns encrypted string.
- `decrypt_reverse`: Decrypts messages  by reversing the characters and returns decrypted string.
- `createEncryptMap`: Stores the map/list for all standard printable ASCI characters needed for encryption
- `createDecryptMap`: Same as `createEncryptMap` but in revers
- `encrypt_mappin`: Encrypt message by using encryptMap list for the characters and returns encrypted string.
- `decrypt_mapping`: Decrypts messages custem maps for the characters and returns decrypted string.

# What the Cypher Tool does

The Cypher Tool allows users to encrypt and decrypt data using several encryption methods. The first two methods, **ROT13** and **Reverse**, are designed specifically for Latin alphabet characters. The third method, **Random Mapping**, can encrypt and decrypt all standard printable ASCII characters.


# Features

- Greets the user.
- Allows the user to select the operation (encrypt or decrypt).
- Allows the user to select the encryption type (ROT13, Reverse, or Mapping).
- Accepts user input for messages to encrypt/decrypt.
- Outputs the result of the encryption or decryption operation.

# How to Use Cypher Tool

1. **Open Command Prompt**  
   Navigate to the directory where the Cypher Tool is located.

    Example:
   ```bash
    cd /path/to/cyphertool
   ```
2. **Run Cypher Tool**  
   Use the following command to launch the tool:

   ```bash
   ./cyphertool
   ```

3. **Select an Operation**  
   After launching, you will be prompted to select an operation:
   - Press `1` and after that `Enter` to **Encrypt** a message.
   - Press `2` and after that `Enter` to **Decrypt** a message.

   Example input:
    ```bash
    Select operation 1 or 2 and press "Enter":
            1. Encrypt.
            2. Decrypt.
    Enter your input: 1
    ```


4. **Choose a Cypher Method**  
   Once you select an operation, choose the desired encryption or decryption method:
   - Press `1` and after that `Enter` for **ROT13**.
   - Press `2` and after that `Enter` for **Reverse** (reverses the string).
   - Press `3` and after that `Enter` for **Random Mapping** (maps characters randomly).

   Example input:
    ```bash
    Select cypher and press "Enter":
            1. ROT13.
            2. Reverse.
            3. Random mapping.
    Enter your input: 1
    ```
5. **Enter the Message**  
     You will then be asked to enter the message you want to encrypt or decrypt. Type your message and press `Enter`.

   Example input:
    ```bash
    Enter the message and press "Enter":
    This is a secret message
    ```
6. **View the Result**  
   After processing, the tool will display the encrypted or decrypted message.

    Example result (for ROT13):
    ```bash
    Encrypted message using ROT13:
    Guvf vf n frperg zrffntr
    ```
7. **Continue or Exit**  
   The tool will prompt you if you want to continue using it.
    - Type `Y` and press `Enter` to continue.
    - Press any other key and `Enter` to exit.
   ```bash
    If you want to continue using the tool choose 'Y' and press "Enter":
   ```


## The Cyphers Used:

### ROT13
    ROT13 encryption is a simple letter substitution cipher that shifts each letter in the alphabet by 13 positions. For example, 'A' becomes 'N', 'B' becomes 'O', and so on. After shifting by 13, the letters wrap around, so 'N' becomes 'A'. It is commonly used for basic obfuscation, but it is not a secure encryption method.

### Reverse
    The Reverse method is a cipher where each letter is replaced by its opposite in the alphabet. For example, 'A' becomes 'Z', 'B' becomes 'Y', and so on. It swaps the position of every letter with its mirror image in the alphabet.

### Random Mapping
    Our custom encryption Random Mapping is a cipher where each letter, number, and printable symbol is replaced with a random corresponding symbol based on a predefined mapping. 
    For example:
    "a" becomes "B", "b" becomes "f", "d" becomes "m"

    Each letter has a specific random match, making it less predictable than simple substitution ciphers.
