const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CopyPlugin = require("copy-webpack-plugin");

module.exports = {
    entry: {
        all: './web/assets/main.js'
    },
    mode: 'production',
    devtool: 'source-map',
    module: {
        rules: [
            {
                test: /\.scss$/i,
                use: [
                    {
                        loader: MiniCssExtractPlugin.loader
                    },
                    {
                        loader: 'css-loader',
                        options: {
                            url: false,
                        },
                    },
                    {
                        loader: 'sass-loader'
                    }
                ]
            },
            {
                test: /\.js$/,
                exclude: /node_modules/,
                loader: 'babel-loader'
            },
            // {
            //   test: /\.svg$/,
            //   use: 'file-loader'
            // }
        ]
    },
    output: {
        filename: 'javascript/[name].js',
        path: path.resolve(__dirname, 'web/static')
    },
    plugins: [
        new CopyPlugin({
            patterns: [
                {
                  from: path.resolve(__dirname, "web/assets/images"),
                  to: path.resolve(__dirname, "web/static/images")},
            ],
        }),
        new MiniCssExtractPlugin({
            filename: 'stylesheets/[name].css'
        }),
    ]
}