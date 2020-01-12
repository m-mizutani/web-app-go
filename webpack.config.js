const path = require('path');
const VueLoaderPlugin = require("vue-loader/lib/plugin");

module.exports = {
  mode: "development",
  entry: ["babel-polyfill", path.resolve("src", "js", "index.js")],
  output: {
    filename: "bundle.js",
    path: path.join(__dirname, "static/js/"),
    publicPath: "/js"
  },
  module: {
    rules: [
      {
        test: /\.vue$/,
        loader: "vue-loader"
      },
      {
        test: /\.js$/,
        loader: "babel-loader"
      },
      {
        test: /\.s[ac]ss$/i,
        use: [
          // Creates `style` nodes from JS strings
          'vue-style-loader',
          // Translates CSS into CommonJS
          'css-loader',
          // Compiles Sass to CSS
          'sass-loader',
        ],
      }
    ]
  },
  resolve: {
    extensions: [".js", "json", "jsx", "vue"],
    alias: {
      vue$: "vue/dist/vue.esm.js"
    }
  },
  devServer: {
    contentBase: "static",
    proxy: {
      "/api": "http://localhost:9080"
    }
  },
  plugins: [new VueLoaderPlugin()]
};

