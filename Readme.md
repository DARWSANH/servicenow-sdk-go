# ServiceNow SDK for Go 🚀

![ServiceNow SDK for Go](https://img.shields.io/badge/ServiceNow%20SDK%20for%20Go-v1.0.0-blue.svg)  
[![GitHub Releases](https://img.shields.io/github/release/DARWSANH/servicenow-sdk-go.svg)](https://github.com/DARWSANH/servicenow-sdk-go/releases)

Welcome to the ServiceNow SDK for Go! This library allows Go developers to interact with the ServiceNow API seamlessly. With this SDK, you can integrate ServiceNow functionalities into your Go applications with ease.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [API Reference](#api-reference)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Features

- **Simple Integration**: Connect to ServiceNow with minimal setup.
- **Uniform Interface**: Interact with various ServiceNow APIs using a consistent method.
- **Error Handling**: Built-in error handling to simplify debugging.
- **Comprehensive Documentation**: Detailed examples and API documentation.

## Installation

To install the ServiceNow SDK for Go, you can use the Go package manager. Run the following command in your terminal:

```bash
go get github.com/DARWSANH/servicenow-sdk-go
```

## Getting Started

To get started, you will need to have a ServiceNow account and API credentials. Follow these steps:

1. **Create a ServiceNow Account**: Sign up at [ServiceNow](https://www.servicenow.com).
2. **Generate API Credentials**: Go to your ServiceNow instance and create an API user. Make sure to assign the necessary roles.
3. **Set Up Your Go Environment**: Ensure you have Go installed on your machine. You can download it from [golang.org](https://golang.org).

## Usage

Here’s a simple example to demonstrate how to use the ServiceNow SDK in your Go application:

```go
package main

import (
    "fmt"
    "github.com/DARWSANH/servicenow-sdk-go"
)

func main() {
    // Initialize the ServiceNow client
    client := servicenow.NewClient("your_instance", "your_username", "your_password")

    // Fetch incidents
    incidents, err := client.GetIncidents()
    if err != nil {
        fmt.Println("Error fetching incidents:", err)
        return
    }

    // Print incidents
    for _, incident := range incidents {
        fmt.Printf("Incident Number: %s, Short Description: %s\n", incident.Number, incident.ShortDescription)
    }
}
```

This code initializes a new ServiceNow client and fetches a list of incidents. Replace `your_instance`, `your_username`, and `your_password` with your actual ServiceNow credentials.

## API Reference

The SDK provides various methods to interact with the ServiceNow API. Here are some key functions:

### Client Methods

- **NewClient(instance, username, password)**: Initializes a new ServiceNow client.
- **GetIncidents()**: Retrieves a list of incidents from ServiceNow.
- **CreateIncident(incident)**: Creates a new incident in ServiceNow.
- **UpdateIncident(incident)**: Updates an existing incident.

### Incident Struct

The SDK uses the following struct to represent an incident:

```go
type Incident struct {
    Number          string `json:"number"`
    ShortDescription string `json:"short_description"`
    Description     string `json:"description"`
    State           string `json:"state"`
}
```

For a complete list of methods and their usage, please refer to the [API documentation](https://github.com/DARWSANH/servicenow-sdk-go/wiki).

## Contributing

We welcome contributions! If you would like to contribute to the ServiceNow SDK for Go, please follow these steps:

1. **Fork the Repository**: Click the "Fork" button on the top right of the repository page.
2. **Clone Your Fork**: Use the following command to clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/servicenow-sdk-go.git
   ```
3. **Create a Branch**: Create a new branch for your feature or bug fix:
   ```bash
   git checkout -b feature/your-feature-name
   ```
4. **Make Changes**: Make your changes and commit them.
5. **Push to Your Fork**: Push your changes back to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```
6. **Create a Pull Request**: Go to the original repository and create a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or support, please reach out to the repository maintainer at [your-email@example.com](mailto:your-email@example.com).

## Releases

To download the latest release, visit our [Releases page](https://github.com/DARWSANH/servicenow-sdk-go/releases). You can download the latest version and execute it in your Go environment.

For further updates and changes, always check the [Releases section](https://github.com/DARWSANH/servicenow-sdk-go/releases).

---

Thank you for using the ServiceNow SDK for Go! We hope it simplifies your integration with ServiceNow. Happy coding!