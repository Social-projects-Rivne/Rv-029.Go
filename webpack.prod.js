const publicPath = '/frontend/public/';
const path = require('path');
const webpack = require('webpack');

module.exports = {
 entry: './frontend/src/index.js',
 output: {
    path: path.join(__dirname, publicPath),
    filename: '[name].bundle.js',
    publicPath: publicPath,
    sourceMapFilename: '[name].map'
  },

  plugins: [
    new webpack.LoaderOptionsPlugin({
      minimize: true,
      debug: false
    }),
    new webpack.optimize.UglifyJsPlugin({
      beautify: false,
      mangle: {
        screw_ie8: true,
        keep_fnames: true
      },
      compress: {
        screw_ie8: true
      },
      comments: false
    })
  ]
}
