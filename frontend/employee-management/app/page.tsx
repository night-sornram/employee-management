"use client"

import Link from "next/link";
import { useSession } from "next-auth/react";


export default function Home() {
  const { data: session } = useSession();
  return (
    <main className="bg-slate-100 w-screen h-screen">
        {
          session ? (
          <div className="flex flex-row h-screen justify-center">
              <div className="flex flex-col justify-center self-center">
                <h1 className="text-2xl font-semibold">Welcome to Employee Management System</h1>
                <Link href="/attendance/checkin" className="custom-btn-dark hover:custom-btn-dark-hover mt-5 w-[50%] self-center text-center">
                  <h1>Go to Dashboard</h1>
                </Link>
              </div>
            </div>
          )
          :
          (
            <div className="flex flex-row h-screen justify-center">
                <div className="flex flex-col justify-center self-center">
                  <h1 className="text-2xl font-semibold">Please login to use service</h1>
                  <Link href="/api/auth/signin" className="custom-btn-dark hover:custom-btn-dark-hover mt-5 w-[50%] self-center text-center">
                    <h1>Go to Login</h1>
                  </Link>
                </div>
                
            </div>
          )}
    </main>
  );
}
