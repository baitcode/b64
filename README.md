# About

Cli tool written with Golang to perform b64 encryption and decryption. Nothing special, just a tool that is supposed to work the same on Linux, Max and Windows for the sake of writing simpler instructions for the software I make. Because, you know. You can't just use base64 encoding in shell in windows.

Feel free to open issues or any other way of initiating conversations

# Usage

```bash
go install github.com/baitcode/b64
```

## Encode:

```bash
b64 encode "some data"
>> c29tZSBkYXRh
```

Pipes are also supported

```bash
echo "some data" | b64 encode
>> c29tZSBkYXRh
```

## Decode:

```bash
b64 decode "c29tZSBkYXRh"
>> some data
```

Pipes are also supported

```bash
echo "c29tZSBkYXRh" | b64 decode
>> some data
```
