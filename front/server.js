import http from 'http';
import fs from 'fs';
import axios from 'axios'


http.createServer(async function (request, response) {
    switch (request.url) {
        case "/style.css":
            fs.readFile("./style.css", function (error, css) {
                response.writeHead(200, {"Content-Type": "text/css"});
                response.write(css);
                response.end();
            });
            break;
        case "/main.js": {
            fs.readFile("./main.js", function (error, js) {
                response.writeHead(200, {"Content-Type": "text/javascript"});
                response.write(js);
                response.end();
            });
            break;
        }
        case "/api/form":
            request.on('data', function (data) {
                axios.post('http://localhost:8080/api/form', data, {
                    headers: {
                        'Content-Type': 'application/json',
                    }
                })
                    .then(function (response) {
                        console.log(response.data);
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
            })
            break;
        default:
            fs.readFile("./index.html", function (error, html) {
                response.writeHead(200, {"Content-Type": "text/html"});
                response.write(html);
                response.end();
            })
    }
}).listen(3000)