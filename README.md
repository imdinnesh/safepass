# Safepass - Zero Trust API Gateway (Roadmap & Features)

Safepass is a high-performance, zero-trust API gateway written in Go. Designed to be developer-friendly, self-hosted, and secure-by-default, Safepass acts as a smart proxy that protects your backend services with authentication, rate limiting, auditing, IP firewall, and plugin-based extensibility.

## 🌐 Why Safepass?

Modern apps rely on APIs, and APIs need protection. Safepass puts a powerful gatekeeper in front of your APIs, ensuring:  
- Only authorized requests get through  
- Rate limits and policies are enforced  
- API keys, tokens, and sessions are centrally managed  
- Infrastructure is hardened without touching your app code  

---

## 🧩 Core Features (v1 MVP)

### 🔐 Authentication & Authorization
- ✅ JWT verification with support for RS256/HS256  
- ✅ Per-route auth enforcement (optional/public/protected)  
- ✅ API key support for programmatic clients  
- ✅ Role-based access control (RBAC)  
- ✅ Redis-backed OTP session store (multi-device/session-aware)  

### 🛡️ Security Enforcement
- ✅ IP allow/block lists  
- ✅ Dynamic IP banning for repeated abuse  
- ✅ Path-level route protection rules  
- ✅ Signature verification for HMAC-secured routes  
- ✅ Expiring session tokens with refresh flow  

### 📈 Rate Limiting & Throttling
- ✅ Per-route and per-user rate limits  
- ✅ Redis-based rate limiting engine  
- ✅ Burst & sliding window strategies  
- ✅ Global IP rate control  

### 🔁 Request Proxy & Routing
- ✅ Reverse proxy with intelligent routing  
- ✅ Path-based service forwarding (`/api/users/*` → service A)  
- ✅ Request rewriting support  
- ✅ Custom header injection (e.g., `X-User-ID`, `X-Role`)  

### 🪝 Extensibility (Plugins)
- ✅ Plugin system (Go interface)  
- ✅ Lua/WASM plugin support (optional)  
- ✅ Lifecycle hooks: `onRequest`, `onAuth`, `onResponse`  
- ✅ External plugin loading via config  

### 📊 Monitoring & Audit
- ✅ Central request logging (method, status, latency, user)  
- ✅ Audit logs for access violations  
- ✅ Redis or file-based logging backend  
- ✅ Integration with Prometheus (planned)  

### ⚙️ Configuration & Admin
- ✅ YAML-based route config (`safepass.yaml`)  
- ✅ Hot-reload support for config changes  
- ✅ Web Admin UI (optional, Next.js frontend)  
- ✅ CLI tool (`spg`) for managing:  
  - API keys  
  - Route rules  
  - Token/session cache  

---

## 📦 Use Cases

- Protect Go/Express/FastAPI backends without modifying them  
- Centralize auth and rate-limiting for multiple microservices  
- Implement secure ingress for internal tools/dashboards  
- Create fine-grained API gateways for public APIs  
- Manage device-based logins and multi-session control  

---