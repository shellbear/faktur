<p align="center">
  <img alt="Gopher" src=".github/images/gopher.png" height="140" />
  <h3 align="center">Faktur</h3>
  <p align="center">Invoice generator as a service.</p>
</p>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/shellbear/faktur" alt="Go Report Card">
    <img src="https://goreportcard.com/badge/github.com/shellbear/faktur" />
  </a>
  <a href="https://github.com/shellbear/faktur/actions?query=workflow%3Alint" alt="Pipeline status">
    <img src="https://github.com/shellbear/faktur/workflows/lint/badge.svg" />
  </a>
  <img src="https://img.shields.io/github/go-mod/go-version/shellbear/faktur" alt="Go version" />
  <a href="https://opensource.org/licenses/MIT" alt="Go version">
    <img src="https://img.shields.io/badge/license-MIT-brightgreen.svg" />
  </a>
</p>

---

## Get Faktur

With go CLI:
```shell script
go build -o faktur .
./faktur
```

With docker:
```shell script
docker run -p 8080:8080 docker.pkg.github.com/shellbear/faktur/faktur
```

Or download binary from the release [page](https://github.com/shellbear/faktur/releases).

## How it works?

![faktur](.github/images/faktur.png)

## Build with

- [go-wkhtmltopdf](https://github.com/SebastiaanKlippert/go-wkhtmltopdf)
- [basscss](https://github.com/basscss/basscss)

## Credits

- Artwork by [Ashley McNamara](https://twitter.com/ashleymcnamara)
- Inspired by [Renee French](http://reneefrench.blogspot.co.uk/)

## License

© shellbear, 2020~time.Now

Released under the [MIT License](LICENSE).
