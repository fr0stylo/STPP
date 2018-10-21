const express = require('express');
const httpProxy = require('express-http-proxy');
const jwt = require('express-jwt');
const cors = require('cors');

const jwksRsa = require('jwks-rsa');

const app = express();
app.use(cors())

const routes = require('./routes.json');

const checkJwt = jwt({
	secret: jwksRsa.expressJwtSecret({
		cache: true,
		rateLimit: true,
		jwksRequestsPerMinute: 5,
		jwksUri: `https://stpp.eu.auth0.com/.well-known/jwks.json`
	}),
	audience: 'dBDJmkD4Htne3e3lX1x7rCDFWKl0hRMH',
	issuer: `https://stpp.eu.auth0.com/`,
	algorithms: ['RS256']
});

console.log("started");

for (let i in routes) {
	app.use(routes[i].route, checkJwt,async (req, res, next) => {
		console.time('request');
		console.log(`Route called ${routes[i].route} passed to ${routes[i].proxy}`);

		await httpProxy(routes[i].proxy)(req, res, next);
		
		console.timeEnd('request');
	});
}

app.listen(3000, () => {
	console.log('Listening on port', 3000);
});