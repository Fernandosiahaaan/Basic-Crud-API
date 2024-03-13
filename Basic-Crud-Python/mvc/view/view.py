from flask import jsonify, request


class ProductView:
    def response_with_data(status, message, data):
        message_json = jsonify(
            {
                "status": status,
                "message": message,
                "data": data,
            }
        )

        return (
            message_json,
            status,
        )

    def response_without_data(status, message):
        message_json = jsonify(
            {
                "status": status,
                "message": message,
            }
        )
        return (
            message_json,
            status,
        )
