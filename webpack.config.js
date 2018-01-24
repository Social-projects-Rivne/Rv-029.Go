const path = require('path');

const config = {
  entry: './frontend/src/index.js',

  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'frontend/public')
  },

  devtool: 'inline-source-map',

  module: {
    rules:[
      {
        test: /\.css$/,
        use: ['style-loader', 'css-loader']
      },
      {
        test: /\.svg$|\.jpg$/,
        use: {
          loader: 'file-loader',
          options: {
            publicPath: 'public/'
          }
        }
      },
      {
        test: /\.js$/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ["react", "es2015", "stage-2"]
          }
        }
      }
    ]
  }
};

module.exports = config;
