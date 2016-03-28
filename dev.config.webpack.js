var webpack = require('webpack'),
    path = require('path');

var nodeModulesPath = path.resolve(__dirname, 'node_modules'),
    buildPath = path.resolve(__dirname, './client/dist/js'),
    appPath = path.resolve(__dirname, './client/src/client.js');

var webpackConfig = {

    devtool: 'eval',

    entry: [
        'webpack/hot/dev-server',
        'webpack-dev-server/client?http://localhost:8081',
        appPath
    ],

    output: {
        path: buildPath,
        filename: 'bundle.js',
        publicPath: '/js/'
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
    },

    plugins: [new webpack.HotModuleReplacementPlugin()]
};

module.exports = webpackConfig;
