# Email Verifier

A command-line tool for verifying email domain configuration.

## Installation

Clone the repository:

~~~javascript
git clone https://github.com/yourusername/email-verifier.git
~~~

Install dependencies:

~~~javascript
cd email-verifier
go mod download
~~~

## 🛠 Usage
To verify a single domain:

~~~javascript
go run main.go example.com
~~~

To verify multiple domains, provide them as separate arguments:

~~~javascript
go run main.go example.com example.net example.org
~~~

# License
This project is licensed under the MIT[link](https://opensource.org/license/mit/) License.