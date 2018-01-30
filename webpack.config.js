const path = require('path');
const webpack = require('webpack');
const publicPath = 'http://localhost:3000';

module.exports = {
  entry: './frontend/src/index.js',
  devtool: 'cheap-module-source-map',

  plugins: [
    new webpack.HotModuleReplacementPlugin(),
  ],

  output: {
    path: path.resolve(__dirname, './frontend/public'),
    filename: '[name].bundle.js',
    publicPath: this.hmr ? 'http://localhost:3000/' : '',
    sourceMapFilename: '[name].map',
  },

  devServer: {
    port: 3000,
    host: 'localhost',
    historyApiFallback: true,
    noInfo: false,
    stats: 'minimal',
    publicPath: publicPath,
    contentBase: path.join(__dirname, publicPath),
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
