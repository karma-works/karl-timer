# Karl Timer

Eine einfache, elegante Stoppuhr-Applikation in Go-Lang, die für maximale Lesbarkeit und einfache Bedienung im Vollbildmodus entwickelt wurde.

## Features

- **Riesige Anzeige**: Die Zeitzahlen sind extra groß für beste Sichtbarkeit aus der Ferne.
- **Vollbild-Modus**: Startet automatisch im Vollbild. Mit `ESC` kann der Modus gewechselt werden.
- **Einfache Steuerung**:
  - **Linksklick**: Startet oder stoppt den Timer.
  - **F5**: Setzt den Timer auf Null zurück.
- **Dark Mode**: Ein kontrastreicher, dunkler Hintergrund schont die Augen.
- **Multi-Monitor Support**: Startet auf dem zuletzt verwendeten Bildschirm.

## Voraussetzungen

- [Go](https://golang.org/dl/) (getestet mit 1.20+)
- C-Compiler (für Fyne/GLFW, z.B. GCC oder Clang)
- Grafiktreiber mit OpenGL-Unterstützung

## Installation & Ausführung

### 1. Abhängigkeiten installieren
Stellen Sie sicher, dass alle Go-Abhängigkeiten geladen sind:
```bash
go mod tidy
```

### 2. Anwendung starten
Um die Applikation direkt aus dem Quellcode zu starten:
```bash
go run main.go
```

### 3. Kompilieren
Um eine ausführbare Datei zu erstellen:
```bash
go build -o karl-timer
./karl-timer
```

## Fehlerbehebung (macOS)
Falls Fehlermeldungen bezüglich fehlender Header erscheinen, stellen Sie sicher, dass die Xcode Command Line Tools installiert sind:
```bash
xcode-select --install
```
