<div align="center">
  <img src="https://github.com/saumsy/localizer/assets/64534496/af788aac-9bef-4f14-ba64-17608789162c" width="300" alt="localizer">
  <h1 align="center">Localizer</h1>
  Quickly and easily localize your Xcode application
</div>

***

<div align="left">

## Overview
Localizer helps you localize (add internationalisation) within your Xcode application within seconds, so you just have to focus on the code!
As it is built with Go, so speed isn't a compromise 🚀.

![localizer-demo](https://github.com/saumsy/localizer/assets/64534496/be26ea72-3825-426e-9cfd-3b879193c2b7)

## Installation

### Using Go
```bash
go install github.com/saumsy/localizer@latest
```

### Using Github Releases
Navigate to [releases](https://github.com/saumsy/localizer/releases/latest), find the suitable binary for your OS and install to get started

## Usage
```bash
localizer help
```
- You need to have a base `.lproj/` directory (for example:- `en.lproj`) and a `Localizable.strings` file containing the strings that need be to localized and thats it
- Separate the languages with a `space` you want to localize.

### Example Usage
```bash
localizer localize en.lproj fr de es
``` 

</div>
