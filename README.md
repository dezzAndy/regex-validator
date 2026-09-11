# Validador de Expresiones Regulares con GUI

El sistema implementa un conjunto de reconocedores léxicos para validar formatos estructurados bajo el estándar y contexto mexicano. El núcleo de reconocimiento está desarrollado en **Go** y se compila como una biblioteca compartida nativa (`.so` / `.dll`), la cual es consumida mediante **P/Invoke** por una interfaz gráfica de escritorio desarrollada en **C#** con **Avalonia UI**.

---

## Características y Formatos Soportados

El analizador valida 8 tipos de cadenas:

| Tipo | Patrón / Regla | Ejemplo válido |
|---|---|---|
| **Teléfono** | 10 dígitos numéricos unificados | `3312345678` |
| **Correo Electrónico** | RFC 5322 simplificado (`usuario@dominio.tld`) | `usuario@dominio.com` |
| **CURP** | 18 caracteres alfanuméricos con validación posicional de entidad, sexo y consonantes | `GARC820714HJCRRL09` |
| **Contraseña Segura** | Mínimo 8 caracteres, al menos una mayúscula, un número y un símbolo especial | `Hola123!` |
| **RFC** | Persona física (13 caracteres) o persona moral (12 caracteres, admite `&` y `Ñ`) | `GARC820714H10` / `XYZ890101ABC` |
| **Dirección IP** | IPv4 (octetos 0–255) e IPv6 (soporte de compresión `::` y zonas de interfaz) | `192.168.1.1` / `2001:db8::1` |
| **Cumpleaños** | Formato `dd/mm/aaaa` con validación real de calendario (años bisiestos) | `29/02/2000` |
| **Placas Vehiculares** | Formato federal estándar para automóviles particulares (`AAA-1234`) | `ABC-1234` |

---

## Arquitectura

```
regex-validator/
├── lib/                                   # Bibliotecas nativas compiladas
│   ├── libvalidador.so                    # Biblioteca compartida para Linux
│   └── validador.dll                      # Biblioteca dinámica para Windows
├── scripts/                               # Scripts de automatización de compilación
│   ├── build-linux.sh
│   ├── build-windows.sh
│   └── build-windows.bat
├── src/
│   ├── core/                              # Módulo en Go (Reconocedores y CGo)
│   │   ├── automatas/                     # Implementación de autómatas y pruebas
│   │   ├── go.mod
│   │   └── main.go                        # Funciones exportadas con CGo
│   └── gui/
│       └── ValidadorGUI/                  # Aplicación de escritorio en C# (Avalonia UI)
│           ├── MainWindow.axaml
│           ├── MainWindow.axaml.cs
│           ├── NativeValidator.cs         # Interfaz P/Invoke (DllImport)
│           └── ValidadorGUI.csproj
└── README.md
```

---

## Requisitos Previos

Asegúrate de contar con las siguientes herramientas instaladas en tu sistema:

* **Go** (versión 1.20 o superior).
* **.NET SDK** (versión 8.0, 9.0 o 10.0).
* **Compilador C (GCC)** con soporte para CGo:
  * En **Arch Linux / CachyOS / Manjaro**: `sudo pacman -S base-devel`
  * En **Ubuntu / Debian**: `sudo apt install build-essential`
  * En **Fedora**: `sudo dnf groupinstall "Development Tools"`
* *(Opcional, para compilar la DLL de Windows desde Linux)*:
  * `sudo pacman -S mingw-w64-gcc` (Arch)
  * `sudo apt install gcc-mingw-w64` (Ubuntu/Debian)

---

## Guía de Instalación y Ejecución

### 1. Clonar el repositorio

```bash
git clone https://github.com/dezzAndy/regex-validator.git
cd regex-validator
```

### 2. Compilar la biblioteca nativa en Go

#### En Linux (`.so`):
```bash
./scripts/build-linux.sh
```
*Esto generará el archivo `lib/libvalidador.so` y su cabecera `lib/libvalidador.h`.*

#### En Windows o Cross-compilación (`.dll`):
```bash
# Desde Linux con mingw-w64:
./scripts/build-windows.sh

# O desde Windows (CMD / PowerShell):
scripts\build-windows.bat
```

### 3. Ejecutar las pruebas unitarias del núcleo (Go)

Para verificar que todos los autómatas y expresiones regulares pasen su suite de pruebas:

```bash
cd src/core
go test -v ./...
cd ../..
```

### 4. Ejecutar la interfaz gráfica (C# / Avalonia)

El proyecto de C# está configurado para copiar automáticamente la biblioteca nativa correspondiente (`.so` en Linux o `.dll` en Windows) a su carpeta de salida binaria:

```bash
dotnet run --project src/gui/ValidadorGUI
```

---

## Tecnologías Utilizadas

* **[Go](https://go.dev/):** Implementación de expresiones regulares (`regexp`/RE2), validación procedural de estados y exportación de ABI estándar de C mediante CGo.
* **[C# / .NET](https://dotnet.microsoft.com/):** Capa de aplicación e invocación de funciones no administradas mediante `System.Runtime.InteropServices` (`DllImport`).
* **[Avalonia UI](https://avaloniaui.net/):** Framework de interfaz gráfica XAML multiplataforma para escritorio (Linux, Windows y macOS).

---

## Licencia

Proyecto con fines académicos.
