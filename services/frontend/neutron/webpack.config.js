const path = require('path');
const HtmlWebPackPlugin = require('html-webpack-plugin');

module.exports = {
    entry: {
        app: ['./src/index.tsx'],
        vendor: ['react', 'react-dom']
    },
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: 'js/[name].bundle.js'
    },
    devtool: "source-map",
    resolve: {
        extensions: [".ts", ".tsx", ".js", ".jsx", ".json", ".css"]
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                loader: "awesome-typescript-loader"
            },
            {
                test: /\.css$/i,
                use: ['style-loader', 'css-loader'],
            },
        ]
    },

    plugins: [
        new HtmlWebPackPlugin({
            template: "./src/assets/index.html",
            filename: "./index.html"
        })
    ],
    devServer: {
        // contentBase: path.join(__dirname, 'dist'),
        compress: true,
        // Port
        port: 3000,
        host: 'localhost',
        // Automatically open page
        open: true,
        // Serves index.html (contains 404 page in react-router) in place of any 404 responses
        historyApiFallback: true,
        // Shows a full-screen overlay when there are compiler errors
        overlay: true,
    }
};