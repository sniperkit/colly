# Gowap [[Wappalyzer](https://github.com/AliasIO/Wappalyzer) implementation in Go]

## Notes

* Current implementation does not support analyzing JS variables, cause I could not find easy way to get global js variables from Window object in GO, anyone welcome to contribute!

* Some features was not implemented, such as confidence rate, as I do not see need for it. I also did not implement recursive crawling, but it is pretty easy to do.
* If you want to participate in development, fork it, make your advancements and create merge request, everyone is welcome!

## Usage

`go get github.com/altsab/gowap`

Call `Init()` function from package providing path to `apps.json` file and boolean value to choose between JSON (`true`) and raw output, it will return `Wappalyzer` object on which you can call Analyze method with URL string as argument.

```golang

func Init(appsJSONPath string, JSON bool) (wapp *Wappalyzer, err error)
func (wapp *Wappalyzer) Analyze(url string) (result string, err error)

```