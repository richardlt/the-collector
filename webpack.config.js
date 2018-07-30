const webpack = require('webpack'),
    HtmlWebpackPlugin = require('html-webpack-plugin'),
    path = require('path');

const packageJSON = require("./package.json");

const BUILD_DIR = path.resolve(__dirname, 'client/dist');
const APP_DIR = path.resolve(__dirname, 'client/src');
const PRODUCTION = process.env.NODE_ENV === 'production';

let config = {
    devtool: PRODUCTION ? "source-map" : "eval-source-map",
    entry: {
        bundle: APP_DIR + '/client.js',
        vendor: Object.keys(packageJSON.dependencies)
    },
    output: {
        path: BUILD_DIR,
        filename: PRODUCTION ? '[name].[chunkhash].js' : '[name].js'
    },
    module: {
        loaders: [{
            test: /\.js$/,
            loader: 'babel-loader',
            query: {
                presets: ['react', 'es2015']
            },
            include: APP_DIR
        }]
    },
    plugins: [
        new webpack.optimize.CommonsChunkPlugin({ name: 'vendor' }),
        new webpack.DefinePlugin({
            'process.env': {
                NODE_ENV: JSON.stringify(PRODUCTION ? 'production' : 'develoment')
            }
        }),
        new HtmlWebpackPlugin({
            title: 'The collector' + (!PRODUCTION ? ' dev' : ''),
            template: APP_DIR + '/index.ejs'
        })
    ]
};

if (PRODUCTION) {
    config.plugins = [
        new webpack.HashedModuleIdsPlugin(),
        new webpack.optimize.ModuleConcatenationPlugin(),
        new webpack.optimize.UglifyJsPlugin({
            sourceMap: true,
            compress: {
                warnings: false,
                drop_debugger: true
            },
            output: {
                comments: false
            }
        })
    ].concat(config.plugins);
} else {
    config.devServer = {
        port: 8081,
        contentBase: path.join(__dirname, 'client/dist'),
        compress: true,
        proxy: [{
            context: ["/api"],
            ws: true,
            target: "http://localhost:8080"
        }],
        historyApiFallback: true,
        hot: true
    };
    config.plugins = [
        new webpack.HotModuleReplacementPlugin(),
        new webpack.NamedModulesPlugin()
    ].concat(config.plugins);
}

module.exports = config;
