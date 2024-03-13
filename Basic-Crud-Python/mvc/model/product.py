import psycopg2 as postgree
from flask import request


# Buat class untuk data yang akan disimpan di database
class Product:
    def __init__(self, id, name, price):
        self.id = id
        self.name = name
        self.price = price


class ProductDAO:
    def __init__(self, db_name, user, password, host, port):
        self.conn = postgree.connect(
            dbname=db_name, user=user, password=password, host=host, port=port
        )
        self.cur = self.conn.cursor()

    def get_all_products(self):
        self.cur.execute("SELECT id, name, price FROM products")
        rows = self.cur.fetchall()
        products = [Product(row[0], row[1], row[2]).__dict__ for row in rows]
        return products

    def get_product_by_id(self, id):
        self.cur.execute(f"SELECT id, name, price FROM products WHERE id={id}")
        row = self.cur.fetchone()
        if row:
            return Product(row[0], row[1], row[2]).__dict__
        return None

    def create_product(self):
        data = request.json
        name = data.get("name")
        price = data.get("price")
        self.cur.execute(
            "INSERT INTO products (name, price) VALUES (%s, %s) RETURNING id",
            (name, price),
        )
        id = self.cur.fetchone()[0]
        self.conn.commit()
        print(f"id = {id}")
        return Product(id, name, price).__dict__

    def update_product(self, id):
        data = request.json
        name = data.get("name")
        price = data.get("price")
        self.cur.execute(
            "UPDATE products SET name=%s, price=%s WHERE id=%s", (name, price, id)
        )
        self.conn.commit()
        if self.cur.rowcount == 0:
            return None
        return Product(id, name, price).__dict__

    def delete_product(self, id):
        self.cur.execute("DELETE FROM products WHERE id=%s", (id,))
        self.conn.commit()
        if self.cur.rowcount == 0:
            return None
        return Product(id, "", 0).__dict__
