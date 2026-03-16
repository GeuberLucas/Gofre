"use client";
import { useEffect } from "react";
import { DoLogin } from "../services/auth-service";

export default function Login() {
  useEffect(() => {
    DoLogin("gebe", "123456");
  }, []);

  return <div>Login Page</div>;
}
