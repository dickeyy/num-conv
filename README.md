# num-conv

A simple little utility to convert any decimal, hex, or binary number to the other representations.

I made this for a class where we regularly had to convert numbers between different bases, and I wanted a quick way to check my work.

## Usage

#### Clone source

```bash
git clone https://github.com/dickeyy/num-conv
cd num-conv
go run main.go <number>
```

#### Install binary

```bash
go install github.com/dickeyy/num-conv@latest
num-conv <number>
```

For `number`, you can use decimal (positive, negative, or fractional), hex, or binary.

```bash
num-conv 0xAFB
num-conv 0b10010
num-conv 12
num-conv -12
num-conv 12.34
num-conv -12.34
```

## License

Licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
