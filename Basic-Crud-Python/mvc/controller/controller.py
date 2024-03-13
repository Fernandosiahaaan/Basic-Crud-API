from ..model.product import ProductDAO as dao
from ..view.view import ProductView as view
from enum import Enum

db_name = "basic_database"
user = "postgres"
password = "postgre"
host = "localhost"
port = "5433"


class HTTP_Status_Code(Enum):
    StatusOK = 200
    StatusCreated = 201
    StatusMultipleChoices = 300
    StatusBadRequest = 400


class ProductController:
    def __init__(self):
        self.dao = dao(db_name, user, password, host, port)

    def get_all_products(self):
        products = self.dao.get_all_products()
        return view.response_with_data(
            HTTP_Status_Code.StatusOK.value, "Product Success Get", products
        )

    def get_products_by_id(self, id):
        products = self.dao.get_product_by_id(id)
        if products == None:
            return view.response_without_data(
                HTTP_Status_Code.StatusBadRequest.value, f"Not Found ID {id} Product "
            )
        else:
            return view.response_with_data(
                HTTP_Status_Code.StatusOK.value,
                f"Product Success Get ID {id}",
                products,
            )

    def create_product(self):
        products = self.dao.create_product()
        return view.response_with_data(
            HTTP_Status_Code.StatusOK.value,
            "Product Created Successfully",
            products,
        )

    def update_product(self, id):
        products = self.dao.update_product(id)
        if products == None:
            return view.response_without_data(
                HTTP_Status_Code.StatusBadRequest.value, f"Not Found Product ID {id}"
            )
        else:
            return view.response_with_data(
                HTTP_Status_Code.StatusOK.value,
                "Updated Product successfully",
                products,
            )

    def delete_product(self, id):
        products = self.dao.delete_product(id)
        if products == None:
            return view.response_without_data(
                HTTP_Status_Code.StatusBadRequest.value, f"Not Found Product ID {id}"
            )
        else:
            return view.response_without_data(
                HTTP_Status_Code.StatusOK.value,
                f"Product Deleted ID {id} successfully",
            )
