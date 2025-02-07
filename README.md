# BaiaoC2

**BaiaoC2** is a Command & Control (C2) platform under development, designed for advanced and customizable operations. This repository is in the initial development phase, and we are currently implementing the GUI using the Fyne framework to make the C2 more intuitive and efficient.

---

## 🚀 **Planned Features**

### **Supported Agents**
- **TCP**
- **HTTPS**
- **DoH (DNS over HTTPS)**
- **QUIC**
- **Pivot over TCP**

### **Frontend Features**
- **Multi-User Support**: Support for multiple users with admin and user roles.
- **Visual Interactions**: Graphs and visual interaction for listeners and agents.
- **Password Security**: Enforces strong password requirements.
- **2FA Authentication**: Two-factor authentication using Google MFA.
- **Websocket API Calls**: API calls using Websockets.

### **Implant Features**
- The **Implant** is written in **Golang** and is compatible with:
  - **Windows**
  - **Linux**
  - **MacOS**
  - (Mobile platforms are being considered for future updates).

### **Teamserver**
- The **Teamserver** is written in **.NET 6.0** and **does not require the .NET Core environment** to run.

### **Controller Features**
- **Reverse shell**.
- File management.
- Process management.
- Network traffic monitoring.
- Memory loading.
- Reverse proxy (based on **IOX**).
- Screenshots.
- **PE file loading in memory** for Windows/Linux.
- **Process injection and migration**.
- Allows files to be executed without touching the disk.

### **Memory Execution**
- Supports in-memory execution of:
  - **.NET Assemblies** (execute-assembly, inline-assembly).
  - **Custom RDI shellcode** (64-bit only; 32-bit requires manual client compilation).
  - Tools like **donut** and **Godonut** for shellcode generation.

### **Extensions and Integrations**
- **Custom extensions** via Lua scripts (similar to CNA scripts).
- Support for hosting binary files, text, and images in the Teamserver (similar to SimpleHttpServer).
- Telegram online host notifications by modifying `Chat ID` and `API Token` parameters in `profile.json`.

### **Performance**
- The **Controller** and **Teamserver** are lightweight and support high concurrency for intensive operations.

---

## 📋 **Development Status**
- The graphical user interface (GUI) is currently under development using the Fyne framework.
- The backend is being structured to fully integrate the features listed above.

---

## 💻 **Requirements**
- **Go 1.19+**
- **.NET 6.0** (for the Teamserver)
- **Fyne.io** (for the GUI)

---

## 🛠 **How to Run**
1. Clone this repository:
   ```bash
   git clone https://github.com/CyberSecurityUP/baiaoc2.git
   cd baiaoc2
   ```

2. Install dependencies:
   ```bash
   go get fyne.io/fyne/v2
   ```

3. Run the project:
   ```bash
   go run gui/main.go
   ```

---

## 📢 **Contributions**
We welcome contributions! Feel free to open issues or submit PRs to help improve **BaiaoC2**.

---

## 📄 **License**
This project is licensed under the [MIT License](LICENSE).
```

---

### **Changes Made**
1. Translated the content into English.
2. Updated the repository link and username to **CyberSecurityUP**.
3. Kept the same structure with relevant features and instructions.

Let me know if you need further adjustments or additional details! 🚀
