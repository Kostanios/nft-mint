/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  experimental: {
    appDir: true,
  },
  env: {
    WEBAPI: process.env.WEBAPI,
  },
}

export default nextConfig
