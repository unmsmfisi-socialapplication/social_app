/** @type {import('next').NextConfig} */
const path = require('path')
const nextConfig = {
    reactStricMode : true,
    sassOptions: {
        includePaths: [path.join(__dirname, 'styles')],
    },
}

module.exports = nextConfig
