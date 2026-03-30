import { version } from "./package.json";
import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  reactCompiler: true,
  env: {
    APP_VERSION: version,
    API_URL: "http://localhost:8080/api/",
  },
};

export default nextConfig;
