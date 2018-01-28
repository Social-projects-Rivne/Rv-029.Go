const path = require('path');
const webpack = require('webpack');

module.exports = {
  devtool: 'cheap-module-eval-source-map',
  devServer: {
    port: 8080,
    hot: true,
    inline: true,
    headers: { 'Access-Control-Allow-Origin': '*' }
  },
  entry: {
    'app': [
      'babel-polyfill',
      'react-hot-loader/patch',
      './frontend/src/index.js'
    ],
  },
  output: {
    path: path.resolve(__dirname, './frontend/public'),
    filename: 'bundle.js',
    publicPath: 'http://localhost:8080/'
  },
  plugins: [
    new webpack.NamedModulesPlugin(),
    new webpack.HotModuleReplacementPlugin()
  ],
  module: {
    rules: [
      { test: /\.js$|\.jsx$/, exclude: /node_modules/, loader: 'babel-loader' },
      { test: /\.css$/, use: ['style-loader', 'css-loader'] },
      { test: /\.svg$|\.png$/, loader: 'file-loader' }
    ]
  },
};
