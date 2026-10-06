"use client";
import { GetProfile } from "@/lib/services/profile-service";

export default function Home() {
  void GetProfile().then();

  return (
    <div className="">
      <h1 className="text-white">Hello world</h1>
    </div>
  );
}
