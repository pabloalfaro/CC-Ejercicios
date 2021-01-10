#!/usr/bin/env node

var express=require('express');
var app = express();
var port = process.env.PORT || 8080;

const coche_marca = "Opel";

app.get('/marca', function (req, res) {
    res.send( coche_marca );
});

app.get('/', function (req, res) {
    res.send( { Marca: true } );
});

app.listen(port);
console.log('Server running at http://127.0.0.1:'+port+'/');
