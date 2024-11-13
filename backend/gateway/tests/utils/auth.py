import requests
import jwt


test_user = {
    "login": "Example", 
    "password": "Asdfasdf"
}

test_user2 = {
    "login": "Director", 
    "password": "iDirecTor42"
}

test_user3 = {
    "login": "TestCustomer", 
    "password": "MyUniquePswd1"
}

def delete_user_by_login(mysql, user_login):
    mysql.cursor().execute(
        f"DELETE FROM `users` WHERE `login` = '{user_login}'"
    )
    mysql.commit()

class Auth:
    user = test_user
    access_token = ""
    refresh_token = ""
    cookies = None

    def __init__(self, api_url, user=test_user):
        self.api_url = api_url
        self.user = user
    
    def register_user(self, do_assertion=True):
        response = requests.post(f"{self.api_url}/register", json=self.user)

        if do_assertion:
            self.__assert_tokens_response(response)

        self.access_token = response.cookies.get("access_token")
        self.refresh_token = response.cookies.get("refresh_token")
        self.tokens_cookies = response.cookies

        return response

    def __assert_tokens_response(self, response: requests.Response):
        assert response.status_code == 200

        new_access_token = response.cookies.get("access_token")
        new_refresh_token = response.cookies.get("refresh_token")
        # Проверка cookie токенов
        assert new_access_token is not None
        assert new_refresh_token is not None

        # Проверка параметров cookie токенов
        access_token_cookie = next(
            (c for c in response.cookies if c.name == "access_token"), None
        )
        assert access_token_cookie is not None
        assert not access_token_cookie.has_nonstandard_attr("HttpOnly")

        refresh_token_cookie = next(
            (c for c in response.cookies if c.name == "refresh_token"), None
        )
        assert refresh_token_cookie is not None
        assert refresh_token_cookie.has_nonstandard_attr("HttpOnly")

    def __assert_refresh_response(self, response: requests.Response):
        # Проверка ответа и новых токенов
        self.__assert_tokens_response(response)

        # Сравнение новых и старых токенов
        new_access_token = response.cookies.get("access_token")
        new_refresh_token = response.cookies.get("refresh_token")

        assert new_access_token != self.access_token
        assert new_refresh_token != self.refresh_token

    def __assert_logout_response(self, response: requests.Response):
        assert response.status_code == 200

        assert response.cookies.get("access_token") is None
        assert response.cookies.get("refresh_token") is None
