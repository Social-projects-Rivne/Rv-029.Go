const path = require('path');
const webpack = require('webpack');

module.exports = {
  entry:'./frontend/src/index.js',
  devtool: 'cheap-module-source-map',

  plugins: [
    new webpack.HotModuleReplacementPlugin(),
  ],

  output: {
    path: path.resolve(__dirname, './frontend/public'),
    filename: '[name].bundle.js',
    sourceMapFilename: '[name].map',
  },

  devServer: {
    port: 80,
    host: '0.0.0.0',
    historyApiFallback: true,
    noInfo: false,
    stats: 'minimal',
    publicPath: "/",
    contentBase: path.resolve(__dirname, './frontend/public'),
    headers: { 'Access-Control-Allow-Origin': '*' },
    inline: true,
    hot: true
  },

  module: {
    rules: [
      { 
        test: /\.svg$|\.png$/, 
        loader: 'file-loader' 
      },
      {
        test: /\.(png|svg|jpg|gif)$/,
        use: ['file-loader']
      },
      {
        test: /\.js|.jsx?$/,
        exclude: /(node_modules)/,
        loaders: ["babel-loader"]
      }]
  },
}
