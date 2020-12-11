'use strict'

// copy from `https://github.com/flipped-aurora/gin-vue-admin/blob/master/web/vue.config.js`


const path = require('path')


function __resolve(dir) {
  return path.join(__dirname, dir)
}

module.exports = {
  // base config
  publicPath: './',
  outputDir: 'dist',
  assetsDir: 'static',
  lintOnSave: process.env.NODE_ENV === 'development',
  productionSourceMap: false,
  devServer: {
    port: 8080,
    open: true,
    overlay: {
      warning: false,
      errors: true
    },
    proxy: {
       // 把key的路径代理到target位置
      // detail: https://cli.vuejs.org/config/#devserver-proxy
      [process.env.VUE_APP_BASE_API]: { //需要代理的路径   例如 '/api'
        target: `http://127.0.0.1:8888/`, //代理到 目标路径
        changeOrigin: true,
        pathRewrite: { // 修改路径数据
          ['^' + process.env.VUE_APP_BASE_API]: '' // 举例 '^/api:""' 把路径中的/api字符串删除
        }
      }
    }
  },
  configureWebpack: {
    resolve: {
      alias: {
        '@': __resolve('src')
      }
    }
  },
  chainWebpack(config) {
    // set preserveWhitespace
    config.module
      .rule('vue')
      .use('vue-loader')
      .loader('vue-loader')
      .tap(options => {
          options.compilerOptions.preserveWhitespace = true
          return options
      })
      .end()
    config
    // https://webpack.js.org/configuration/devtool/#development
        .when(process.env.NODE_ENV === 'development',
        config => config.devtool('cheap-source-map')
    )

    config
        .when(process.env.NODE_ENV !== 'development',
            config => {
                // 不打包 begin
                // 1.目前已经测试通过[vue,axios,echarts]可以cdn引用，其它组件测试通过后可继续添加
                // 2.此处添加不打包后，需在public/index.html head中添加相应cdn资源链接
                config.set('externals', {
                  vue: 'Vue',
                  axios: 'axios',
                  echarts: 'echarts'
                });
                // 不打包 end
                config
                  .plugin('ScriptExtHtmlWebpackPlugin')
                  .after('html')
                  .use('script-ext-html-webpack-plugin', [{
                      // `runtime` must same as runtimeChunk name. default is `runtime`
                      inline: /single\..*\.js$/
                  }])
                  .end();
                config.optimization.splitChunks({
                    chunks: 'all',
                    cacheGroups: {
                      libs: {
                        name: 'chunk-libs',
                        test: /[\\/]node_modules[\\/]/,
                        priority: 10,
                        chunks: 'initial' // only package third parties that are initially dependent
                      },
                      elementUI: {
                        name: 'chunk-elementUI', // split elementUI into a single package
                        priority: 20, // the weight needs to be larger than libs and app or it will be packaged into libs or app
                        test: /[\\/]node_modules[\\/]_?element-ui(.*)/ // in order to adapt to cnpm
                      },
                      commons: {
                        name: 'chunk-commons',
                        test: resolve('src/components'), // can customize your rules
                        minChunks: 3, //  minimum common number
                        priority: 5,
                        reuseExistingChunk: true
                      }
                    }
                  })
                config.optimization.runtimeChunk('single')
            }
        )
  }
}