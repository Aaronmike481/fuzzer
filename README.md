# Redhat Fuzzer

An AI Prompt Injection Scanner built with Go, Python, and Next.js.

## Architecture
- **Frontend:** Next.js (TypeScript + Tailwind)
- **Backend:** Go (Worker Pool via pyproc)
- **AI Logic:** Python (Mutation engine & Payload generation)

## Prerequisites
- Node.js (v22+)
- Go (1.21+)
- Python (3.9+)
- WSL 2 (If on Windows)

## How to Run (Development)

### 1. Start the Python Worker (inside WSL)
```bash
cd backend/python-worker
pip install pyproc-worker
python3 worker.py
