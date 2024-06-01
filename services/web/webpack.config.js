const path = require('path');

module.exports = {
    mode: 'development',
    entry: {
        main: './src/main.ts'
    },
    output: {
        path: path.join(__dirname,'dist'),
        filename: '[name].js'
    },
    resolve: {
        extensions:['.ts','.js']
    },
    module: {
        rules: [
            {
                test: /\.ts$/,
                use: 'ts-loader'
            }
        ]
    },
    devServer: {
        static: {
            directory: path.join(__dirname, "dist"),
        },
        compress: true,
        port: 8000
    },
}
