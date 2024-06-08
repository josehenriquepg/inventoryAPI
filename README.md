# Inventory Management API
This is a RESTful API for managing an inventory of products. It is built using Go (Golang) and provides basic CRUD operations for products.
## 📚 Features
- **Add Product**: Add a new product to the inventory.
- **List Products**: List all products in the inventory.
- **Get Product by ID**: Retrieve a specific product by its ID.
- **Update Product**: Update the details of an existing product.
- **Delete Product**: Remove a product from the inventory.
## 🧱 Data Structure
The product data structure is defined as follows:
```go
type Product struct {
  ID          int     `json:"id"`
  Name        string  `json:"name"`
  Description string  `json:"description"`
  Price       float64 `json:"price"`
  Quantity    int     `json:"quantity"`
}
```
## :gear: How to Use
#### Prerequisites
  - Go 1.16+ installed on your system.
#### Installation
  1. Clone the repository:
```bash
git clone https://github.com/josehenriquepg/microservice-inventory.git
cd microservice-inventory
```
  2. Build the project:
```bash
go build -o inventoryAPI
```
#### Running the Application
Run the compiled executable:
```bash
./inventoryAPI
```
The server will start running at `http://localhost:8000`
#### Endpoints
1. Add Product
  - URL: /products
  - Method: POST
  - Request Body:
```json
{
  "name": "Product Name",
  "description": "Product Description",
  "price": 99.99,
  "quantity": 10
}
```
  - Response:
```json
{
  "id": 1,
  "name": "Product Name",
  "description": "Product Description",
  "price": 99.99,
  "quantity": 10
}
```
2. List Products
  - URL: /products
  - Method: GET
  - Response:
```json
[
  {
    "id": 1,
    "name": "Product Name",
    "description": "Product Description",
    "price": 99.99,
    "quantity": 10
  }
]
```
3. Get Product by ID
  - URL: /products/{id}
  - Method: GET
  - Response:
```json
{
  "id": 1,
  "name": "Product Name",
  "description": "Product Description",
  "price": 99.99,
  "quantity": 10
}
```
4. Update Product
  - URL: /products/{id}
  - Method: PUT
  - Request Body:
```json
{
  "name": "Updated Name",
  "description": "Updated Description",
  "price": 109.99,
  "quantity": 15
}
```
  - Response:
```json
{
  "id": 1,
  "name": "Updated Name",
  "description": "Updated Description",
  "price": 109.99,
  "quantity": 15
}
```
5. Delete Product
  - URL: /products/{id}
  - Method: DELETE
  - Response:
```json
[
  {
    "id": 1,
    "name": "Product Name",
    "description": "Product Description",
    "price": 99.99,
    "quantity": 10
  }
]
```
#### Running Tests
The project includes unit tests for all endpoints using the net/http/httptest package. To run the tests, follow these steps:
1. Open a terminal and navigate to the project directory.
2. Run the tests using the following command:
```bash
go test -v
```
The -v flag is used to run the tests in verbose mode, which will display detailed output for each test case.
#### Code Structure
  - **main.go**: Contains the main logic of the application, including handlers for adding, listing, updating, and deleting products.
  - **main_test.go**: Contains the unit tests for the application.
## 💻 Technologies used
[![techs](https://skillicons.dev/icons?i=go,git&theme=dark)](https://skillicons.dev)
## 🤝 Collaborators 
<table>
  <tr>
    <td align="center">
      <a href="http://github.com/josehenriquepg">
        <img src="https://avatars.githubusercontent.com/josehenriquepg" width="100px;" alt="Foto de José Henrique no GitHub"/><br>
        <sub>
          <b>José Henrique</b>
        </sub>
      </a>
    </td>
  </tr>
</table>
