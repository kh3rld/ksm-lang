# KISUMU Language

A programming language written in Go with a syntax similar to Go. KSM is designed to provide a basic yet functional programming environment with essential data structures and methods.

## Project Overview

KSM aims to be an educational tool for learning about programming languages and interpreters. It supports several fundamental data structures and provides a minimal set of methods for each, allowing for manipulation and interaction with data.

## Key Features

- **Syntax**: Similar to Go, making it easy to understand for those familiar with Go.
- **Data Structures**:
  - **Number**: Supports both integers and floats.
  - **String**: Basic string handling.
  - **Boolean**: True/false values.
  - **Null**: Represents the absence of a value.
  - **Array**: Ordered collection of elements with methods like `length` and `first`.
  - **Object/Hash**: Collection of key-value pairs with methods to access and modify elements.

## Installation

To install and use KSM, follow these steps:

1. **Clone the repository**:
    ```bash
    git clone https://github.com/kh3rld/ksm-lang.git
    cd ksm-lang && cd cmd
    ```

2. **Build the project**:
    ```bash
    go build -o ksm
    ```

3. **Run the interpreter**:
    ```bash
    ./ksm <path_to_your_ksm_file.ksm>
    ```

## Contributing

We welcome contributions to KSM! If you'd like to contribute, please fork the repository, make your changes, and submit a pull request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For questions, suggestions, or support, please reach out to us via GitHub issues or contact us directly through the project's page.
