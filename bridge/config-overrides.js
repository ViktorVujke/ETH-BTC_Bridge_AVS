const webpack = require('webpack');

module.exports = function override(config, env) {
  config.resolve.fallback = {
    buffer: require.resolve('buffer/'),
    stream: require.resolve('stream-browserify')
  };

  config.plugins = (config.plugins || []).concat([
    new webpack.ProvidePlugin({
      Buffer: ['buffer', 'Buffer'],
    }),
  ]);

  // Add WebAssembly support
  config.experiments = {
    asyncWebAssembly: true,
    syncWebAssembly: true,
  };

  // Add rule to handle WASM files
  config.module.rules.push({
    test: /\.wasm$/,
    type: 'webassembly/async',
  });

  return config;
};
