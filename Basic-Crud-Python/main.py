from flask import Flask, request, jsonify
import psycopg2
from enum import Enum

app = Flask(__name__)

# Koneksi ke database
port_number = 8000
conn = psycopg2.connect(
    dbname="basic_database",
    user="postgres",
    password="postgre",
    host="localhost",
    port="5433",
)
cur = conn.cursor()


# Buat class untuk data yang akan disimpan di database
class Product:
    def __init__(self, id, name, price):
        self.id = id
        self.name = name
        self.price = price


class HTTP_Status_Code(Enum):
    StatusOK = 200
    StatusCreated = 201
    StatusMultipleChoices = 300
    StatusBadRequest = 400


# Routes
@app.route("/products", methods=["GET"])
def get_products():
    print(f"StatusOK = {HTTP_Status_Code.StatusOK.value}")
    cur.execute("SELECT id, name, price FROM products")
    rows = cur.fetchall()
    products = [Product(row[0], row[1], row[2]).__dict__ for row in rows]
    print(f"products = {products}")
    return jsonify(
        {
            "status": HTTP_Status_Code.StatusOK.value,
            "message": "Produk berhasil diambil",
            "data": products,
        }
    )


@app.route("/products/<int:id>", methods=["GET"])
def get_product(id):
    cur.execute("SELECT id, name, price FROM products WHERE id=%s", (id,))
    row = cur.fetchone()
    if row:
        product = Product(row[0], row[1], row[2]).__dict__
        return jsonify(
            {
                "status": HTTP_Status_Code.StatusOK.value,
                "message": "Produk berhasil diambil",
                "data": product,
            }
        )
    else:
        return jsonify(
            {
                "status": HTTP_Status_Code.StatusBadRequest.value,
                "message": "Produk Error Request",
            }
        )


@app.route("/products", methods=["POST"])
def create_product():
    data = request.json
    name = data.get("name")
    price = data.get("price")
    cur.execute(
        "INSERT INTO products (name, price) VALUES (%s, %s) RETURNING id", (name, price)
    )
    id = cur.fetchone()[0]
    conn.commit()
    product = Product(id, name, price).__dict__
    return jsonify(
        {
            "status": HTTP_Status_Code.StatusCreated.value,
            "message": "Product created successfully",
            "data": product,
        }
    )


@app.route("/products/<int:id>", methods=["PUT"])
def update_product(id):
    data = request.json
    name = data.get("name")
    price = data.get("price")
    cur.execute("UPDATE products SET name=%s, price=%s WHERE id=%s", (name, price, id))
    conn.commit()
    if cur.rowcount == 0:
        return jsonify(
            {
                "status": HTTP_Status_Code.StatusBadRequest.value,
                "message": f"Tidak ada produk dengan id : {id}",
            }
        )
    else:
        product = Product(id, name, price).__dict__
        return jsonify(
            {
                "status": HTTP_Status_Code.StatusOK.value,
                "message": "Product updated successfully",
                "data": product,
            }
        )


@app.route("/products/<int:id>", methods=["DELETE"])
def delete_product(id):
    cur.execute("DELETE FROM products WHERE id=%s", (id,))
    conn.commit()
    if cur.rowcount == 0:
        return jsonify(
            {
                "status": HTTP_Status_Code.StatusBadRequest.value,
                "message": f"Tidak ada produk dengan id : {id}",
            }
        )
    else:
        return jsonify(
            {
                "status": HTTP_Status_Code.StatusOK.value,
                "message": "Product deleted successfully",
            }
        )


@app.route("/products", methods=["DELETE"])
def delete_products():
    cur.execute("DELETE FROM products")
    conn.commit()
    return jsonify(
        {
            "status": HTTP_Status_Code.StatusOK.value,
            "message": "Product deleted successfully",
        }
    )


if __name__ == "__main__":
    app.run(debug=True, port=port_number)
