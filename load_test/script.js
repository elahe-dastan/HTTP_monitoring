import http from 'k6/http';

export default function () {
    var url = 'http://127.0.0.1:8080/register';
    var emailLength = 23;
    var passLength = 5;
    var characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    var email = '';
    var pass = '';
    for ( var i = 0; i < emailLength; i++ ) {
        email += characters.charAt(Math.floor(Math.random() * characters.length));
    }
    for ( i = 0; i < passLength; i++ ) {
        pass += characters.charAt(Math.floor(Math.random() * characters.length));
    }
    var payload = JSON.stringify({
        email: email,
        password: pass,
    });

    var params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);
}