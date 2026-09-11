# 🎨 ASCII Art Web Generator

<p align="center">

# ✨ Turn Text Into Beautiful ASCII Art ✨

A modern **Go Web Application** that transforms supported text into stunning **ASCII Art** using multiple fonts, customizable colors, and a clean responsive interface.

<img src="static/home.png" width="900"/>

</p>

---

## 🌟 Features

✨ Convert supported text into ASCII Art instantly

🎨 Multiple banner fonts

* Standard
* Shadow
* Thinkertoy

🌈 Color customization

📝 Multi-line text support

⚡ Fast Go backend

📱 Responsive Design

📚 Additional Pages

* 🏠 Home
* 📖 History
* 💡 About
* 📩 Contact

🛡️ Built-in Error Handling

---

# 🛠️ Built With

<p>

🟦 **Go (Golang)**

🌐 **HTML5**

🎨 **CSS3**

📄 **Go Templates**

⚙️ **net/http**

</p>

---

# 📂 Project Structure

```text
ascii-art-web
│
├── ascii
│   ├── fonts
│   │   ├── standard.txt
│   │   ├── shadow.txt
│   │   └── thinkertoy.txt
│   ├── renderer.go
│   └── renderer_test.go
│
├── handler
│   ├── handler.go
│   └── handler_test.go
│
├── templates
│   ├── layout.html
│   ├── index.html
│   ├── about.html
│   ├── history.html
│   ├── contact.html
│   └── error.html
│
├── static
│   ├── style.css
│   ├── image.png
│   └── home.png
│
├── main.go
├── go.mod
└── README.md
```

---

# ⚙️ How It Works

```text
              🌍 Browser
                   │
                   ▼
            HTTP Request
                   │
                   ▼
             Go HTTP Server
                   │
                   ▼
              HTTP Router
                   │
           ┌───────┼─────────┐
           ▼       ▼         ▼
         Pages  ASCII Art  Static
                   │
                   ▼
               Read Form
                   │
                   ▼
            Validate Input
                   │
                   ▼
             Load Banner
                   │
                   ▼
           Generate ASCII
                   │
                   ▼
           Render Template
                   │
                   ▼
              🌍 Browser
```

---

# 🧠 Implementation

The application follows a simple flow:

1. The browser sends the entered text, selected banner, and color to `/ascii-art`.
2. The handler validates the HTTP method, and selected banner.
3. The selected banner file is loaded from `ascii/fonts/`.
4. The banner data is validated and prepared for rendering.
5. Each supported character is mapped to its 8-line ASCII representation.
6. Multi-line input is processed line by line.
7. The generated ASCII result is passed to the HTML template.
8. Go's `html/template` package renders the final page safely.
9. Errors such as invalid routes, methods, banners, characters, or template failures are handled with the appropriate HTTP response.

---

# 🌐 Available Routes

| Route | Method | Description |
| :--- | :---: | --- |
| `/` | GET | 🏠 Home Page |
| `/ascii-art` | POST | 🎨 Generate ASCII Art |
| `/about` | GET | 💡 About Project |
| `/history` | GET | 📖 ASCII History |
| `/contact` | GET | 📩 Contact Page |
| `/static/*` | GET | 📁 Static Files |

---

# 🎯 Supported Fonts

| Font | Preview |
| --- | --- |
| Standard | ⭐⭐⭐⭐⭐ |
| Shadow | ⭐⭐⭐⭐⭐ |
| Thinkertoy | ⭐⭐⭐⭐⭐ |

---

# ❗ Error Handling

✅ Invalid HTTP Method

✅ Invalid Route

✅ Missing Banner File

✅ Invalid Banner

✅ Unsupported Characters

✅ Input Length Validation

✅ Template Rendering Errors

---

# 🚀 Getting Started

### 1️⃣ Clone the repository

```bash
git clone https://01.nextera.education/git/momahmoud/ascii-art-web
```

### 2️⃣ Navigate to the project

```bash
cd ascii-art-web
```

### 3️⃣ Run the application

```bash
go run .
```

### 4️⃣ Open your browser

```text
http://localhost:8080
```

---

# 🧪 Testing

Run all tests with:

```bash
go test ./...
```

The test suite covers:

* ASCII rendering behavior
* Banner loading and validation
* Multi-line input
* Unsupported characters
* HTTP status handling
* Invalid fonts
* Input length limits
* Template rendering failures
* Official ASCII Art audit cases

---

# 💻 Example

Input:

```text
Hello
```

↓

Output:

```text
 _    _          _ _
| |  | |        | | |
| |__| |   ___  | | |  ___
|  __  |  / _ \\ | | | / _ \\
| |  | | |  __/ | | || (_) |
|_|  |_|  \\___| |_|_| \\___/
```

---

# 🚧 Future Improvements

* 📋 Copy ASCII with one click
* 📥 Download as TXT
* 🌙 Dark / Light Theme
* 🔥 More Fonts
* ⚡ Live Preview
* 🕘 Save Previous Generations

---

### ❤️ Developed by

**Enas Essam**

**Mostafa Mahmoud**

---

<p align="center">

### ✨ Made with Go ❤️ and lots of Coffee ☕

</p>
