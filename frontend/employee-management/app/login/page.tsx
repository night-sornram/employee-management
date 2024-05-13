'use client'

import { useRouter } from "next/navigation";

export default function LoginPage () {

    const router = useRouter();

    const handleLogin= async (e: React.FormEvent<HTMLFormElement>) => {
        alert("login!");
        router.push('/attendance');
        router.refresh();
    }

    return (
      <div className="flex flex-row h-screen">
        <div className="bg-p-darkgray w-[60%] flex flex-col justify-center">
            <h1 className="text-white text-center">Cover Image</h1>
        </div>
        <div className="w-[40%] flex flex-col">
            <div className="px-[10%] pt-[30%]">
                <h1 className="text-2xl font-bold mb-[10%]">Login</h1>
                <form onSubmit={handleLogin}>
                    <div className="mb-[5%]">
                        <h2>ID</h2>
                        <input type="text" placeholder="Your ID" className="bg-p-lightgray placeholder-gray-500 rounded px-3 py-2 w-full"/>
                    </div>
                    <div className="mb-[5%]">
                        <h2>Password</h2>
                        <input type="password" placeholder="Your ID" className="bg-p-lightgray placeholder-gray-500 rounded px-3 py-2 w-full"/>
                    </div>
                    <div className="flex flex-col items-center">
                        <button type="submit" className="bg-p-darkgray px-3 py-2 rounded text-white">
                            Login
                        </button>
                    </div>  
                </form>
            </div>
        </div>
      </div>  
    );
}