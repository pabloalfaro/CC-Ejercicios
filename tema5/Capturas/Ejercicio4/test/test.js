const request = require('supertest');

const app = require('../ejer4.js');

request(app)
  .get('/')
  .expect('Content-Type', /json/)
  .expect(200)
  .end(function(err, res) {
    if (err) throw err;
  });
  
request(app)
  .get('/proc')
  .expect('Content-Type', /json/)
  .expect(200)
  .end(function(err, res) {
    if (err) throw err;
  });
