import http from 'k6/http';

export default function () {
    var url = 'http://127.0.0.1:8080/register';
    var payload = JSON.stringify({
        email: 'parham.alvani@gmail.com',
        password: '1378',
    });

    var params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);
}