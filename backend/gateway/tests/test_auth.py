import pytest
import requests

import tests.utils.auth as auth


def test_register_1_correct_user(api_url, mysql_conn):
    auth.delete_user_by_login(mysql_conn, auth.test_user["login"])

    user_auth = auth.Auth(api_url, user=auth.test_user)

    _ = user_auth.register_user(do_assertion=True)

def test_register_2_correct_user2(api_url, mysql_conn):
    auth.delete_user_by_login(mysql_conn, auth.test_user2["login"])

    user_auth = auth.Auth(api_url, user=auth.test_user2)

    _ = user_auth.register_user(do_assertion=True)

def test_register_3_correct_user3(api_url, mysql_conn):
    auth.delete_user_by_login(mysql_conn, auth.test_user3["login"])

    user_auth = auth.Auth(api_url, user=auth.test_user3)

    _ = user_auth.register_user(do_assertion=True)


def test_register_2_short_password(api_url):
    user_auth = auth.Auth(api_url, user={
            "login": "Example", 
            "password": "Asdf"
        })

    response = user_auth.register_user(do_assertion=False)

    assert response.status_code == 400
    # assert response.text == "password must be at least 8 characters long"

def test_register_3_long_password(api_url):
    user_auth = auth.Auth(api_url, user={
            "login": "Example", 
            "password": "0123456789012345678aA"
        })

    response = user_auth.register_user(do_assertion=False)

    assert response.status_code == 400

def test_register_4_password_equals_login(api_url):
    user_auth = auth.Auth(api_url, user={
            "login": "Example", 
            "password": "Example"
        })

    response = user_auth.register_user(do_assertion=False)

    assert response.status_code == 400

def test_register_5_password_containss_login(api_url):
    user_auth = auth.Auth(api_url, user={
            "login": "Example", 
            "password": "123Example321"
        })

    response = user_auth.register_user(do_assertion=False)

    assert response.status_code == 400

def test_register_6_password_no_uppercase(api_url):
    user_auth = auth.Auth(api_url, user={
            "login": "Example", 
            "password": "asdfasdf"
        })

    response = user_auth.register_user(do_assertion=False)

    assert response.status_code == 400

def test_register_7_password_no_lowercase(api_url):
    user_auth = auth.Auth(api_url, user={
            "login": "Example", 
            "password": "ASDFASDF"
        })

    response = user_auth.register_user(do_assertion=False)

    assert response.status_code == 400

def test_register_7_password_short_and_no_uppercase(api_url):
    user_auth = auth.Auth(api_url, user={
            "login": "Example", 
            "password": "asd1"
        })

    response = user_auth.register_user(do_assertion=False)

    assert response.status_code == 400
    assert len(response.json()["messages"]) == 2