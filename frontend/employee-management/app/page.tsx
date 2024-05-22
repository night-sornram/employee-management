"use client"

import { useSession } from "next-auth/react";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";


export default function Home() {
  const { data: session } = useSession();
  const router = useRouter();
  return (
    <main className="bg-slate-100 w-screen h-screen">
        {
          session ? (
          <div className="flex flex-row h-screen justify-center">
              <div className="flex flex-col justify-center  self-center space-y-3">
                <h1 className="text-2xl font-semibold">Welcome to Employee Management System</h1>
                <Button onClick={()=>router.push("/attendance/checkin")} className="flex justify-center items-center">
                  Go to Dashboard
                </Button>
              </div>
            </div>
          )
          :
          (
            <div className="flex flex-row h-screen justify-center">
                <div className="flex flex-col justify-center self-center space-y-3">
                  <h1 className="text-2xl font-semibold">Please login to use service</h1>
                  <Button onClick={()=>router.push("/api/auth/signin")} className="flex justify-center items-center">
                    <h1>Go to Login</h1>
                  </Button>
                </div>
                
            </div>
          )}
    </main>
  );
}
