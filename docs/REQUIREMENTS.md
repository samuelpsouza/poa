## Functional Requirements

### Multitenant

Each tenant needs to have their own separate and secure storage space within the application, ensuring that their data is isolated from other tenants and only accessible to authorized users. 

#### Features
- data encryption ?
- user access control
- prevent data leakage between tenants

### Types of Authentication

Users should be able to select from various authentication methods such as basic, bearer token, oauth2, biometrics, and two-factor authentication. Eache option will be used to protect the mocked api endpoint so it allows them to test authentication mechanisms. This requirement ensures flexibility and adaptability to accommodate different user preferences and security needs.

### Echo authentication

The application should have the ability to accurately reproduce the exact input message provided by the user without any alterations or additions. This would ensure that the API is successfully echoing back the user's input in a timely and precise manner, allowing for seamless communication between the user and the application.

## Non-funcional Requirements

- Use ScyllaDB
