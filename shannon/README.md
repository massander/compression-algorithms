# Shannon-Fano coding algorithm

Link: [Shanon Fano coding](https://en.wikipedia.org/wiki/Shannon%E2%80%93Fano_coding)

```Go
    package main

    import "github.com/2thousandmax/CompressionAlgorithms/shannon"

    func main() {
        testString := "shannon fano algorithm"
        shanonFanoCode := NewShannonFanoCode(testString)

        fmt.Println(shannonFanoCode)
    }
```

## Example

![Shanon-Fano example](../assets/ShannonFano.png "Shannon-Fano example")
<!-- ![Shanon-Fano example](https://github.com/2thousandmax/algorithms/blob/main/assets/ShannonFano.png?raw=true) -->
