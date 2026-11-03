# 🎨 ASCII Art Web

> A simple and interactive web application that transforms your text into **ASCII Art** using different banners and colors.

---

## ✨ Description

**ASCII Art Web** is a web application built with **Go** that provides a graphical interface for generating ASCII Art.

Instead of using the command line, users can interact with the application directly from their browser.

### 🚀 Features

* 📝 Enter custom text
* 🎨 Choose a text color
* 🔤 Choose between different ASCII Art banners
* 🖥️ Generate ASCII Art directly in the browser
* 📖 About page
* 🕐 History page
* 📩 Contact page
* 📱 Simple and user-friendly interface



---

## 🛠️ Technologies

| Technology         | Usage                 |
| ------------------ | --------------------- |
| 🐹 Go              | Backend & HTTP Server |
| 🌐 HTML            | Web pages             |
| 🎨 CSS             | Styling               |
| 📦 `html/template` | HTML Templates        |
| 🌍 `net/http`      | Web Server            |
| 🔤 ASCII Art       | Text Rendering        |

---

## ▶️ Usage

### 1. Clone the repository

```bash
git clone <repository-url>
cd ascii-art-web
```

### 2. Run the server

```bash
go run .
```

When the server starts, you should see:

```text
server started on :8080
```

### 3. Open the application

Open your browser and visit:

**http://localhost:8080**

---

## 🧩 Project Structure

```text
ascii-art-web/
│
├── main.go
│
├── ascii/
│   ├── ...
│   └── ...
│
├── handler/
│   └── handler.go
│
├── templates/
│   ├── layout.html
│   ├── index.html
│   ├── about.html
│   ├── history.html
│   └── contact.html
│
├── static/
│   ├── css/
│   └── ...
│
└── README.md
```

---

# ⚙️ Implementation Details

## 🌐 HTTP Server

The application uses Go's standard `net/http` package to create the web server.

The server listens on:

```text
:8080
```

The following routes are registered:

| Route      | Method | Description        |
| ---------- | ------ | ------------------ |
| `/`        | GET    | Home page          |
| `/ascii`   | POST   | Generate ASCII Art |
| `/about`   | GET    | About page         |
| `/history` | GET    | History page       |
| `/contact` | GET    | Contact page       |
| `/static/` | GET    | Static files       |

---

## 🔄 How It Works

The application follows this flow:

```text
        User
          │
          ▼
     Enter Text
          │
          ▼
    Select Font
          │
          ▼
    Select Color
          │
          ▼
      Submit
          │
          ▼
   ┌───────────────┐
   │  Go Handler   │
   └───────┬───────┘
           │
           ▼
    Load ASCII Banner
           │
           ▼
    Render ASCII Art
           │
           ▼
      Apply Color
           │
           ▼
    Render HTML Page
           │
           ▼
        Browser
```

---

# 🧠 Algorithm

When the user submits the form, the application performs the following steps:

### 1️⃣ Get the user input

The handler receives:

* `text`
* `font`
* `color`

```go
text := r.FormValue("text")
font := r.FormValue("font")
color := r.FormValue("color")
```

---

### 2️⃣ Validate the input

If the text is empty, the application returns to the home page.

If no font is selected, the default font is:

```text
standard
```

---

### 3️⃣ Load the banner

The selected ASCII Art banner is loaded using:

```go
banner, err := ascii.LoadBanner(font)
```

If the banner does not exist, the server returns:

```text
404 Not Found
```

---

### 4️⃣ Generate the ASCII Art

The text and banner are passed to:

```go
result, err := ascii.RenderAscii(text, banner)
```

The renderer converts each character of the input into its corresponding ASCII Art representation.

The characters are combined line by line to create the final result.

For example:

```text
Input:

HELLO
```

becomes a multi-line ASCII representation based on the selected banner.

---

### 5️⃣ Apply the color

If the user chooses a color, the generated ASCII Art is wrapped inside an HTML `<span>`:

```html
<span style="color: ...">
    ASCII ART
</span>
```

This allows the result to be displayed using the selected color.

---

### 6️⃣ Render the result

The generated ASCII Art is passed to the template through `PageData`:

```go
type PageData struct {
    Result template.HTML
    Text   string
    Font   string
    Color  string
}
```

The template then displays the final result to the user.

---

# 🧱 Main Components

### `main.go`

Responsible for:

* Starting the HTTP server
* Registering routes
* Serving static files
* Connecting URLs to their handlers

---

### `handler/handler.go`

Responsible for:

* Handling HTTP requests
* Validating requests
* Reading form data
* Loading banners
* Calling the ASCII renderer
* Rendering HTML templates
* Handling errors

---

### `ascii/`

Contains the ASCII Art logic:

* Loading banners
* Validating banners
* Converting text into ASCII Art

---

### `templates/`

Contains the HTML pages and shared layout.

The application uses Go's:

```go
html/template
```

package to render the pages.

---

### `static/`

Contains static resources such as:

* CSS
* Images
* Other frontend files

---

# 🚨 Error Handling

The application handles different errors and returns appropriate HTTP status codes.

| Error                   |                      Status |
| ----------------------- | --------------------------: |
| Invalid request         |           `400 Bad Request` |
| Font not found          |             `404 Not Found` |
| Wrong HTTP method       |    `405 Method Not Allowed` |
| Template/internal error | `500 Internal Server Error` |

---

# 🎯 Goal

The goal of this project is to transform the previous **ASCII Art** command-line project into an interactive **web application**, while practicing:

* Go HTTP servers
* Routing
* HTTP methods
* HTML templates
* Form handling
* Error handling
* Frontend/backend integration
* ASCII Art rendering

---

## 💻 Run It

```bash
go run .
```

Then open:

```text
http://localhost:8080
```

---

## 👩‍💻 Authors

**Mostafa Mahmoud**
**Enas Elashrey**

### ⭐ Enjoy creating ASCII Art!
