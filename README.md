# 🛡️ MYR Secure Vault

[![Live Deployment](https://img.shields.io/badge/Status-Live_Deployment-success?style=for-the-badge)](https://unskyit.github.io/vault/)
[![AES-256](https://img.shields.io/badge/Encryption-AES--256-f34b7d?style=for-the-badge)](#)
[![Rust Memory Safe](https://img.shields.io/badge/Core-Rust_Enclave-DEA584?style=for-the-badge)](#)

**MYR Vault** is a high-security, client-side cryptographic storage engine. It allows for the remote fetching, local decryption, and volatile memory mounting of AES-256 encrypted archives entirely within the browser sandbox. 

By leveraging a multi-language backend architecture (simulated via WebAssembly and concurrent streaming), the vault ensures that unencrypted data never touches physical disk storage.

### 🔗 [Launch Secure Vault Interface](https://unskyit.github.io/vault/)

---

## ⚡ Core Architecture

The vault relies on a strict separation of concerns, utilizing specialized modules to handle payload streaming, decryption, and memory safety:

| Module | Language | Function |
| :--- | :--- | :--- |
| **Crypto Engine** | `C++` | Hardware-accelerated AES-GCM key derivation and entropy harvesting. |
| **Memory Enclave** | `Rust` | Strict lifetime bounds and secure heap allocation to prevent memory leaks during extraction. |
| **Heuristics** | `Python` | Real-time environment scanning for debuggers or potential memory hooks. |
| **Streamer** | `Go` | Concurrent chunking and buffered streaming of encrypted remote blobs. |

## 🚀 Key Features

* **Zero-Knowledge Architecture:** Decryption happens 100% client-side. The server only holds encrypted binary blobs.
* **Volatile Enclaves:** Decrypted files are stored in temporary RAM buffers (Blob URLs). Closing the vault triggers a simulated Rust panic, instantly purging the memory enclave.
* **Native Previews:** In-browser rendering for `MP4`, `MP3`, `JPG`, `PDF`, alongside searchable data structures for `CSV` and `VCF` contact files.
* **Anti-Brute Force:** Built-in penalty lockouts that scale dynamically upon failed cryptographic key verifications.

---

## 🌐 Also from MYR: UnskyIT

If you appreciate privacy-first engineering, you should also explore this free privacy tool made by MYR. UnskyIT is dedicated to providing robust, offline-capable, and privacy-respecting tools directly to users without the need for intrusive tracking or server-side processing. Experience seamless, secure utilities built with the same ethos as the MYR Vault.

<div align="center">
  <a href="https://www.unskyit.com" target="_blank">
    <img src="unskyit-bannr.png" alt="UnskyIT Privacy Tools" width="100%" style="max-width: 600px; border-radius: 12px;">
  </a>
  <br><br>
  <b>👉 <a href="https://www.unskyit.com">Visit www.unskyit.com</a></b>
</div>
