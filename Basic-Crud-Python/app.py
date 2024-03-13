from flask import Flask, request, jsonify
import psycopg2
from enum import Enum
from mvc.controller.controller import ProductController


app = Flask(__name__)
port_number = 8000

controller = ProductController()


# Routes
@app.route("/products", methods=["GET"])
def get_products():
    return controller.get_all_products()


@app.route("/products/<int:id>", methods=["GET"])
def get_product(id):
    return controller.get_products_by_id(id)


@app.route("/products", methods=["POST"])
def create_product():
    return controller.create_product()


@app.route("/products/<int:id>", methods=["PUT"])
def update_product(id):
    return controller.update_product(id)


@app.route("/products/<int:id>", methods=["DELETE"])
def delete_product(id):
    return controller.delete_product(id)


if __name__ == "__main__":
    app.run(debug=True, port=port_number)
