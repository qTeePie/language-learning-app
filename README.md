# 🧠 Language Learning App

> ⚠️ This project is a work in progress — not fully functional yet, but development is in motion! Stay tuned. 💙

> Learn. Quiz. Repeat. 💫 — a mobile-first language learning experience built with love, React Native, and a spicy Go backend.

---

## 📚 Dictionary & Translation Features (In Progress)

This app uses a combo of free APIs and dictionaries to power translation, verb conjugation, and word sense disambiguation — all the good nerdy stuff behind the scenes 😏

### 🧬 Translation Flow:
- Source: **Dictionary Password (English)** ➜ Target language
- Translations are fetched via **Lexica API**
- Outputs are filtered through **Lexica Multigloss Dictionary** to ensure correct *part of speech (POS)* and limit meanings to the most frequent ones 🔍

### 🔄 Verb Conjugation:
We’re exploring a few cool tools for dynamic verb conjugation:
- [`api.verbix.com`](https://api.verbix.com/conjugator/html)
- [`mlconjug3`](https://pypi.org/project/mlconjug3/) — Python package for multilingual conjugation
- [`apertium-apy`](https://github.com/apertium/apertium-apy) — Open-source machine translation API

---

## ⚙️ Backend (Go + MongoDB)

The backend is written in **Go (Golang)** using the `gin` framework and stores user data with MongoDB. It also powers translation lookups and manages the learning flow.

### 🧠 Services

#### 🌐 Translation Service
Handles all external API calls for translations and conjugations.
- Uses Lexica, Apertium, and Verbix integrations (with fallback logic in the works)

#### 📋 Management Service
Primarily for admin usage — assigns vocabulary (by POS type) to users per time period, structures the learning journey, and tracks completion.

---

## 📱 Frontend

React Native (Expo) app built with TypeScript and styled-components. Designed for swipey, daily learning on-the-go.

---

## 🛠️ Installation (Local Dev)

### Prerequisites:
- Node.js + Expo CLI
- Go 1.21+
- MongoDB (local or cloud)

### Run Frontend:
```bash
cd frontend
npm install
npm start
```

### Run Backend:
```bash
cd backend
go mod tidy
go run main.go
```
