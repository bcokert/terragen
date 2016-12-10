"use strict";
require("webpack");
var path = require("path");

var BUILD_DIR = path.resolve(__dirname, "build/static");
var APP_DIR = path.resolve(__dirname, "static/js");

var config = {
    entry: {
        main: APP_DIR + "/index.jsx"
    },
    output: {
        path: BUILD_DIR,
        publicPath: "/static/",
        filename: "[name].js"
    },
    module: {
        loaders: [
            {test: /\.jsx?/, include: APP_DIR, loader: "babel"},
            {test: /\.css$/, loader: "style-loader!css-loader"},
            {test: /\.less$/, loader: "style-loader!css-loader!less-loader"},
            {test: /\.json$/, loader: "json"},
            {test: /\.glsl/, loader: "raw"},
            {test: /\.(jpe?g|gif|png|svg)$/i, loader: "url?limit=10000"}
        ],
        noParse: [
            /plotly\.js/
        ]
    },
    resolve: {
        extensions: ["", ".webpack.js", ".web.js", ".js", ".jsx"]
    }
};

module.exports = config;
