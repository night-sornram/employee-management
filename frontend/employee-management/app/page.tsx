import NavBar from "@/components/NavBar/NavBar";
import TopBar from "@/components/TopBar/TopBar";
import Image from "next/image";
import Link from "next/link";

export default function Home() {
  return (
    <main className="bg-gray-200">
        <div className="flex flex-row h-screen justify-center">
            <div className="flex flex-col justify-center self-center">
              <h1 className="text-2xl font-semibold">Please login to use service</h1>
              <Link href="/login" className="custom-btn-dark hover:custom-btn-dark-hover mt-5 w-[50%] self-center">
                <h1>Go to Login</h1>
              </Link>
            </div>
            
        </div>
    </main>
  );
}
