import { version } from "./package.json";
import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  reactCompiler: true,
  env: {
    APP_VERSION: version,
    API_URL: "http://localhost:8080/api/",
    //TODO:IMPLEMENTS LOGIN AND MIDDLEWARE REQUEST
    TOKEN:
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3NzI0MDcyNzAsInVzZXJJZCI6NX0.GiBQAnf_bokaFjOm16bHgWqYpFSgM6wY6dcknmq7U1I",
  },
};

export default nextConfig;
