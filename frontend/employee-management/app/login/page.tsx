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
        <div className="bg-slate-800 w-[60%] flex flex-col justify-center">
            <h1 className="text-white text-center">Some Organization</h1>
        </div>
        <div className="w-[40%] flex flex-col">
            <div className="px-[10%] pt-[30%]">
                <h1 className="text-2xl font-bold mb-[10%]">Login</h1>
                <form onSubmit={handleLogin}>
                    <div className="mb-[5%]">
                        <h2>ID</h2>
                        <input type="text" placeholder="Your ID" className="bg-slate-100 placeholder-slate-400 rounded px-3 py-2 w-full
                        border border-slate-300 shadow-sm focus:outline-none focus:border-slate-400 invalid:border-red-600 invalid:text-red-600 focus:invalid:border-red-700"/>
                    </div>
                    <div className="mb-[5%]">
                        <h2>Password</h2>
                        <input type="password" placeholder="Password" className="bg-slate-100 placeholder-slate-400 rounded px-3 py-2 w-full
                        border border-slate-300 shadow-sm focus:outline-none focus:border-slate-400 invalid:border-red-600 invalid:text-red-600 focus:invalid:border-red-700"/>
                    </div>
                    <div className="flex flex-col items-center">
                        <button type="submit" className="custom-btn-dark hover:custom-btn-dark-hover">
                            Login
                        </button>
                    </div>  
                </form>
            </div>
        </div>
      </div>  
    );
}