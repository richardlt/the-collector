var webpack = require('webpack'),
    path = require('path');

var nodeModulesPath = path.resolve(__dirname, 'node_modules'),
    buildPath = path.resolve(__dirname, './client/dist/js'),
    appPath = path.resolve(__dirname, './client/src/client.js');

var webpackConfig = {

    devtool: 'source-map',

    entry: appPath,

    output: {
        path: buildPath,
        filename: 'bundle.js'
    },

    module: {
        loaders: [{
            test: /\.js$/,
            loader: 'babel-loader',
            query: {
                presets: ['react', 'es2015']
            },
            exclude: nodeModulesPath
        }]
    }

};

module.exports = webpackConfig;
