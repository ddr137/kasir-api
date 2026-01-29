# Kasir API

Simple Point of Sale (POS) API built with Go (Golang).

## 🚀 Production URL
**[http://gokasir.zeabur.app](http://gokasir.zeabur.app)**

## 📚 API Documentation
Swagger UI is available at:
- **Local**: `http://localhost:8094/swagger/index.html`
- **Production**: `http://gokasir.zeabur.app/swagger/index.html` (if deployed)

## 🛠 Tech Stack
- **Language**: Go 1.25+
- **Database**: PostgreSQL (Supabase)
- **Routing**: Standard `net/http`
- **Documentation**: Swagger (`swaggo/swag`)

## ✨ Features
- **Products**: CRUD, Pagination, Category Join.
- **Categories**: CRUD, Pagination.
- **RESTful Response**: Standard JSON format with metadata.

## 📦 Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/ddr137/kasir-api.git
   cd kasir-api
   ```

2. **Setup Database**
   - Create a PostgreSQL database (e.g., via Supabase).
   - Run the SQL script in `database/schema.sql`.

3. **Environment Variables**
   Create a `.env` file:
   ```env
   PORT=8094
   DB_SOURCE=postgresql://user:password@host:port/dbname?sslmode=require
   ```

4. **Run the Application**
   ```bash
   go run main.go
   ```

## 📝 API Endpoints

### Products
- `GET /api/products?page=1&page_size=10` - List products
- `POST /api/products` - Create product
- `GET /api/products/{id}` - Get product detail
- `PUT /api/products/{id}` - Update product
- `DELETE /api/products/{id}` - Delete product

### Categories
- `GET /api/categories` - List categories
- `POST /api/categories` - Create category
- `GET /api/categories/{id}` - Get category detail
- `PUT /api/categories/{id}` - Update category
- `DELETE /api/categories/{id}` - Delete category
