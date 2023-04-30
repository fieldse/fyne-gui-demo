## Fyne GUI widget

A simple GUI widget for testing Fyne on Mac M1 silicon

### Requirements

-   Macbook with M1 silicon hardware
-   Go 1.19+ (may work with older versions)
-   Fyne application libraries (see installation instructions below)

### Installation

1. Install fyne

```
go install fyne.io/fyne/v2/cmd/fyne@latest
```

2. Build the project

```
go mod tidy
go build app/main.go

```

**(Optional) Compiling desktop application**

This installs the application with icons etc. in your system applications directory

```
fyne install
```

### Usage

Run the application

```
go run app/main.go
```

### Screenshots

![Screenshot 1](img/screenshots/1.png?raw=true)

![Screenshot 2](img/screenshots/2.png?raw=true)
