# hashvalue

A program to calculate a hash value from input or file contents.

[![Go Report Card](https://goreportcard.com/badge/github.com/xformerfhs/hashvalue)](https://goreportcard.com/report/github.com/xformerfhs/hashvalue)
[![License](https://img.shields.io/github/license/xformerfhs/hashvalue)](https://github.com/xformerfhs/hashvalue/blob/main/LICENSE)

## Introduction

Sometimes one needs the hash value of a text or of the contents of a file.
There are several utilities that can calculate the value of a single hash algorithm.
The output is mostly just hex bytes.

This utility has been created in order to be able to calculate the values of several hash algorithms with a single program and to get the output in one of several formats.

## Call

The program is called like this:

```
hashvalue [--hash <algorithm>] {--source <text> | --file <path>} [--hex] [--base16] [--base32] [--base64] [--prefix <text>] [--separator <text>] [--lower | --upper]
```

The options have the following meaning:

| Option      | Meaning                                                                                                         |
|-------------|-----------------------------------------------------------------------------------------------------------------|
| `hash`      | Name of the hash algoritm.                                                                                      |
| `source`    | Text that is to be hashed (Mutually exclusive with `file`).                                                     |
| `file`      | File path of a file whose content is to be hashed (mutually exclusive with `source`).                           |
| `hex`       | Hash value is encoded as a [hexadecimal](https://en.wikipedia.org/wiki/Hexadecimal) (base16) string  (default). |
| `base16`    | Alias for `hex`.                                                                                                |
| `base32`    | Hash value is encoded as a [base32](https://en.wikipedia.org/wiki/Base32) string.                               |
| `base64`    | Hash value is encoded as a [base64](https://en.wikipedia.org/wiki/Base64) string.                               |
| `z85`       | Hash value is encoded as a [Z85](https://rfc.zeromq.org/spec/32) string.                                        |
| `prefix`    | Prefix text for hex encoded bytes.  Only used for `hex` encoding.                                               |
| `separator` | Separator text for hex encoded bytes. Only used for `hex` encoding.                                             |
| `lower`     | Hexadecimal values `A`-`F` are printed in lower case. Only used for `hex` encoding.                             |
| `upper`     | Hexadecimal values `A`-`F` are printed in upper case (default). Only used for `hex` encoding.                   |
| `version`   | Print the version information and exit.                                                                         |

The options can be started with either `--` or `-`.

Specify only one encoding.
If there is more than one encoding specified an error message is printed.

The hash algorithm names consist up to three parts:

1. Algorithm
2. Calculated hash size in bits
3. Output hash size in bits

The hash size is separated by a minus character (`-`) from the algorithm.
When an algorithm has only one hash size (`md5`, `sha1`), the hash size is not specified.
When the output hash is only a part of the calculated hash, this is separated by an underscore character (`_`).

| Algorithm | Meaning                                                                                                                                                                                                   |
|-----------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `blake2b` | A [family](https://en.wikipedia.org/wiki/BLAKE_(hash_function)#BLAKE2) of hash functions that has been designed as a faster alternative to the `SHA2` and `SHA3` hashes, optimized for 64 bit processors. |
| `blake2s` | A [family](https://en.wikipedia.org/wiki/BLAKE_(hash_function)#BLAKE2) of hash functions that has been designed as a faster alternative to the `SHA2` and `SHA3` hashes, optimized for 32 bit processors. |
| `md5`     | One of the first [message digests](https://en.wikipedia.org/wiki/MD5) with a fixed hash size of 128 bits. It is no longer considered secure.                                                              |
| `sha1`    | [Secure Hash Algorithm 1](https://en.wikipedia.org/wiki/SHA-1) with a fixed hash size of 160 bits. It is no longer considered secure.                                                                     |
| `sha2`    | [Secure Hash Algorithm 2](https://en.wikipedia.org/wiki/SHA-2) is a family of hash functions that has been designed as the successor of `SHA-1`.                                                          |
| `sha3`    | [Secure Hash Algorithm 3](https://en.wikipedia.org/wiki/SHA-3) is a family of hash functions that has been designed as the successor of `SHA-2`.                                                          |

The list of supported hash algorithms is as follows:

- `blake2b-256`
- `blake2b-384`
- `blake2b-512`
- `blake2s-256`
- `md5`
- `sha1`
- `sha2-224`
- `sha2-256`
- `sha2-384`
- `sha2-512`
- `sha2-512_224`
- `sha2-512_256`
- `sha3-224`
- `sha3-256`
- `sha3-384`
- `sha3-512`

If the program is called without arguments or with wrong arguments a usage text is printed.

### Examples

In the first example a simple text is hashed:

```
hashvalue --source "There should be a meaning."
```

This prints the following output:

```
9AFA63F5C5BE4BEFEC3D1499470F255EDEDCE0B02B916564111886DCD72597CD
```

If the algorithm is not specified the default is `sha3-256`.
The hash value is printed in hex encoding with upper case letters, since this is the default if no encoding is specified.

Now an example with another output encoding:

```
hashvalue --source "There should be a meaning." --base32
```

This prints the following output:

```
TL5GH5OFXZF673B5CSMUODZFL3PNZYFQFOIWKZARDCDNZVZFS7GQ
```

The hash value is the same as before but now it is printed in base32 encoding.

Now a different hash algorithm is specified:

```
hashvalue --source "There should be a meaning." --z85 --hash blake2b-384
```

```
Xov:#>ia!Rpp[?j9%J]O]&8o8q:EM-h9PwW}RH88&XoG=jNnzgT0<E<7l8$2
```

The hexadecimal out can be modified, so that it can be incorporated in a program source code.
E.g. if one wants the hash formatted for use in [Go](https://go.dev/), or [Java](https://www.java.com/), this could be specified like this:

```
hashvalue --source "There should be a meaning." --hash sha2-256 --hex --prefix 0x --separator ", "  --lower
```

This prints the following output:

```
0x23, 0x85, 0xad, 0x7d, 0x11, 0xb1, 0xd8, 0x33, 0x93, 0x4f, 0x9e, 0x29, 0x4b, 0x4a, 0x8b, 0xeb, 0x37, 0x2b, 0xc1, 0x92, 0x40, 0x3f, 0x4e, 0xe4, 0x70, 0xc7, 0x1d, 0xe1, 0xbf, 0x81, 0xf3, 0xab
```

And, last, but not least, one can calculate the hash value of a file:

```
hashvalue -file main.go --hash sha3-384 --base32
```

This prints the following output:

```
4KRKAQP4UHQTKMU7T6BZLDIPN4TULQ3WOKR5SCTOIBLRTQH5YRJELPC62FWHI4AEHTGAHIGXEOA7I
```

### Return codes

The possible return codes are the following:

| Code | Meaning                   |
|------|---------------------------|
| `0`  | Successful processing     |
| `1`  | Error in the command line |
| `2`  | Error while processing    |

## Program build

You must have Go installed to create the program.
This creates a directory that is specified in the `GOPATH` environment variable.
Under Windows, this is the home directory, e.g. `D:\Users\username\go`.
Under Linux this is `${HOME}/go`.
In that directory there is a subdirectory `src`.

To create the program, the source code must be stored under `%GOPATH%\src\hashvalue` or `${HOME}/go/src/hashvalue`.
Then one has to start the batch file `gb.bat` or the shell script `gb`, which builds the executables.
These scripts expect the UPX program to be in a specific location.
This location can be adapted to the local path.
If UPX is not available, no compression is performed.

As a result, the files `hashvalue` for Linux and `hashvalue.exe` for Windows are created.

## Contact

Frank Schwab ([Mail](mailto:github.sfdhi@slmails.com "Mail"))

## License

This source code is published under the [Apache License V2](https://www.apache.org/licenses/LICENSE-2.0.txt).
