<p align="center">
  <img alt="Gopher" src=".github/images/gopher.png" height="140" />
  <h3 align="center">Faktur</h3>
  <p align="center">Invoice generator as a service.</p>
</p>

---

[![Go Report Card](https://goreportcard.com/badge/github.com/shellbear/faktur)](https://goreportcard.com/report/github.com/shellbear/faktur)
[![Pipeline status](https://github.com/shellbear/faktur/workflows/lint/badge.svg)](https://github.com/shellbear/faktur/actions)
![Go version](https://img.shields.io/github/go-mod/go-version/shellbear/faktur)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

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
